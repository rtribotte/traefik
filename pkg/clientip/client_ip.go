package clientip

import (
	"context"
	"net"
	"net/http"
)

// clientIPKey is the context key under which the resolved client IP is stored.
type clientIPKey struct{}

// WithClientIP returns a context carrying the given resolved client IP.
func WithClientIP(ctx context.Context, addr string) context.Context {
	return context.WithValue(ctx, clientIPKey{}, addr)
}

// ClientIP returns the client IP (the resolved client IP or the remote addr) for the given request.
func ClientIP(req *http.Request) string {
	if addr, ok := fromContext(req.Context()); ok {
		return addr
	}

	host, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return req.RemoteAddr
	}

	return host
}

// fromContext returns the resolved client IP previously stored with
// WithClientIP and whether it was present.
func fromContext(ctx context.Context) (string, bool) {
	addr, ok := ctx.Value(clientIPKey{}).(string)
	return addr, ok
}
