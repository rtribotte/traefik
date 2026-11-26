package tcpaccesslog

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
)

type servedKey struct{}

// HTTPConnLogger logs the connections an HTTP server closes without serving any request,
// e.g. because of a failed TLS handshake or bytes which are not an HTTP request.
// The HTTP access logs never see these connections, as no request reaches the HTTP routers.
// The http.ConnState hook cannot tell them apart by itself,
// as a connection becomes active as soon as its first byte is read.
type HTTPConnLogger struct {
	logger         *Logger
	entryPointName string

	// served holds, for each tracked connection, whether a request reached the handler.
	served sync.Map
}

// NewHTTPConnLogger creates a new HTTPConnLogger.
// It returns nil when the TCP access logs are disabled, i.e. when the given logger is nil,
// and then tracks no connection.
func NewHTTPConnLogger(logger *Logger, entryPointName string) *HTTPConnLogger {
	if logger == nil {
		return nil
	}

	return &HTTPConnLogger{
		logger:         logger,
		entryPointName: entryPointName,
	}
}

// ConnContext tracks the given connection.
// It is meant to be used by the http.Server ConnContext hook.
func (h *HTTPConnLogger) ConnContext(ctx context.Context, c net.Conn) context.Context {
	if h == nil {
		return ctx
	}

	served := &atomic.Bool{}
	h.served.Store(c, served)

	return context.WithValue(ctx, servedKey{}, served)
}

// ConnState logs the given connection when it is closed without having served any request.
// It is meant to be called by the http.Server ConnState hook.
func (h *HTTPConnLogger) ConnState(c net.Conn, state http.ConnState) {
	if h == nil {
		return
	}

	// A hijacked connection never reports the closed state, and has served a request.
	if state != http.StateClosed && state != http.StateHijacked {
		return
	}

	served, ok := h.served.LoadAndDelete(c)
	if !ok || state == http.StateHijacked || served.(*atomic.Bool).Load() {
		return
	}

	entry := Entry{
		EntryPointName: h.entryPointName,
		ErrorType:      ErrorTypeNoHTTPRequest,
	}

	if tlsConn, ok := c.(*tls.Conn); ok {
		if state := tlsConn.ConnectionState(); !state.HandshakeComplete {
			entry.ServerName = state.ServerName
			entry.ErrorType = ErrorTypeTLSHandshake
		}
	}

	h.logger.LogConnection(c, entry)
}

// Wrap marks the connections of the requests reaching the given handler as served.
// It has to be the first handler executed by the HTTP server.
func (h *HTTPConnLogger) Wrap(next http.Handler) http.Handler {
	if h == nil {
		return next
	}

	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if served, ok := req.Context().Value(servedKey{}).(*atomic.Bool); ok {
			served.Store(true)
		}

		next.ServeHTTP(rw, req)
	})
}
