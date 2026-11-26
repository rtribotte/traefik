package tcpaccesslog

import (
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"errors"
	"io"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	traefiktls "github.com/traefik/traefik/v3/pkg/tls"
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
	serverAddressField      = "server.address"
	tlsProtocolVersionField = "tls.protocol.version"
	tlsCipherField          = "tls.cipher"
	tlsClientSubjectField   = "tls.client.subject"
	tlsClientHashField      = "tls.client.hash.sha256"
	errorTypeField          = "error.type"
)

// The values of the error type field, for the connections closed without a router handling them.
const (
	// ErrorTypeTimeout is used when the client did not send its first bytes in time.
	ErrorTypeTimeout = "timeout"
	// ErrorTypeConnectionClosed is used when the client closed the connection before sending its first bytes.
	ErrorTypeConnectionClosed = "connection_closed"
	// ErrorTypeInvalidClientHello is used when the TLS ClientHello could not be read.
	ErrorTypeInvalidClientHello = "invalid_client_hello"
	// ErrorTypeInvalidConnectionData is used when the connection data used for routing could not be read.
	ErrorTypeInvalidConnectionData = "invalid_connection_data"
	// ErrorTypeNoMatchingRouter is used when no router matches the connection.
	ErrorTypeNoMatchingRouter = "no_matching_router"
	// ErrorTypePostgres is used when the Postgres STARTTLS negotiation failed.
	ErrorTypePostgres = "postgres_error"
	// ErrorTypeInvalidTLSOptions is used when the TLS options of the matching router are invalid.
	ErrorTypeInvalidTLSOptions = "invalid_tls_options"
	// ErrorTypeTLSHandshake is used when the TLS handshake of a connection forwarded to the HTTPS server failed.
	ErrorTypeTLSHandshake = "tls_handshake_error"
	// ErrorTypeNoHTTPRequest is used when a connection forwarded to an HTTP server was closed without serving any request.
	ErrorTypeNoHTTPRequest = "no_http_request"
)

// Entry holds what the caller knows about a logged connection.
// The empty values are not logged.
type Entry struct {
	EntryPointName string
	RouterName     string
	ServiceName    string
	TLSMode        string

	// ServerName is only used when the TLS handshake of the connection was not completed,
	// as the server name is then not part of its TLS connection state.
	ServerName string

	// ErrorType is only set for the connections closed without a router handling them.
	// It is replaced when Err shows that the client closed the connection, or did not send its first bytes in time.
	ErrorType string
	Err       error

	// Stats is nil for the connections closed without a router handling them,
	// as these values would only cover the routing, the error type telling more about it.
	Stats *ConnStats
}

// ConnStats holds the bytes exchanged with the client, and the duration of the connection.
type ConnStats struct {
	ReceivedBytes int64
	SentBytes     int64
	Duration      time.Duration
}

// connectionStater is implemented by the connections whose TLS is terminated by Traefik,
// i.e. the *tls.Conn, and the connections wrapping it, such as the metrics middleware ones.
type connectionStater interface {
	ConnectionState() tls.ConnectionState
}

// LogConnection logs the given connection.
// It does nothing when the TCP access logs are disabled, i.e. when the logger is nil.
func (l *Logger) LogConnection(conn net.Conn, entry Entry) {
	if l == nil {
		return
	}

	fields := logrus.Fields{
		entryPointNameField: entry.EntryPointName,
	}

	addField(fields, routerNameField, entry.RouterName)
	addField(fields, serviceNameField, entry.ServiceName)
	addField(fields, tlsModeField, entry.TLSMode)

	if entry.Stats != nil {
		fields[receivedBytesField] = entry.Stats.ReceivedBytes
		fields[sentBytesField] = entry.Stats.SentBytes
		fields[durationField] = entry.Stats.Duration.Seconds()
	}

	addField(fields, errorTypeField, errorType(entry.ErrorType, entry.Err))

	// The listening address is not logged, as the entry point identifies it.
	addAddressFields(fields, clientAddressField, clientPortField, conn.RemoteAddr())

	// The handshake is not forced, as it has to stay driven by the handlers,
	// so the TLS parameters are only known once the connection has been served.
	var state tls.ConnectionState
	if stater, ok := conn.(connectionStater); ok {
		state = stater.ConnectionState()
	}

	if state.HandshakeComplete {
		addTLSFields(fields, state)
	} else {
		addField(fields, serverAddressField, entry.ServerName)
	}

	l.log(fields)
}

func addField(fields logrus.Fields, field, value string) {
	if value != "" {
		fields[field] = value
	}
}

// errorType returns the given error type,
// unless the error shows that the client closed the connection, or did not send its first bytes in time.
func errorType(errorType string, err error) string {
	if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
		return ErrorTypeConnectionClosed
	}

	if netErr, ok := errors.AsType[net.Error](err); ok && netErr.Timeout() {
		return ErrorTypeTimeout
	}

	return errorType
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
	addField(fields, serverAddressField, state.ServerName)
	fields[tlsProtocolVersionField] = traefiktls.GetVersion(&state)
	fields[tlsCipherField] = traefiktls.GetCipherName(&state)

	if len(state.PeerCertificates) == 0 {
		return
	}

	// As for the HTTP access logs, the subject of the client certificate is logged.
	// It is not unique, e.g. for a renewed certificate,
	// so the certificate is also identified by its SHA-256 fingerprint,
	// formatted as an uppercase hash, as the semantic conventions require.
	cert := state.PeerCertificates[0]
	hash := sha256.Sum256(cert.Raw)

	fields[tlsClientSubjectField] = cert.Subject.String()
	fields[tlsClientHashField] = strings.ToUpper(hex.EncodeToString(hash[:]))
}
