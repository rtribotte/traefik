package accesslog

import (
	"bufio"
	"compress/gzip"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	otypes "github.com/traefik/traefik/v3/pkg/observability/types"
	"github.com/traefik/traefik/v3/pkg/tcp"
	"github.com/traefik/traefik/v3/pkg/tls/generate"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/plog/plogotlp"
)

func TestRouterAccessLogHandler(t *testing.T) {
	testCases := []struct {
		desc    string
		logger  *Logger
		tls     *dynamic.RouterTCPTLSConfig
		wrapped bool
		tlsMode string
	}{
		{
			desc: "TCP access logs disabled",
		},
		{
			desc:    "without TLS",
			logger:  &Logger{},
			wrapped: true,
			tlsMode: "none",
		},
		{
			desc:    "TLS terminated",
			logger:  &Logger{},
			tls:     &dynamic.RouterTCPTLSConfig{},
			wrapped: true,
			tlsMode: "terminated",
		},
		{
			desc:    "TLS passthrough",
			logger:  &Logger{},
			tls:     &dynamic.RouterTCPTLSConfig{Passthrough: true},
			wrapped: true,
			tlsMode: "passthrough",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			next := tcp.HandlerFunc(func(conn tcp.WriteCloser) {})

			handler, err := RouterAccessLogHandler(t.Context(), test.logger, "web", "router@file", "service@file", test.tls)(next)
			require.NoError(t, err)

			middleware, wrapped := handler.(*accessLogMiddleware)
			require.Equal(t, test.wrapped, wrapped)

			if wrapped {
				assert.Equal(t, test.tlsMode, middleware.tlsMode)
			}
		})
	}
}

func TestAccessLogMiddleware_ServeTCP(t *testing.T) {
	logger, logFile := newLogger(t)

	serverConn, clientConn := net.Pipe()
	t.Cleanup(func() {
		_ = serverConn.Close()
		_ = clientConn.Close()
	})

	go func() {
		_, _ = clientConn.Write([]byte("ping"))
		_, _ = io.ReadFull(clientConn, make([]byte, 5))
	}()

	next := tcp.HandlerFunc(func(conn tcp.WriteCloser) {
		_, err := io.ReadFull(conn, make([]byte, 4))
		assert.NoError(t, err)

		_, err = conn.Write([]byte("pong!"))
		assert.NoError(t, err)
	})

	middleware, err := RouterAccessLogHandler(t.Context(), logger, "web", "router@file", "service@file", nil)(next)
	require.NoError(t, err)

	middleware.ServeTCP(pipeConn{Conn: serverConn})

	entries := readEntries(t, logFile)
	require.Len(t, entries, 1)

	entry := entries[0]
	assert.Equal(t, "web", entry["traefik.entrypoint.name"])
	assert.Equal(t, "router@file", entry["traefik.router.name"])
	assert.Equal(t, "service@file", entry["traefik.service.name"])
	assert.Equal(t, "none", entry["traefik.tls.mode"])
	assert.InDelta(t, 4, entry["traefik.connection.received_bytes"], 0)
	assert.InDelta(t, 5, entry["traefik.connection.sent_bytes"], 0)
	assert.Contains(t, entry, "traefik.connection.duration")
	assert.NotEmpty(t, entry["time"])

	// The address of a pipe cannot be split into a host and a port.
	assert.Equal(t, "pipe", entry["client.address"])
	assert.NotContains(t, entry, "client.port")

	assert.NotContains(t, entry, "tls.protocol.version")
}

