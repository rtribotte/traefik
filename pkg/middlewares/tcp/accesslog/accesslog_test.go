package accesslog

import (
	"bufio"
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	otypes "github.com/traefik/traefik/v3/pkg/observability/types"
	"github.com/traefik/traefik/v3/pkg/tcp"
	"github.com/traefik/traefik/v3/pkg/tls/generate"
	"go.opentelemetry.io/collector/pdata/plog/plogotlp"
)

func TestRouterAccessLogHandler(t *testing.T) {
	testCases := []struct {
		desc    string
		handler *Handler
		tls     *dynamic.RouterTCPTLSConfig
		wrapped bool
		tlsMode string
	}{
		{
			desc: "TCP access logs disabled",
		},
		{
			desc:    "without TLS",
			handler: &Handler{},
			wrapped: true,
			tlsMode: "none",
		},
		{
			desc:    "TLS terminated",
			handler: &Handler{},
			tls:     &dynamic.RouterTCPTLSConfig{},
			wrapped: true,
			tlsMode: "terminated",
		},
		{
			desc:    "TLS passthrough",
			handler: &Handler{},
			tls:     &dynamic.RouterTCPTLSConfig{Passthrough: true},
			wrapped: true,
			tlsMode: "passthrough",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			next := tcp.HandlerFunc(func(conn tcp.WriteCloser) {})

			handler, err := RouterAccessLogHandler(t.Context(), test.handler, "web", "router@file", "service@file", test.tls)(next)
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
	handler, logFile := newHandler(t)

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

	middleware, err := RouterAccessLogHandler(t.Context(), handler, "web", "router@file", "service@file", nil)(next)
	require.NoError(t, err)

	middleware.ServeTCP(pipeConn{Conn: serverConn})

	entries := readEntries(t, logFile)
	require.Len(t, entries, 1)

	entry := entries[0]
	assert.Equal(t, "web", entry["traefik.entrypoint.name"])
	assert.Equal(t, "router@file", entry["traefik.router.name"])
	assert.Equal(t, "service@file", entry["traefik.service.name"])
	assert.Equal(t, "none", entry["traefik.tls.mode"])
	assert.Equal(t, "tcp", entry["network.transport"])
	assert.InDelta(t, 4, entry["traefik.connection.received_bytes"], 0)
	assert.InDelta(t, 5, entry["traefik.connection.sent_bytes"], 0)
	assert.Contains(t, entry, "traefik.connection.duration")
	assert.NotEmpty(t, entry["time"])

	// The address of a pipe cannot be split into a host and a port, and is not a TCP one.
	assert.Equal(t, "pipe", entry["client.address"])
	assert.NotContains(t, entry, "client.port")
	assert.Equal(t, "pipe", entry["network.local.address"])
	assert.NotContains(t, entry, "network.local.port")
	assert.NotContains(t, entry, "network.type")

	assert.NotContains(t, entry, "tls.protocol.version")
}

func TestAccessLogMiddleware_ServeTCP_tls(t *testing.T) {
	handler, logFile := newHandler(t)

	cert, err := generate.DefaultCertificate()
	require.NoError(t, err)

	serverConn, clientConn := net.Pipe()
	t.Cleanup(func() {
		_ = serverConn.Close()
		_ = clientConn.Close()
	})

	go func() {
		client := tls.Client(clientConn, &tls.Config{ServerName: "foo.localhost", InsecureSkipVerify: true})
		_, _ = client.Write([]byte("ping"))
	}()

	next := tcp.HandlerFunc(func(conn tcp.WriteCloser) {
		_, err := io.ReadFull(conn, make([]byte, 4))
		assert.NoError(t, err)
	})

	middleware, err := RouterAccessLogHandler(t.Context(), handler, "websecure", "router@file", "service@file", &dynamic.RouterTCPTLSConfig{})(next)
	require.NoError(t, err)

	middleware.ServeTCP(tls.Server(serverConn, &tls.Config{
		Certificates: []tls.Certificate{*cert},
		MaxVersion:   tls.VersionTLS12,
		CipherSuites: []uint16{tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256},
	}))

	entries := readEntries(t, logFile)
	require.Len(t, entries, 1)

	entry := entries[0]
	assert.Equal(t, "terminated", entry["traefik.tls.mode"])
	assert.Equal(t, "foo.localhost", entry["server.address"])
	assert.Equal(t, "1.2", entry["tls.protocol.version"])
	assert.Equal(t, "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256", entry["tls.cipher"])
	assert.InDelta(t, 4, entry["traefik.connection.received_bytes"], 0)
	assert.NotContains(t, entry, "tls.client.certificate")
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

			logCh := make(chan string, 1)
			collector := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gzr, err := gzip.NewReader(r.Body)
				require.NoError(t, err)

				body, err := io.ReadAll(gzr)
				require.NoError(t, err)

				req := plogotlp.NewExportRequest()
				require.NoError(t, req.UnmarshalProto(body))

				marshalledReq, err := json.Marshal(req)
				require.NoError(t, err)

				logCh <- string(marshalledReq)
			}))
			t.Cleanup(collector.Close)

			logFile := filepath.Join(t.TempDir(), "tcp-access.log")

			handler, err := NewHandler(t.Context(), &otypes.TCPAccessLog{
				FilePath:   logFile,
				DualOutput: test.dualOutput,
				OTLP: &otypes.OTelLog{
					ServiceName: "test",
					HTTP:        &otypes.OTelHTTP{Endpoint: collector.URL},
				},
			})
			require.NoError(t, err)
			t.Cleanup(func() { _ = handler.file.Close() })

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

			middleware, err := RouterAccessLogHandler(t.Context(), handler, "web", "router@file", "service@file", nil)(next)
			require.NoError(t, err)

			middleware.ServeTCP(pipeConn{Conn: serverConn})

			select {
			case <-time.After(5 * time.Second):
				t.Fatal("TCP access log not exported")

			case log := <-logCh:
				assert.Regexp(t, `{"key":"service.name","value":{"stringValue":"test"}}`, log)
				assert.Regexp(t, `{"key":"traefik.router.name","value":{"stringValue":"router@file"}}`, log)
				assert.Regexp(t, `{"key":"traefik.connection.received_bytes","value":{"intValue":"4"}}`, log)
				assert.NotRegexp(t, `"body":{"stringValue":""}`, log)
			}

			assert.Len(t, readEntries(t, logFile), test.fileEntries)
		})
	}
}

func newHandler(t *testing.T) (*Handler, string) {
	t.Helper()

	logFile := filepath.Join(t.TempDir(), "tcp-access.log")

	handler, err := NewHandler(t.Context(), &otypes.TCPAccessLog{FilePath: logFile})
	require.NoError(t, err)

	t.Cleanup(func() { _ = handler.file.Close() })

	return handler, logFile
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
