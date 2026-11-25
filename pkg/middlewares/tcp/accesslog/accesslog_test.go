package accesslog

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"io"
	"net"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	otypes "github.com/traefik/traefik/v3/pkg/observability/types"
	"github.com/traefik/traefik/v3/pkg/tcp"
	"github.com/traefik/traefik/v3/pkg/tls/generate"
)

func TestRouterAccessLogHandler(t *testing.T) {
	testCases := []struct {
		desc    string
		handler *Handler
		wrapped bool
	}{
		{
			desc: "TCP access logs disabled",
		},
		{
			desc:    "TCP access logs enabled",
			handler: &Handler{},
			wrapped: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			next := tcp.HandlerFunc(func(conn tcp.WriteCloser) {})

			handler, err := RouterAccessLogHandler(t.Context(), test.handler, "web", "router@file", "service@file")(next)
			require.NoError(t, err)

			_, wrapped := handler.(*accessLogMiddleware)
			assert.Equal(t, test.wrapped, wrapped)
		})
	}
}

func TestAccessLogMiddleware_ServeTCP(t *testing.T) {
	handler, logFile := newJSONHandler(t)

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

	middleware, err := RouterAccessLogHandler(t.Context(), handler, "web", "router@file", "service@file")(next)
	require.NoError(t, err)

	middleware.ServeTCP(pipeConn{Conn: serverConn})

	entries := readEntries(t, logFile)
	require.Len(t, entries, 2)

	for _, entry := range entries {
		assert.Equal(t, "web", entry["entry_point_name"])
		assert.Equal(t, "router@file", entry["router_name"])
		assert.Equal(t, "service@file", entry["service_name"])
		assert.Equal(t, "pipe", entry["client_addr"])
		assert.Equal(t, "pipe", entry["server_addr"])
		assert.NotEmpty(t, entry["timestamp"])
	}

	assert.Equal(t, "connection_start", entries[0]["event"])
	assert.NotContains(t, entries[0], "bytes_in")

	assert.Equal(t, "connection_end", entries[1]["event"])
	assert.InDelta(t, 4, entries[1]["bytes_in"], 0)
	assert.InDelta(t, 5, entries[1]["bytes_out"], 0)
	assert.Contains(t, entries[1], "duration_ms")
	assert.Equal(t, "success", entries[1]["status"])
	assert.NotContains(t, entries[1], "tls_version")
}

func TestAccessLogMiddleware_ServeTCP_tls(t *testing.T) {
	handler, logFile := newJSONHandler(t)

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

	middleware, err := RouterAccessLogHandler(t.Context(), handler, "websecure", "router@file", "service@file")(next)
	require.NoError(t, err)

	middleware.ServeTCP(tls.Server(serverConn, &tls.Config{
		Certificates: []tls.Certificate{*cert},
		MaxVersion:   tls.VersionTLS12,
		CipherSuites: []uint16{tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256},
	}))

	entries := readEntries(t, logFile)
	require.Len(t, entries, 2)

	// The TLS parameters are not known when the connection starts, as the handshake is not forced.
	assert.NotContains(t, entries[0], "tls_version")

	assert.Equal(t, "foo.localhost", entries[1]["tls_server_name"])
	assert.Equal(t, "1.2", entries[1]["tls_version"])
	assert.Equal(t, "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256", entries[1]["tls_cipher"])
	assert.InDelta(t, 4, entries[1]["bytes_in"], 0)
	assert.NotContains(t, entries[1], "client_cert_pem")
}

func newJSONHandler(t *testing.T) (*Handler, string) {
	t.Helper()

	logFile := filepath.Join(t.TempDir(), "tcp-access.log")

	handler, err := NewHandler(&otypes.TCPAccessLog{FilePath: logFile, Format: "json"})
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
