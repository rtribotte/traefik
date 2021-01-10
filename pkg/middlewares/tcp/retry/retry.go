package tcpretry

import (
	"context"
	"fmt"
	"time"

	"github.com/traefik/traefik/v2/pkg/config/dynamic"
	"github.com/traefik/traefik/v2/pkg/log"
	"github.com/traefik/traefik/v2/pkg/middlewares"
	"github.com/traefik/traefik/v2/pkg/tcp"
)

const (
	typeName = "RetryTCP"
)

// Listener is used to inform about retry attempts.
type Listener interface {
	// Retried will be called when a retry happens.
	// For the first retry this will be attempt 2.
	Retried(attempt int)
}

// Listeners is a convenience type to construct a list of Listener and notify
// each of them about a retry attempt.
type Listeners []Listener

// Retried exists to implement the Listener interface. It calls Retried on each of its slice entries.
func (l Listeners) Retried(attempt int) {
	for _, listener := range l {
		listener.Retried(attempt)
	}
}

// retry is a middleware that retries requests.
type retry struct {
	attempts        int
	initialInterval time.Duration
	next            tcp.Handler
	listener        Listener
	name            string
}

// New returns a new retry middleware.
func New(ctx context.Context, next tcp.Handler, config dynamic.TCPRetry, listener Listener, name string) (tcp.Handler, error) {
	log.FromContext(middlewares.GetLoggerCtx(ctx, name, typeName)).Debug("Creating middleware")

	if config.Attempts <= 0 {
		return nil, fmt.Errorf("incorrect (or empty) value for attempt (%d)", config.Attempts)
	}

	return &retry{
		attempts:        config.Attempts,
		initialInterval: time.Duration(config.InitialInterval),
		next:            next,
		listener:        listener,
		name:            name,
	}, nil
}

func (r *retry) ServeTCP(conn tcp.WriteCloser) {

	retryConn := retryConn{WriteCloser: conn, maxAttempts: r.attempts, attempts: 1}

	for {
		r.next.ServeTCP(&retryConn)

		if !retryConn.shouldRetry() {
			if !retryConn.closed {
				conn.Close()
			}
			return
		}

		retryConn.closed = false
		retryConn.attempts++
		r.listener.Retried(r.attempts)
	}
}

type retryConn struct {
	tcp.WriteCloser
	maxAttempts int
	attempts    int
	closed      bool
	used        bool
}

func (rc retryConn) shouldRetry() bool {
	// Should retry if the conn has been used or closed
	return (rc.used || rc.closed) && rc.attempts < rc.maxAttempts
}

func (rc *retryConn) Close() error {
	rc.closed = true

	// Don't close the underlying conn if it remains some attempts.
	if rc.attempts < rc.maxAttempts {
		return nil
	}

	return rc.WriteCloser.Close()
}

func (rc *retryConn) Read(b []byte) (n int, err error) {
	rc.used = true

	return rc.WriteCloser.Read(b)
}

func (rc *retryConn) Write(b []byte) (n int, err error) {
	rc.used = true

	return rc.WriteCloser.Write(b)
}