func TestAccessLogMiddleware_ServeTCP_tls(t *testing.T) {
	cert, err := generate.DefaultCertificate()
	require.NoError(t, err)

	leaf, err := x509.ParseCertificate(cert.Certificate[0])
	require.NoError(t, err)

	testCases := []struct {
		desc            string
		clientCert      bool
		expectedSubject string
		expectedHash    string
	}{
		{
			desc: "without client certificate",
		},
		{
			desc:            "with client certificate",
			clientCert:      true,
			expectedSubject: leaf.Subject.String(),
			expectedHash:    fmt.Sprintf("%X", sha256.Sum256(leaf.Raw)),
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			logger, logFile := newLogger(t)

			serverConn, clientConn := net.Pipe()
			t.Cleanup(func() {
				_ = serverConn.Close()
				_ = clientConn.Close()
			})

			clientConfig := &tls.Config{ServerName: "foo.localhost", InsecureSkipVerify: true}
			if test.clientCert {
				clientConfig.Certificates = []tls.Certificate{*cert}
			}

			go func() {
				client := tls.Client(clientConn, clientConfig)
				_, _ = client.Write([]byte("ping"))
			}()

			next := tcp.HandlerFunc(func(conn tcp.WriteCloser) {
				_, err := io.ReadFull(conn, make([]byte, 4))
				assert.NoError(t, err)
			})

			middleware, err := RouterAccessLogHandler(t.Context(), logger, "websecure", "router@file", "service@file", &dynamic.RouterTCPTLSConfig{})(next)
			require.NoError(t, err)

			middleware.ServeTCP(tls.Server(serverConn, &tls.Config{
				Certificates: []tls.Certificate{*cert},
				MaxVersion:   tls.VersionTLS12,
				CipherSuites: []uint16{tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256},
				ClientAuth:   tls.RequestClientCert,
			}))

			entries := readEntries(t, logFile)
			require.Len(t, entries, 1)

			entry := entries[0]
			assert.Equal(t, "terminated", entry["traefik.tls.mode"])
			assert.Equal(t, "foo.localhost", entry["server.address"])
			assert.Equal(t, "1.2", entry["tls.protocol.version"])
			assert.Equal(t, "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256", entry["tls.cipher"])
			assert.InDelta(t, 4, entry["traefik.connection.received_bytes"], 0)

			if test.expectedSubject == "" {
				assert.NotContains(t, entry, "tls.client.subject")
				assert.NotContains(t, entry, "tls.client.hash.sha256")
				return
			}

			assert.Equal(t, test.expectedSubject, entry["tls.client.subject"])
			assert.Equal(t, test.expectedHash, entry["tls.client.hash.sha256"])
		})
	}
}

func TestAddAddressFields(t *testing.T) {
	testCases := []struct {
		desc     string
		addr     net.Addr
		expected map[string]any
	}{
		{
			desc:     "IPv4 TCP address",
			addr:     &net.TCPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 4242},
			expected: map[string]any{"address": "10.0.0.1", "port": 4242},
		},
		{
			desc:     "IPv6 TCP address",
			addr:     &net.TCPAddr{IP: net.ParseIP("::1"), Port: 4242},
			expected: map[string]any{"address": "::1", "port": 4242},
		},
		{
			desc:     "splittable address",
			addr:     stringAddr("192.168.1.1:8080"),
			expected: map[string]any{"address": "192.168.1.1", "port": 8080},
		},
		{
			desc:     "unsplittable address",
			addr:     stringAddr("/run/traefik.sock"),
			expected: map[string]any{"address": "/run/traefik.sock"},
		},
		{
			desc:     "no address",
			expected: map[string]any{},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			fields := make(map[string]any)
			addAddressFields(fields, "address", "port", test.addr)

			assert.Equal(t, test.expected, fields)
		})
	}
}

func TestAccessLogMiddleware_ServeTCP_otlp(t *testing.T) {
	testCases := []struct {
		desc        string
		dualOutput  bool
		fileEntries int
	}{
		{
			desc: "OTLP only",
		},
		{
			desc:        "OTLP with dual output",
			dualOutput:  true,
			fileEntries: 1,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			collectorURL, exported := newCollector(t)

			logFile := filepath.Join(t.TempDir(), "tcp-access.log")

			logger, err := NewLogger(t.Context(), &otypes.TCPAccessLog{
				FilePath:   logFile,
				DualOutput: test.dualOutput,
				OTLP: &otypes.OTelLog{
					ServiceName: "test",
					HTTP:        &otypes.OTelHTTP{Endpoint: collectorURL},
				},
			})
			require.NoError(t, err)
			t.Cleanup(func() { require.NoError(t, logger.Close()) })

			serverConn, clientConn := net.Pipe()
			t.Cleanup(func() {
				_ = serverConn.Close()
				_ = clientConn.Close()
			})

			go func() {
				_, _ = clientConn.Write([]byte("ping"))
			}()

			next := tcp.HandlerFunc(func(conn tcp.WriteCloser) {
				_, err := io.ReadFull(conn, make([]byte, 4))
				assert.NoError(t, err)
			})

			middleware, err := RouterAccessLogHandler(t.Context(), logger, "web", "router@file", "service@file", nil)(next)
			require.NoError(t, err)

			middleware.ServeTCP(pipeConn{Conn: serverConn})

			resource, record := waitForRecord(t, exported)

			serviceName, _ := resource.Attributes().Get("service.name")
			assert.Equal(t, "test", serviceName.Str())

			routerName, _ := record.Attributes().Get("traefik.router.name")
			assert.Equal(t, "router@file", routerName.Str())

			receivedBytes, _ := record.Attributes().Get("traefik.connection.received_bytes")
			assert.Equal(t, int64(4), receivedBytes.Int())

			// The body is the formatted log line, as written to the file with the dual output.
			body := make(map[string]any)
			require.NoError(t, json.Unmarshal([]byte(record.Body().Str()), &body))
			assert.Equal(t, "router@file", body["traefik.router.name"])
			assert.Empty(t, body["msg"])

			entries := readEntries(t, logFile)
			require.Len(t, entries, test.fileEntries)

			if test.dualOutput {
				// The written log line does not embed the body in its message.
				assert.Empty(t, entries[0]["msg"])
				assert.Equal(t, body, entries[0])
			}
		})
	}
}

