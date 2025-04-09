package tls

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ocsp"
)

const certWithOCSPServer = `-----BEGIN CERTIFICATE-----
MIIBgjCCASegAwIBAgICIAAwCgYIKoZIzj0EAwIwEjEQMA4GA1UEAxMHVGVzdCBD
QTAeFw0yMzAxMDExMjAwMDBaFw0yMzAyMDExMjAwMDBaMCAxHjAcBgNVBAMTFU9D
U1AgVGVzdCBDZXJ0aWZpY2F0ZTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABIoe
I/bjo34qony8LdRJD+Jhuk8/S8YHXRHl6rH9t5VFCFtX8lIPN/Ll1zCrQ2KB3Wlb
fxSgiQyLrCpZyrdhVPSjXzBdMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgwFoAU+Eo3
5sST4LRrwS4dueIdGBZ5d7IwLAYIKwYBBQUHAQEEIDAeMBwGCCsGAQUFBzABhhBv
Y3NwLmV4YW1wbGUuY29tMAoGCCqGSM49BAMCA0kAMEYCIQDg94xY/+/VepESdvTT
ykCwiWOS2aCpjyryrKpwMKkR0AIhAPc/+ZEz4W10OENxC1t+NUTvS8JbEGOwulkZ
z9yfaLuD
-----END CERTIFICATE-----`

const certWithoutOCSPServer = `-----BEGIN CERTIFICATE-----
MIIBUzCB+aADAgECAgIgADAKBggqhkjOPQQDAjASMRAwDgYDVQQDEwdUZXN0IENB
MB4XDTIzMDEwMTEyMDAwMFoXDTIzMDIwMTEyMDAwMFowIDEeMBwGA1UEAxMVT0NT
UCBUZXN0IENlcnRpZmljYXRlMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEih4j
9uOjfiqifLwt1EkP4mG6Tz9LxgddEeXqsf23lUUIW1fyUg838uXXMKtDYoHdaVt/
FKCJDIusKlnKt2FU9KMxMC8wDAYDVR0TAQH/BAIwADAfBgNVHSMEGDAWgBT4Sjfm
xJPgtGvBLh254h0YFnl3sjAKBggqhkjOPQQDAgNJADBGAiEA3rWetLGblfSuNZKf
5CpZxhj3A0BjEocEh+2P+nAgIdUCIQDIgptabR1qTLQaF2u0hJsEX2IKuIUvYWH3
6Lb92+zIHg==
-----END CERTIFICATE-----`

// certKey is the private key for both certWithOCSPServer and certWithoutOCSPServer.
const certKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEINnVcgrSNh4HlThWlZpegq14M8G/p9NVDtdVjZrseUGLoAoGCCqGSM49
AwEHoUQDQgAEih4j9uOjfiqifLwt1EkP4mG6Tz9LxgddEeXqsf23lUUIW1fyUg83
8uXXMKtDYoHdaVt/FKCJDIusKlnKt2FU9A==
-----END EC PRIVATE KEY-----`

// caCert is the issuing certificate for certWithOCSPServer and certWithoutOCSPServer.
const caCert = `-----BEGIN CERTIFICATE-----
MIIBazCCARGgAwIBAgICEAAwCgYIKoZIzj0EAwIwEjEQMA4GA1UEAxMHVGVzdCBD
QTAeFw0yMzAxMDExMjAwMDBaFw0yMzAyMDExMjAwMDBaMBIxEDAOBgNVBAMTB1Rl
c3QgQ0EwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASdKexSor/aeazDM57UHhAX
rCkJxUeF2BWf0lZYCRxc3f0GdrEsVvjJW8+/E06eAzDCGSdM/08Nvun1nb6AmAlt
o1cwVTAOBgNVHQ8BAf8EBAMCAQYwEwYDVR0lBAwwCgYIKwYBBQUHAwkwDwYDVR0T
AQH/BAUwAwEB/zAdBgNVHQ4EFgQU+Eo35sST4LRrwS4dueIdGBZ5d7IwCgYIKoZI
zj0EAwIDSAAwRQIgGbA39+kETTB/YMLBFoC2fpZe1cDWfFB7TUdfINUqdH4CIQCR
ByUFC8A+hRNkK5YNH78bgjnKk/88zUQF5ONy4oPGdQ==
-----END CERTIFICATE-----`

const caKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIDJ59ptjq3MzILH4zn5IKoH1sYn+zrUeq2kD8+DD2x+OoAoGCCqGSM49
AwEHoUQDQgAEnSnsUqK/2nmswzOe1B4QF6wpCcVHhdgVn9JWWAkcXN39BnaxLFb4
yVvPvxNOngMwwhknTP9PDb7p9Z2+gJgJbQ==
-----END EC PRIVATE KEY-----`

