package accesslog

import (
	"net"
	"time"

	"github.com/sirupsen/logrus"
)

// The values of the error type field, for the connections closed by an entry point without a router handling them.
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
)

// EntryPointLogger logs the connections an entry point closes without a router handling them,
// as the access log middleware of the routers never sees them.
type EntryPointLogger struct {
	logger         *Logger
	entryPointName string
}

// NewEntryPointLogger creates a new EntryPointLogger.
// It returns nil when the TCP access logs are disabled, i.e. when the given logger is nil.
func NewEntryPointLogger(logger *Logger, entryPointName string) *EntryPointLogger {
	if logger == nil {
		return nil
	}

	return &EntryPointLogger{
		logger:         logger,
		entryPointName: entryPointName,
	}
}

// LogUnhandled logs a connection closed by the entry point without a router handling it.
// The router name is only known when the matching router has invalid TLS options,
// and the server name only when the TLS ClientHello was read.
func (e *EntryPointLogger) LogUnhandled(conn net.Conn, start time.Time, receivedBytes int64, routerName, serverName, errorType string) {
	if e == nil {
		return
	}

	fields := logrus.Fields{
		entryPointNameField: e.entryPointName,
		receivedBytesField:  receivedBytes,
		sentBytesField:      int64(0),
		durationField:       time.Since(start).Seconds(),
		errorTypeField:      errorType,
	}

	if routerName != "" {
		fields[routerNameField] = routerName
	}

	if serverName != "" {
		fields[serverAddressField] = serverName
	}

	addAddressFields(fields, clientAddressField, clientPortField, conn.RemoteAddr())

	e.logger.log(fields)
}
