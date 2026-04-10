package forwardedheaders

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/ip"
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

func TestXForwarded_ResolveClientIP_UntrustedPeer_StripsSourceHeader(t *testing.T) {
	capture := &captureHandler{}
	m, err := NewXForwarded(Config{
		TrustedIPs:              []string{"10.0.0.0/8"},
		ResolveClientIP:         true,
		ClientIPHeader:          "CF-Connecting-IP",
		ComputeFullForwardedFor: true,
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
	// No resolved IP in context since peer is untrusted.
	_, ok := clientip.FromContext(capture.req.Context())
	assert.False(t, ok)
}

func TestXForwarded_ResolveClientIP_TrustedPeer_StashesContextAndXRealIP(t *testing.T) {
	capture := &captureHandler{}
	m, err := NewXForwarded(Config{
		TrustedIPs:              []string{"10.0.0.0/8"},
		ResolveClientIP:         true,
		ClientIPHeader:          "CF-Connecting-IP",
		ComputeFullForwardedFor: true,
	}, capture)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	req.RemoteAddr = "10.0.0.1:1234" // trusted
	req.Header.Set("CF-Connecting-IP", "1.2.3.4")

	m.ServeHTTP(nil, req)

	resolved, ok := clientip.FromContext(capture.req.Context())
	assert.True(t, ok)
	assert.Equal(t, "1.2.3.4", resolved)
	// X-Real-IP written by rewrite() reflects the resolved client, not the peer.
	assert.Equal(t, "1.2.3.4", capture.req.Header.Get("X-Real-Ip"))
}

func TestXForwarded_ResolveClientIP_XForwardedForChain(t *testing.T) {
	capture := &captureHandler{}
	m, err := NewXForwarded(Config{
		TrustedIPs:              []string{"10.0.0.0/8"},
		ResolveClientIP:         true,
		ClientIPHeader:          "X-Forwarded-For",
		ComputeFullForwardedFor: true,
	}, capture)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	req.RemoteAddr = "10.0.0.1:1234"
	req.Header.Set("X-Forwarded-For", "1.2.3.4, 10.0.0.2")

	m.ServeHTTP(nil, req)

	resolved, ok := clientip.FromContext(capture.req.Context())
	assert.True(t, ok)
	assert.Equal(t, "1.2.3.4", resolved)
	// ComputeFullForwardedFor=true preserves the incoming XFF chain as-is.
	assert.Equal(t, "1.2.3.4, 10.0.0.2", capture.req.Header.Get("X-Forwarded-For"))
	// Proxy append is still enabled.
	assert.False(t, httputil.ShouldNotAppendXFF(capture.req.Context()))
}

func TestXForwarded_ResolveClientIP_ComputeFullForwardedFor_False(t *testing.T) {
	capture := &captureHandler{}
	m, err := NewXForwarded(Config{
		TrustedIPs:              []string{"10.0.0.0/8"},
		ResolveClientIP:         true,
		ClientIPHeader:          "X-Forwarded-For",
		ComputeFullForwardedFor: false,
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

func TestXForwarded_ResolveClientIP_Disabled_NoBehaviorChange(t *testing.T) {
	capture := &captureHandler{}
	m, err := NewXForwarded(Config{
		Insecure: true,
	}, capture)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	req.RemoteAddr = "10.0.0.1:1234"
	req.Header.Set("X-Forwarded-For", "1.2.3.4")

	m.ServeHTTP(nil, req)

	// No context key, no X-Real-IP override beyond the historical behavior.
	_, ok := clientip.FromContext(capture.req.Context())
	assert.False(t, ok)
	assert.Equal(t, "10.0.0.1", capture.req.Header.Get("X-Real-Ip"))
	assert.Equal(t, "1.2.3.4", capture.req.Header.Get("X-Forwarded-For"))
}
