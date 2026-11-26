package server

import (
	"context"
	"net"
	"net/http"
	"sync"
	"sync/atomic"

	accesslogMiddle "github.com/traefik/traefik/v3/pkg/middlewares/tcp/accesslog"
	tcprouter "github.com/traefik/traefik/v3/pkg/server/router/tcp"
)

type servedKey struct{}

// unservedConnLogger logs in the TCP access logs the connections an HTTP server closes without serving any request,
// e.g. because of a failed TLS handshake or bytes which are not an HTTP request.
// The HTTP access logs never see these connections, as no request reaches the HTTP routers.
// The http.ConnState hook cannot tell them apart by itself,
// as a connection becomes active as soon as its first byte is read.
type unservedConnLogger struct {
	// accessLogger is nil when the TCP access logs are disabled.
	accessLogger *accesslogMiddle.EntryPointLogger

	// served holds, for each tracked connection, whether a request reached the handler.
	served sync.Map
}

// connContext tracks the given connection, when the TCP access logs are enabled.
func (u *unservedConnLogger) connContext(ctx context.Context, c net.Conn) context.Context {
	if u.accessLogger == nil {
		return ctx
	}

	served := &atomic.Bool{}
	u.served.Store(c, served)

	return context.WithValue(ctx, servedKey{}, served)
}

// connState logs the given connection when it is closed without having served any request.
func (u *unservedConnLogger) connState(c net.Conn, state http.ConnState) {
	// A hijacked connection never reports the closed state, and has served a request.
	if state != http.StateClosed && state != http.StateHijacked {
		return
	}

	served, ok := u.served.LoadAndDelete(c)
	if !ok || state == http.StateHijacked || served.(*atomic.Bool).Load() {
		return
	}

	tcprouter.LogUnservedHTTPConn(u.accessLogger, c)
}

// wrap marks the connections of the requests reaching the given handler as served.
func (u *unservedConnLogger) wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if served, ok := req.Context().Value(servedKey{}).(*atomic.Bool); ok {
			served.Store(true)
		}

		next.ServeHTTP(rw, req)
	})
}
