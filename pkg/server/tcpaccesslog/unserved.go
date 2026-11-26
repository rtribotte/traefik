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

// UnservedConnLogger logs the connections an HTTP server closes without serving any request,
// e.g. because of a failed TLS handshake or bytes which are not an HTTP request.
// The HTTP access logs never see these connections, as no request reaches the HTTP routers.
// The http.ConnState hook cannot tell them apart by itself,
// as a connection becomes active as soon as its first byte is read.
type UnservedConnLogger struct {
	logger         *Logger
	entryPointName string

	// served holds, for each tracked connection, whether a request reached the handler.
	served sync.Map
}

// NewUnservedConnLogger creates a new UnservedConnLogger.
// It returns nil when the TCP access logs are disabled, i.e. when the given logger is nil,
// and then tracks no connection.
func NewUnservedConnLogger(logger *Logger, entryPointName string) *UnservedConnLogger {
	if logger == nil {
		return nil
	}

	return &UnservedConnLogger{
		logger:         logger,
		entryPointName: entryPointName,
	}
}

// ConnContext tracks the given connection.
// It is meant to be used by the http.Server ConnContext hook.
func (h *UnservedConnLogger) ConnContext(ctx context.Context, c net.Conn) context.Context {
	if h == nil {
		return ctx
	}

	served := &atomic.Bool{}
	h.served.Store(c, served)

	return context.WithValue(ctx, servedKey{}, served)
}

// ConnState logs the given connection when it is closed without having served any request.
// It is meant to be called by the http.Server ConnState hook.
func (h *UnservedConnLogger) ConnState(c net.Conn, state http.ConnState) {
	if h == nil {
		return
	}

	// A hijacked connection never reports the closed state, and has served a request.
	if state != http.StateClosed && state != http.StateHijacked {
		return
	}

	// The connection stops being tracked here, so that it is logged at most once,
	// even if the closed state were reported again.
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

// MarkServed marks the connection of the request with the given context as served,
// when it is tracked by an UnservedConnLogger.
func MarkServed(ctx context.Context) {
	if served, ok := ctx.Value(servedKey{}).(*atomic.Bool); ok {
		served.Store(true)
	}
}