func TestLogger_export_formatError(t *testing.T) {
	collectorURL, exported := newCollector(t)

	logger, err := NewLogger(t.Context(), &otypes.TCPAccessLog{
		OTLP: &otypes.OTelLog{
			HTTP: &otypes.OTelHTTP{Endpoint: collectorURL},
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, logger.Close()) })

	// A channel cannot be serialized to JSON, which makes the formatting fail.
	logger.log(logrus.Fields{"traefik.router.name": "router@file", "unserializable": make(chan int)})

	_, record := waitForRecord(t, exported)

	// The log line is exported with its attributes, but the formatting error is not sent as its body.
	routerName, _ := record.Attributes().Get("traefik.router.name")
	assert.Equal(t, "router@file", routerName.Str())
	assert.Equal(t, "{}", record.Body().Str())
}

// newCollector starts an OTLP HTTP collector, which sends the export requests it receives to the returned channel.
func newCollector(t *testing.T) (string, <-chan plogotlp.ExportRequest) {
	t.Helper()

	exported := make(chan plogotlp.ExportRequest, 10)
	collector := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gzr, err := gzip.NewReader(r.Body)
		if !assert.NoError(t, err) {
			return
		}

		body, err := io.ReadAll(gzr)
		if !assert.NoError(t, err) {
			return
		}

		req := plogotlp.NewExportRequest()
		if !assert.NoError(t, req.UnmarshalProto(body)) {
			return
		}

		exported <- req
	}))
	t.Cleanup(collector.Close)

	return collector.URL, exported
}

// waitForRecord waits for the first exported log record, and returns it with its resource.
func waitForRecord(t *testing.T, exported <-chan plogotlp.ExportRequest) (pcommon.Resource, plog.LogRecord) {
	t.Helper()

	select {
	case <-time.After(5 * time.Second):
		t.Fatal("TCP access log not exported")
		return pcommon.Resource{}, plog.LogRecord{}

	case req := <-exported:
		require.Positive(t, req.Logs().LogRecordCount())

		resourceLogs := req.Logs().ResourceLogs().At(0)

		return resourceLogs.Resource(), resourceLogs.ScopeLogs().At(0).LogRecords().At(0)
	}
}

func newLogger(t *testing.T) (*Logger, string) {
	t.Helper()

	logFile := filepath.Join(t.TempDir(), "tcp-access.log")

	logger, err := NewLogger(t.Context(), &otypes.TCPAccessLog{FilePath: logFile})
	require.NoError(t, err)

	t.Cleanup(func() { require.NoError(t, logger.Close()) })

	return logger, logFile
}

func readEntries(t *testing.T, logFile string) []map[string]any {
	t.Helper()

	file, err := os.Open(logFile)
	require.NoError(t, err)
	defer file.Close()

	var entries []map[string]any
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry := make(map[string]any)
		require.NoError(t, json.Unmarshal(scanner.Bytes(), &entry))

		entries = append(entries, entry)
	}
	require.NoError(t, scanner.Err())

	return entries
}

// pipeConn adds the CloseWrite method required by tcp.WriteCloser to a net.Pipe connection.
type pipeConn struct {
	net.Conn
}

func (p pipeConn) CloseWrite() error {
	return p.Close()
}

type stringAddr string

func (s stringAddr) Network() string { return "tcp" }

func (s stringAddr) String() string { return string(s) }
