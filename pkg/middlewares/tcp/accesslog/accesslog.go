package accesslog

import (
	"context"
	"crypto/tls"
	"net"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/middlewares"
	"github.com/traefik/traefik/v3/pkg/tcp"
	traefiktls "github.com/traefik/traefik/v3/pkg/tls"
)

const (
	typeName   = "AccessLogTCP"
	nameRouter = "accesslog-router"
)

// The fields are named after the OpenTelemetry semantic conventions attributes,
// and after the TCP routers metrics labels for the Traefik specific ones.
const (
	entryPointNameField     = "traefik.entrypoint.name"
	routerNameField         = "traefik.router.name"
	serviceNameField        = "traefik.service.name"
	tlsModeField            = "traefik.tls.mode"
	receivedBytesField      = "traefik.connection.received_bytes"
	sentBytesField          = "traefik.connection.sent_bytes"
	durationField           = "traefik.connection.duration"
	clientAddressField      = "client.address"
	clientPortField         = "client.port"
	localAddressField       = "network.local.address"
	localPortField          = "network.local.port"
	networkTransportField   = "network.transport"
	networkTypeField        = "network.type"
	serverAddressField      = "server.address"
	tlsProtocolVersionField = "tls.protocol.version"
	tlsCipherField          = "tls.cipher"
	tlsClientSubjectField   = "tls.client.subject"
)

// The values of the TLS mode field.
const (
	tlsModeNone        = "none"
	tlsModeTerminated  = "terminated"
	tlsModePassthrough = "passthrough"
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
	tlsMode        string
}

// RouterAccessLogHandler returns the access log TCP router handler.
// It is idempotent when the TCP access logs are disabled, i.e. when the given handler is nil.
func RouterAccessLogHandler(ctx context.Context, handler *Handler, entryPointName, routerName, serviceName string, tlsConfig *dynamic.RouterTCPTLSConfig) tcp.Constructor {
	return func(next tcp.Handler) (tcp.Handler, error) {
		if handler == nil {
			return next, nil
		}

		middlewares.GetLogger(ctx, nameRouter, typeName).Debug().Msg("Creating middleware")

		tlsMode := tlsModeNone
		if tlsConfig != nil {
			if tlsConfig.Passthrough {
				tlsMode = tlsModePassthrough
			} else {
				tlsMode = tlsModeTerminated
			}
		}

		return &accessLogMiddleware{
			next:           next,
			handler:        handler,
			entryPointName: entryPointName,
			routerName:     routerName,
			serviceName:    serviceName,
			tlsMode:        tlsMode,
		}, nil
	}
}

// ServeTCP logs the given connection once it has been served.
func (a *accessLogMiddleware) ServeTCP(conn tcp.WriteCloser) {
	start := time.Now()

	cConn := &countingConn{WriteCloser: conn}
	a.next.ServeTCP(cConn)

	fields := logrus.Fields{
		entryPointNameField:   a.entryPointName,
		routerNameField:       a.routerName,
		serviceNameField:      a.serviceName,
		tlsModeField:          a.tlsMode,
		networkTransportField: "tcp",
		receivedBytesField:    cConn.bytesRead,
		sentBytesField:        cConn.bytesWritten,
		durationField:         time.Since(start).Seconds(),
	}

	addAddressFields(fields, clientAddressField, clientPortField, conn.RemoteAddr())
	addAddressFields(fields, localAddressField, localPortField, conn.LocalAddr())

	if addr, ok := conn.LocalAddr().(*net.TCPAddr); ok {
		fields[networkTypeField] = "ipv6"
		if addr.IP.To4() != nil {
			fields[networkTypeField] = "ipv4"
		}
	}

	// The handshake is not forced, as it has to stay driven by the next handlers,
	// so the TLS parameters are only known once the connection has been served.
	if stater, ok := conn.(connectionStater); ok {
		addTLSFields(fields, stater.ConnectionState())
	}

	a.handler.log(fields)
}

// addAddressFields adds the address and the port of the given network address.
// The whole address is used when it cannot be split, e.g. for a Unix socket.
func addAddressFields(fields logrus.Fields, addressField, portField string, addr net.Addr) {
	if addr == nil {
		return
	}

	if tcpAddr, ok := addr.(*net.TCPAddr); ok {
		fields[addressField] = tcpAddr.IP.String()
		fields[portField] = tcpAddr.Port
		return
	}

	host, port, err := net.SplitHostPort(addr.String())
	if err != nil {
		fields[addressField] = addr.String()
		return
	}

	fields[addressField] = host
	if p, err := strconv.Atoi(port); err == nil {
		fields[portField] = p
	}
}

func addTLSFields(fields logrus.Fields, state tls.ConnectionState) {
	if !state.HandshakeComplete {
		return
	}

	if state.ServerName != "" {
		fields[serverAddressField] = state.ServerName
	}
	fields[tlsProtocolVersionField] = traefiktls.GetVersion(&state)
	fields[tlsCipherField] = traefiktls.GetCipherName(&state)

	// As for the HTTP access logs, only the subject of the client certificate is logged.
	if len(state.PeerCertificates) > 0 {
		fields[tlsClientSubjectField] = state.PeerCertificates[0].Subject.String()
	}
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
