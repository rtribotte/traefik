package forwardedheaders

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/clientip"
	"github.com/traefik/traefik/v3/pkg/proxy/httputil"
)

// captureHandler records the request it sees so the test can inspect what
// XForwarded passed downstream.
type captureHandler struct {
	req *http.Request
}

func (c *captureHandler) ServeHTTP(_ http.ResponseWriter, r *http.Request) {
	c.req = r
}

func TestXForwarded_ClientIPResolution_UntrustedPeer_StripsSourceHeader(t *testing.T) {
	capture := &captureHandler{}
	m, err := NewXForwarded(Config{
		TrustedIPs:         []string{"10.0.0.0/8"},
		ClientIPResolution: true,
		ClientIPHeader:     "CF-Connecting-IP",
	}, capture)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	req.RemoteAddr = "8.8.8.8:1234" // untrusted
	req.Header.Set("CF-Connecting-IP", "1.2.3.4")
	req.Header.Set("X-Forwarded-For", "1.2.3.4")

	m.ServeHTTP(nil, req)

	// Untrusted peer: both XFF and the configured source header must be stripped.
	assert.Empty(t, capture.req.Header.Get("CF-Connecting-IP"))
	assert.Empty(t, capture.req.Header.Get("X-Forwarded-For"))
	// No resolved IP in context since peer is untrusted — ClientIP falls back
	// to RemoteAddr.
	assert.Equal(t, "8.8.8.8", clientip.ClientIP(capture.req))
}

func TestXForwarded_ClientIPResolution_TrustedPeer_StashesContextAndXRealIP(t *testing.T) {
	capture := &captureHandler{}
	m, err := NewXForwarded(Config{
		TrustedIPs:         []string{"10.0.0.0/8"},
		ClientIPResolution: true,
		ClientIPHeader:     "CF-Connecting-IP",
	}, capture)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	req.RemoteAddr = "10.0.0.1:1234" // trusted
	req.Header.Set("CF-Connecting-IP", "1.2.3.4")

	m.ServeHTTP(nil, req)

	// Resolved IP is in context — ClientIP returns the resolved value.
	assert.Equal(t, "1.2.3.4", clientip.ClientIP(capture.req))
	// X-Real-IP written by rewrite() reflects the resolved client, not the peer.
	assert.Equal(t, "1.2.3.4", capture.req.Header.Get("X-Real-Ip"))
}

func TestXForwarded_ClientIPResolution_XForwardedForChain(t *testing.T) {
	capture := &captureHandler{}
	m, err := NewXForwarded(Config{
		TrustedIPs:         []string{"10.0.0.0/8"},
		ClientIPResolution: true,
		ClientIPHeader:     "X-Forwarded-For",
	}, capture)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	req.RemoteAddr = "10.0.0.1:1234"
	req.Header.Set("X-Forwarded-For", "1.2.3.4, 10.0.0.2")

	m.ServeHTTP(nil, req)

	assert.Equal(t, "1.2.3.4", clientip.ClientIP(capture.req))
	// ClientIPReplaceXFF=false (default) preserves the incoming XFF chain.
	assert.Equal(t, "1.2.3.4, 10.0.0.2", capture.req.Header.Get("X-Forwarded-For"))
	// Proxy append is still enabled.
	assert.False(t, httputil.ShouldNotAppendXFF(capture.req.Context()))
}

func TestXForwarded_ClientIPResolution_ClientIPReplaceXFF_True(t *testing.T) {
	capture := &captureHandler{}
	m, err := NewXForwarded(Config{
		TrustedIPs:         []string{"10.0.0.0/8"},
		ClientIPResolution: true,
		ClientIPHeader:     "X-Forwarded-For",
		ClientIPReplaceXFF: true,
	}, capture)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	req.RemoteAddr = "10.0.0.1:1234"
	req.Header.Set("X-Forwarded-For", "1.2.3.4, 10.0.0.2")

	m.ServeHTTP(nil, req)

	// Chain is replaced with just the resolved client IP.
	assert.Equal(t, "1.2.3.4", capture.req.Header.Get("X-Forwarded-For"))
	// Proxy append is disabled so the resolved-only value survives to the backend.
	assert.True(t, httputil.ShouldNotAppendXFF(capture.req.Context()))
}

func TestXForwarded_ClientIPResolution_Disabled_NoBehaviorChange(t *testing.T) {
	capture := &captureHandler{}
	m, err := NewXForwarded(Config{
		Insecure: true,
	}, capture)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	req.RemoteAddr = "10.0.0.1:1234"
	req.Header.Set("X-Forwarded-For", "1.2.3.4")

	m.ServeHTTP(nil, req)

	// No resolution — ClientIP falls back to RemoteAddr.
	assert.Equal(t, "10.0.0.1", clientip.ClientIP(capture.req))
	assert.Equal(t, "10.0.0.1", capture.req.Header.Get("X-Real-Ip"))
	assert.Equal(t, "1.2.3.4", capture.req.Header.Get("X-Forwarded-For"))
}
