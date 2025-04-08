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
	leaf       *x509.Certificate
	issuer     *x509.Certificate
	responders []string
	nextUpdate time.Time
	staple     []byte
}

// InMemoryOCSPStapler retrieves staples from OCSP responders and store them in an in-memory cache.
// It also updates the staples on a regular basis and before they expire.
type InMemoryOCSPStapler struct {
	client             *http.Client // FIXME: timeout?
	ocspEntries        cache.Cache
	responderOverrides map[string]string
}

// NewInMemoryOCSPStapler creates a new InMemoryOCSPStapler cache.
func NewInMemoryOCSPStapler(responderOverrides map[string]string) *InMemoryOCSPStapler {
	return &InMemoryOCSPStapler{
		client:             &http.Client{},
		ocspEntries:        *cache.New(30*time.Minute, 5*time.Minute),
		responderOverrides: responderOverrides,
	}
}

// FIXME: force refresh?

// Run updates the OCSP staples every 5 minutes.
func (i *InMemoryOCSPStapler) Run(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		return

	case <-ticker.C:
		for _, item := range i.ocspEntries.Items() {
			select {
			case <-ctx.Done():
				return
			default:
			}

			entry := item.Object.(*ocspEntry)

			if entry.staple != nil && time.Now().Before(entry.nextUpdate) {
				continue
			}

			if err := i.updateStaple(ctx, entry); err != nil {
				log.Error().Err(err).Msgf("Unable to retieve OCSP staple for: %s", entry.leaf.Subject.CommonName)
				continue
			}
		}
	}
}

// GetStaple retrieves the OCSP obtainStaple from corresponding to the given key (public certificate hash).
func (i *InMemoryOCSPStapler) GetStaple(key string) ([]byte, bool) {
	if item, ok := i.ocspEntries.Get(key); ok && item != nil {
		if entry, ok := item.(*ocspEntry); ok {
			return entry.staple, true
		}
	}
	return nil, false
}

// Insert creates a new entry for the given certificate.
// The stapler will then be responsible from retrieving and updating the corresponding OCSP obtainStaple.
// FIXME: should we fetch the entry there?
func (i *InMemoryOCSPStapler) Insert(key string, leaf, issuer *x509.Certificate) error {
	if len(leaf.OCSPServer) == 0 {
		return errors.New("leaf certificate does not contain an OCSP server")
	}

	var responders []string
	for _, url := range leaf.OCSPServer {
		if newURL, ok := i.responderOverrides[url]; ok {
			responders = append(responders, newURL)
		}
	}

	entry := &ocspEntry{
		leaf:       leaf,
		issuer:     issuer,
		responders: responders,
	}
	i.ocspEntries.Set(key, entry, cache.NoExpiration)

	return nil
}

// SetAllItemsTTL sets the expiration time for all items in the cache.
func (i *InMemoryOCSPStapler) SetAllItemsTTL(ttl time.Duration) {
	for _, item := range i.cache.Items() {
		if item.Expiration > 0 {
			continue
		}

		item.Expiration = time.Now().Add(ttl).UnixNano()
	}
}

func (i *InMemoryOCSPStapler) SetNoTTL(key string) {
	if item, ok := i.cache.Get(key); ok && item != nil {
		i.cache.Set(key, item, cache.NoExpiration)
	}
}

// obtainStaple obtains the OCSP stable for the given leaf certificate.
func (i *InMemoryOCSPStapler) updateStaple(ctx context.Context, entry *ocspEntry) error {
	// TODO: check FIPS compliance for SHA1 used as default hash, if set the hash options.
	ocspReq, err := ocsp.CreateRequest(entry.leaf, entry.issuer, nil)
	if err != nil {
		return fmt.Errorf("creating OCSP request: %w", err)
	}

	for _, responder := range entry.responders {
		logger := log.With().Str("responder", responder).Logger()

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, responder, bytes.NewReader(ocspReq))
		if err != nil {
			return fmt.Errorf("creating OCSP request: %w", err)
		}

		req.Header.Set("Content-Type", "application/ocsp-request")

		res, err := i.client.Do(req)
		if err != nil && ctx.Err() != nil {
			return ctx.Err()
		}
		if err != nil {
			logger.Debug().Err(err).Msg("Unable to obtain OCSP response")
			continue
		}
		defer res.Body.Close()

		if res.StatusCode/100 != 2 {
			logger.Debug().Msgf("Unable to obtain OCSP response due to status code: %d", res.StatusCode)
			continue
		}

		ocspResBytes, err := io.ReadAll(res.Body)
		if err != nil {
			logger.Debug().Err(err).Msg("Unable to read OCSP response bytes")
			continue
		}

		ocspRes, err := ocsp.ParseResponseForCert(ocspResBytes, entry.leaf, entry.issuer)
		if err != nil {
			logger.Debug().Err(err).Msg("Unable to parse OCSP response")
			continue
		}

		entry.staple = ocspResBytes
		entry.nextUpdate = ocspRes.ThisUpdate.Add(ocspRes.NextUpdate.Sub(ocspRes.ThisUpdate) / 2)

		return nil
	}

	return errors.New("no OCSP staple obtained from any responders")
}
