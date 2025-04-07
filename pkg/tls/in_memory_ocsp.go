package tls

import (
	"bytes"
	"context"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/ocsp"
)

// Constants for PKIX MustStaple extension.
var (
	tlsFeatureExtensionOID = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 1, 24}
	ocspMustStapleFeature  = []byte{0x30, 0x03, 0x02, 0x01, 0x05}
	mustStapleExtension    = pkix.Extension{
		Id:    tlsFeatureExtensionOID,
		Value: ocspMustStapleFeature,
	}
)

type ocspEntry struct {
	leaf   *x509.Certificate
	issuer *x509.Certificate
	resp   *ocsp.Response
	staple []byte
}

type InMemoryOCSPCache struct {
	cache cache.Cache

	responderOverrides map[string]string
}

// NewInMemoryOCSPCache creates a new inMemoryOCSPCache cache.
func NewInMemoryOCSPCache(responderOverrides map[string]string) *InMemoryOCSPCache {
	inMemoryOCSPCache := &InMemoryOCSPCache{
		cache:              *cache.New(30*time.Minute, 5*time.Minute),
		responderOverrides: responderOverrides,
	}

	return inMemoryOCSPCache
}

// Get retrieves the inMemoryOCSPCache response from the cache.
func (o *InMemoryOCSPCache) Get(key string) ([]byte, bool) {
	if item, ok := o.cache.Get(key); ok && item != nil {
		if entry, ok := item.(*ocspEntry); ok {
			return entry.staple, true
		}
	}

	return nil, false
}

// SetAllItemsTTL sets the expiration time for all items in the cache.
func (o *InMemoryOCSPCache) SetAllItemsTTL(ttl time.Duration) {
	for _, item := range o.cache.Items() {
		if item.Expiration > 0 {
			continue
		}

		item.Expiration = time.Now().Add(ttl).UnixNano()
	}
}

func (o *InMemoryOCSPCache) Set(key string, leaf, issuer *x509.Certificate) {
	o.cache.Set(key, &ocspEntry{
		leaf:   leaf,
		issuer: issuer,
	}, cache.NoExpiration)
}

func (o *InMemoryOCSPCache) SetNoTTL(key string) {
	if item, ok := o.cache.Get(key); ok && item != nil {
		o.cache.Set(key, item, cache.NoExpiration)
	}
}

func (o *InMemoryOCSPCache) Run(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		return
	case <-ticker.C:
		for _, item := range o.cache.Items() {
			entry := item.Object.(*ocspEntry)

			staple, response, err := staple(entry.leaf, entry.resp)
			if err != nil {
				log.Error().Err(err).Msgf("stapling error for %s", entry.leaf.Subject.CommonName)
				continue
			}

			entry.staple = staple
			entry.resp = response
		}
	}
}

// Staple populates the ocsp response of the certificate if needed and not disabled by configuration.
func staple(leaf *x509.Certificate, lastResponse *ocsp.Response) ([]byte, *ocsp.Response, error) {
	if lastResponse != nil &&
		time.Now().Before(lastResponse.ThisUpdate.Add(lastResponse.NextUpdate.Sub(lastResponse.ThisUpdate)/2)) {
		return nil, nil, nil
	}

	ocspRespBytes, ocspResp, ocspErr := o.request()
	if ocspErr != nil {
		return nil, nil, fmt.Errorf("no inMemoryOCSPCache stapling for %v: %w", leaf.Subject.CommonName, ocspErr)
	}

	log.Debug().Msgf("ocsp response: %v", ocspResp)
	if ocspResp.Status == ocsp.Good {
		if ocspResp.NextUpdate.After(o.leaf.NotAfter) {
			return nil, nil, fmt.Errorf("invalid: inMemoryOCSPCache response for %v valid after certificate expiration (%s)", o.leaf.Subject.CommonName, o.leaf.NotAfter.Sub(ocspResp.NextUpdate))
		}

		return ocspRespBytes, ocspResp, nil
	}

	// FIXME
	return nil, nil, errors.New("ocsp response status not good")
}

func (o *ocspEntry) request() ([]byte, *ocsp.Response, error) {
	// TODO: check FIPS compliance for SHA1 used as default hash
	ocspReq, err := ocsp.CreateRequest(o.leaf, o.issuer, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("creating inMemoryOCSPCache request: %w", err)
	}

	if len(o.leaf.OCSPServer) == 0 {
		return nil, nil, errors.New("no inMemoryOCSPCache server specified in certificate")
	}

	ocspServer := o.leaf.OCSPServer
	if len(config.ResponderOverrides) > 0 {
		for i, respURL := range issuedCertificate.OCSPServer {
			if override, ok := config.ResponderOverrides[respURL]; ok {
				ocspServer[i] = override
			}
		}
	}

	reader := bytes.NewReader(ocspReq)
	resp, err := http.Post(ocspServer[0], "application/ocsp-request", reader)
	if err != nil {
		return nil, nil, fmt.Errorf("making inMemoryOCSPCache request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return nil, nil, fmt.Errorf("response error: %d", resp.StatusCode)
	}

	ocspResBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("reading inMemoryOCSPCache response: %w", err)
	}

	// FIXME: use ParseResponseForCert
	ocspRes, err := ocsp.ParseResponse(ocspResBytes, o.issuer)
	if err != nil {
		return nil, nil, fmt.Errorf("parsing inMemoryOCSPCache response: %w", err)
	}

	return ocspResBytes, ocspRes, nil
}
