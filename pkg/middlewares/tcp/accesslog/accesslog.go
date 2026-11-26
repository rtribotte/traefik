package accesslog

import (
	"context"
	"time"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/middlewares"
	"github.com/traefik/traefik/v3/pkg/server/tcpaccesslog"
	"github.com/traefik/traefik/v3/pkg/tcp"
)

const (
	typeName   = "AccessLogTCP"
	nameRouter = "accesslog-router"
)

type accessLogMiddleware struct {
	next   tcp.Handler
	logger *tcpaccesslog.Logger
	entry  tcpaccesslog.Entry
}

// RouterAccessLogHandler returns the access log TCP router handler.
// It is a no-op when the TCP access logs are disabled, i.e. when the given logger is nil.
func RouterAccessLogHandler(ctx context.Context, logger *tcpaccesslog.Logger, entryPointName, routerName, serviceName string, tlsConfig *dynamic.RouterTCPTLSConfig) tcp.Constructor {
	return func(next tcp.Handler) (tcp.Handler, error) {
		if logger == nil {
			return next, nil
		}

		middlewares.GetLogger(ctx, nameRouter, typeName).Debug().Msg("Creating middleware")

		tlsMode := tcpaccesslog.TLSModeNone
		if tlsConfig != nil {
			if tlsConfig.Passthrough {
				tlsMode = tcpaccesslog.TLSModePassthrough
			} else {
				tlsMode = tcpaccesslog.TLSModeTerminated
			}
		}

		return &accessLogMiddleware{
			next:   next,
			logger: logger,
			entry: tcpaccesslog.Entry{
				EntryPointName: entryPointName,
				RouterName:     routerName,
				ServiceName:    serviceName,
				TLSMode:        tlsMode,
			},
		}, nil
	}
}

// ServeTCP logs the given connection once it has been served.
func (a *accessLogMiddleware) ServeTCP(conn tcp.WriteCloser) {
	start := time.Now()

	cConn := &countingConn{WriteCloser: conn}
	a.next.ServeTCP(cConn)

	entry := a.entry
	entry.Stats = &tcpaccesslog.ConnStats{
		ReceivedBytes: cConn.bytesRead,
		SentBytes:     cConn.bytesWritten,
		Duration:      time.Since(start),
	}

	a.logger.LogConnection(conn, entry)
}

// countingConn counts the bytes received from and sent to the client.
// The counters are only read once the connection has been served,
// and each one is only updated by the goroutine copying in its direction.
type countingConn struct {
	tcp.WriteCloser

	bytesRead    int64
	bytesWritten int64
}

func (c *countingConn) Read(b []byte) (int, error) {
	n, err := c.WriteCloser.Read(b)
	c.bytesRead += int64(n)

	return n, err
}

func (c *countingConn) Write(b []byte) (int, error) {
	n, err := c.WriteCloser.Write(b)
	c.bytesWritten += int64(n)

	return n, err
}
