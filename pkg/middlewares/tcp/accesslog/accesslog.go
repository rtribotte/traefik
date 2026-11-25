package accesslog

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/pem"
	"maps"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/traefik/traefik/v3/pkg/middlewares"
	"github.com/traefik/traefik/v3/pkg/tcp"
	traefiktls "github.com/traefik/traefik/v3/pkg/tls"
)

const (
	typeName   = "AccessLogTCP"
	nameRouter = "accesslog-router"
)

// connectionStater is implemented by the connections whose TLS is terminated by Traefik,
// i.e. the *tls.Conn, and the connections wrapping it, such as the metrics middleware ones.
type connectionStater interface {
	ConnectionState() tls.ConnectionState
}

type accessLogMiddleware struct {
	next    tcp.Handler
	handler *Handler

	entryPointName string
	routerName     string
	serviceName    string
}

// RouterAccessLogHandler returns the access log TCP router handler.
// It is idempotent when the TCP access logs are disabled, i.e. when the given handler is nil.
func RouterAccessLogHandler(ctx context.Context, handler *Handler, entryPointName, routerName, serviceName string) tcp.Constructor {
	return func(next tcp.Handler) (tcp.Handler, error) {
		if handler == nil {
			return next, nil
		}

		middlewares.GetLogger(ctx, nameRouter, typeName).Debug().Msg("Creating middleware")

		return &accessLogMiddleware{
			next:           next,
			handler:        handler,
			entryPointName: entryPointName,
			routerName:     routerName,
			serviceName:    serviceName,
		}, nil
	}
}

// ServeTCP logs the start and the end of the given connection.
func (a *accessLogMiddleware) ServeTCP(conn tcp.WriteCloser) {
	start := time.Now()

	fields := logrus.Fields{
		"entry_point_name": a.entryPointName,
		"router_name":      a.routerName,
		"service_name":     a.serviceName,
		"client_addr":      conn.RemoteAddr().String(),
		"server_addr":      conn.LocalAddr().String(),
	}

	a.handler.log(withFields(fields, logrus.Fields{
		"event":     "connection_start",
		"timestamp": start.Format(time.RFC3339Nano),
	}))

	cConn := &countingConn{WriteCloser: conn}
	a.next.ServeTCP(cConn)

	end := withFields(fields, logrus.Fields{
		"event":       "connection_end",
		"timestamp":   time.Now().Format(time.RFC3339Nano),
		"bytes_in":    cConn.bytesRead,
		"bytes_out":   cConn.bytesWritten,
		"duration_ms": time.Since(start).Milliseconds(),
		"status":      "success",
	})

	// The handshake is not forced, as it has to stay driven by the next handlers,
	// so the TLS parameters are only known once the connection has been served.
	if stater, ok := conn.(connectionStater); ok {
		addTLSFields(end, stater.ConnectionState())
	}

	a.handler.log(end)
}

func withFields(fields, extra logrus.Fields) logrus.Fields {
	result := make(logrus.Fields, len(fields)+len(extra))
	maps.Copy(result, fields)
	maps.Copy(result, extra)

	return result
}

func addTLSFields(fields logrus.Fields, state tls.ConnectionState) {
	if !state.HandshakeComplete {
		return
	}

	fields["tls_server_name"] = state.ServerName
	fields["tls_version"] = traefiktls.GetVersion(&state)
	fields["tls_cipher"] = traefiktls.GetCipherName(&state)

	if len(state.PeerCertificates) == 0 {
		return
	}

	buf := new(bytes.Buffer)
	if err := pem.Encode(buf, &pem.Block{Type: "CERTIFICATE", Bytes: state.PeerCertificates[0].Raw}); err == nil {
		fields["client_cert_pem"] = buf.String()
	}
}

// countingConn counts the bytes read from and written to the client.
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
