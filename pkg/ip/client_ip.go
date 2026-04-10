package ip

import (
	"context"
	"net"
	"net/http"
)

// clientIPKey is the context key under which the resolved client IP is stored
// by the forwarded headers middleware when the entrypoint is configured to
// resolve the real client IP from a trusted-peer header chain.
type clientIPKey struct{}

// WithClientIP returns a context carrying the given resolved client IP.
// It is meant to be called by the forwarded headers middleware once the
// entrypoint-level trust walk has produced a resolved address.
func WithClientIP(ctx context.Context, addr string) context.Context {
	return context.WithValue(ctx, clientIPKey{}, addr)
}

// FromContext returns the resolved client IP previously stored with
// WithClientIP and whether it was present.
func FromContext(ctx context.Context) (string, bool) {
	addr, ok := ctx.Value(clientIPKey{}).(string)
	return addr, ok
}

// ClientIP returns the resolved client IP for the given request.
// When the entrypoint has resolved a real client IP and stashed it in the
// request context, that value is returned. Otherwise it falls back to the
// host portion of req.RemoteAddr (matching the historical Traefik behavior).
// The returned value is always stripped of any port.
func ClientIP(req *http.Request) string {
	if addr, ok := FromContext(req.Context()); ok {
		return addr
	}
	host, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return req.RemoteAddr
	}
	return host
}