func TestOCSPStapler_Upsert(t *testing.T) {
	ocspStapler := newOCSPStapler(nil)

	issuerCert, err := tls.X509KeyPair([]byte(caCert), []byte(caKey))
	require.NoError(t, err)

	leafCert, err := tls.X509KeyPair([]byte(certWithOCSPServer), []byte(certKey))
	require.NoError(t, err)

	// Upsert a certificate without an OCSP server should raise an error.
	leafCertWithoutOCSPServer, err := tls.X509KeyPair([]byte(certWithoutOCSPServer), []byte(certKey))
	require.NoError(t, err)

	err = ocspStapler.Upsert("foo", leafCertWithoutOCSPServer.Leaf, issuerCert.Leaf)
	require.Error(t, err)

	// Upsert a certificate with an OCSP server.
	err = ocspStapler.Upsert("foo", leafCert.Leaf, issuerCert.Leaf)
	require.NoError(t, err)

	i, ok := ocspStapler.cache.Get("foo")
	require.True(t, ok)

	e, ok := i.(*ocspEntry)
	require.True(t, ok)

	assert.Equal(t, leafCert.Leaf, e.leaf)
	assert.Equal(t, issuerCert.Leaf, e.issuer)
	assert.Nil(t, e.staple)
	assert.Equal(t, []string{"ocsp.example.com"}, e.responders)
	assert.Equal(t, int64(0), ocspStapler.cache.Items()["foo"].Expiration)

	// Upsert an existing entry to make sure that the existing staple is preserved.
	e.staple = []byte("foo")
	e.nextUpdate = time.Now()
	e.responders = []string{"foo.com"}

	err = ocspStapler.Upsert("foo", leafCert.Leaf, issuerCert.Leaf)
	require.NoError(t, err)

	i, ok = ocspStapler.cache.Get("foo")
	require.True(t, ok)

	e, ok = i.(*ocspEntry)
	require.True(t, ok)

	assert.Equal(t, leafCert.Leaf, e.leaf)
	assert.Equal(t, issuerCert.Leaf, e.issuer)
	assert.Equal(t, []byte("foo"), e.staple)
	assert.NotZero(t, e.nextUpdate)
	assert.Equal(t, []string{"foo.com"}, e.responders)
	assert.Equal(t, int64(0), ocspStapler.cache.Items()["foo"].Expiration)
}

func TestOCSPStapler_Upsert_withResponderOverrides(t *testing.T) {
	ocspStapler := newOCSPStapler(map[string]string{
		"ocsp.example.com": "foo.com",
	})

	issuerCert, err := tls.X509KeyPair([]byte(caCert), []byte(caKey))
	require.NoError(t, err)

	leafCert, err := tls.X509KeyPair([]byte(certWithOCSPServer), []byte(certKey))
	require.NoError(t, err)

	err = ocspStapler.Upsert("foo", leafCert.Leaf, issuerCert.Leaf)
	require.NoError(t, err)

	i, ok := ocspStapler.cache.Get("foo")
	require.True(t, ok)

	e, ok := i.(*ocspEntry)
	require.True(t, ok)

	assert.Equal(t, leafCert.Leaf, e.leaf)
	assert.Equal(t, issuerCert.Leaf, e.issuer)
	assert.Nil(t, e.staple)
	assert.Equal(t, []string{"foo.com"}, e.responders)
}

func TestOCSPStapler_ResetTTL(t *testing.T) {
	ocspStapler := newOCSPStapler(nil)

	issuerCert, err := tls.X509KeyPair([]byte(caCert), []byte(caKey))
	require.NoError(t, err)

	leafCert, err := tls.X509KeyPair([]byte(certWithOCSPServer), []byte(certKey))
	require.NoError(t, err)

	ocspStapler.cache.Set("foo", &ocspEntry{
		leaf:       leafCert.Leaf,
		issuer:     issuerCert.Leaf,
		responders: []string{"foo.com"},
		nextUpdate: time.Now(),
		staple:     []byte("foo"),
	}, cache.NoExpiration)

	ocspStapler.cache.Set("bar", &ocspEntry{
		leaf:       leafCert.Leaf,
		issuer:     issuerCert.Leaf,
		responders: []string{"bar.com"},
		nextUpdate: time.Now(),
		staple:     []byte("bar"),
	}, time.Hour)

	wantBarExpiration := ocspStapler.cache.Items()["bar"].Expiration

	ocspStapler.ResetTTL()

	item, ok := ocspStapler.cache.Items()["foo"]
	require.True(t, ok)

	e, ok := item.Object.(*ocspEntry)
	require.True(t, ok)

	assert.True(t, item.Expiration > 0)
	assert.Equal(t, leafCert.Leaf, e.leaf)
	assert.Equal(t, issuerCert.Leaf, e.issuer)
	assert.Equal(t, []byte("foo"), e.staple)
	assert.NotZero(t, e.nextUpdate)
	assert.Equal(t, []string{"foo.com"}, e.responders)

	item, ok = ocspStapler.cache.Items()["bar"]
	require.True(t, ok)

	e, ok = item.Object.(*ocspEntry)
	require.True(t, ok)

	assert.Equal(t, wantBarExpiration, item.Expiration)
	assert.Equal(t, leafCert.Leaf, e.leaf)
	assert.Equal(t, issuerCert.Leaf, e.issuer)
	assert.Equal(t, []byte("bar"), e.staple)
	assert.NotZero(t, e.nextUpdate)
	assert.Equal(t, []string{"bar.com"}, e.responders)
}

func TestOCSPStapler_GetStaple(t *testing.T) {
	ocspStapler := newOCSPStapler(nil)

	// Get an un-existing staple.
	staple, exists := ocspStapler.GetStaple("foo")

	assert.False(t, exists)
	assert.Nil(t, staple)

	// Get an existing staple.
	ocspStapler.cache.Set("foo", &ocspEntry{staple: []byte("foo")}, cache.NoExpiration)

	staple, exists = ocspStapler.GetStaple("foo")

	assert.True(t, exists)
	assert.Equal(t, []byte("foo"), staple)
}

func startOCSPResponder(t *testing.T, ocspResponse []byte) *httptest.Server {
	t.Helper()

	handler := func(rw http.ResponseWriter, req *http.Request) {
		ct := req.Header.Get("Content-Type")
		assert.Equal(t, ct, "application/ocsp-request")

		reqBytes, err := io.ReadAll(req.Body)
		require.NoError(t, err)

		_, err = ocsp.ParseRequest(reqBytes)
		require.NoError(t, err)

		rw.Header().Set("Content-Type", "application/ocsp-response")

		_, err = rw.Write(ocspResponse)
		require.NoError(t, err)
	}
	return httptest.NewServer(http.HandlerFunc(handler))
}
