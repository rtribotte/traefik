package accesslog

import (
	"bufio"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	otypes "github.com/traefik/traefik/v3/pkg/observability/types"
	"github.com/traefik/traefik/v3/pkg/server/tcpaccesslog"
	"github.com/traefik/traefik/v3/pkg/tcp"
	"github.com/traefik/traefik/v3/pkg/tls/generate"
)

func TestRouterAccessLogHandler(t *testing.T) {
	testCases := []struct {
		desc    string
		logger  *tcpaccesslog.Logger
		tls     *dynamic.RouterTCPTLSConfig
		wrapped bool
		tlsMode string
	}{
		{
			desc: "TCP access logs disabled",
		},
		{
			desc:    "without TLS",
			logger:  &tcpaccesslog.Logger{},
			wrapped: true,
			tlsMode: "none",
		},
		{
			desc:    "TLS terminated",
			logger:  &tcpaccesslog.Logger{},
			tls:     &dynamic.RouterTCPTLSConfig{},
			wrapped: true,
			tlsMode: "terminated",
		},
		{
			desc:    "TLS passthrough",
			logger:  &tcpaccesslog.Logger{},
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
				assert.Equal(t, test.tlsMode, middleware.entry.TLSMode)
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
	assert.NotContains(t, entry, "error.type")

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

func newLogger(t *testing.T) (*tcpaccesslog.Logger, string) {
	t.Helper()

	logFile := filepath.Join(t.TempDir(), "tcp-access.log")

	logger, err := tcpaccesslog.NewLogger(t.Context(), &otypes.TCPAccessLog{FilePath: logFile})
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
