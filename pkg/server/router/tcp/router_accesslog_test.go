package tcp

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"io"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	accesslogMiddle "github.com/traefik/traefik/v3/pkg/middlewares/tcp/accesslog"
	otypes "github.com/traefik/traefik/v3/pkg/observability/types"
	traefiktcp "github.com/traefik/traefik/v3/pkg/tcp"
)

func TestRouter_ServeTCP_unhandledAccessLog(t *testing.T) {
	testCases := []struct {
		desc string
		// setup configures the router, and the server side of the connection.
		setup  func(t *testing.T, router *Router, serverConn net.Conn)
		client func(clientConn net.Conn)

		expectedErrorType  string
		expectedRouterName string
		expectedServerName string
		expectReceived     bool
	}{
		{
			desc: "client closing the connection before sending data",
			client: func(clientConn net.Conn) {
				_ = clientConn.Close()
			},
			expectedErrorType: "connection_closed",
		},
		{
			desc: "client not sending data in time",
			setup: func(t *testing.T, _ *Router, serverConn net.Conn) {
				t.Helper()
				require.NoError(t, serverConn.SetReadDeadline(time.Now().Add(10*time.Millisecond)))
			},
			client:            func(net.Conn) {},
			expectedErrorType: "timeout",
		},
		{
			desc: "no matching router for a non-TLS connection",
			client: func(clientConn net.Conn) {
				_, _ = clientConn.Write([]byte("hello traefik\n"))
				_, _ = io.Copy(io.Discard, clientConn)
			},
			expectedErrorType: "no_matching_router",
			expectReceived:    true,
		},
		{
			desc: "invalid TLS ClientHello",
			client: func(clientConn net.Conn) {
				// A TLS handshake record header, followed by bytes which are not a ClientHello.
				_, _ = clientConn.Write([]byte{0x16, 0x03, 0x01, 0x00, 0x05, 0xff, 0xff, 0xff, 0xff, 0xff})
				_, _ = io.Copy(io.Discard, clientConn)
			},
			expectedErrorType: "invalid_client_hello",
			expectReceived:    true,
		},
		{
			desc: "no matching router for a TLS connection",
			client: func(clientConn net.Conn) {
				_ = tls.Client(clientConn, &tls.Config{ServerName: "foo.localhost", InsecureSkipVerify: true}).Handshake()
			},
			expectedErrorType:  "no_matching_router",
			expectedServerName: "foo.localhost",
			expectReceived:     true,
		},
		{
			desc: "router with invalid TLS options",
			setup: func(t *testing.T, router *Router, _ net.Conn) {
				t.Helper()

				brokenRouter := &brokenTLSRouter{accessLogger: router.accessLogger, routerName: "router@file"}
				require.NoError(t, router.muxerTCPTLS.AddRoute("HostSNI(`foo.localhost`)", "", 1, "file", brokenRouter))
			},
			client: func(clientConn net.Conn) {
				_ = tls.Client(clientConn, &tls.Config{ServerName: "foo.localhost", InsecureSkipVerify: true}).Handshake()
			},
			expectedErrorType:  "invalid_tls_options",
			expectedRouterName: "router@file",
			expectReceived:     true,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			router, logFile := newAccessLoggedRouter(t)

			serverConn, clientConn := net.Pipe()
			t.Cleanup(func() {
				_ = serverConn.Close()
				_ = clientConn.Close()
			})

			if test.setup != nil {
				test.setup(t, router, serverConn)
			}

			go test.client(clientConn)

			router.ServeTCP(pipeWriteCloser{Conn: serverConn})

			entries := readAccessLogEntries(t, logFile)
			require.Len(t, entries, 1)

			entry := entries[0]
			assert.Equal(t, "web", entry["traefik.entrypoint.name"])
			assert.Equal(t, test.expectedErrorType, entry["error.type"])
			assert.Equal(t, "10.0.0.1", entry["client.address"])
			assert.InDelta(t, 4242, entry["client.port"], 0)
			assert.InDelta(t, 0, entry["traefik.connection.sent_bytes"], 0)
			assert.Contains(t, entry, "traefik.connection.duration")
			assert.NotContains(t, entry, "traefik.service.name")

			if test.expectReceived {
				assert.Positive(t, entry["traefik.connection.received_bytes"])
			} else {
				assert.InDelta(t, 0, entry["traefik.connection.received_bytes"], 0)
			}

			if test.expectedRouterName == "" {
				assert.NotContains(t, entry, "traefik.router.name")
			} else {
				assert.Equal(t, test.expectedRouterName, entry["traefik.router.name"])
			}

			if test.expectedServerName == "" {
				assert.NotContains(t, entry, "server.address")
			} else {
				assert.Equal(t, test.expectedServerName, entry["server.address"])
			}
		})
	}
}

func TestRouter_ServeTCP_handledConnectionNotLoggedByEntryPoint(t *testing.T) {
	router, logFile := newAccessLoggedRouter(t)

	handler := traefiktcp.HandlerFunc(func(conn traefiktcp.WriteCloser) {
		_ = conn.Close()
	})
	require.NoError(t, router.muxerTCP.AddRoute("HostSNI(`*`)", "", 1, "file", handler))

	serverConn, clientConn := net.Pipe()
	t.Cleanup(func() {
		_ = serverConn.Close()
		_ = clientConn.Close()
	})

	router.ServeTCP(pipeWriteCloser{Conn: serverConn})

	// The connections handled by a router are logged by the access log middleware of the router.
	assert.Empty(t, readAccessLogEntries(t, logFile))
}

func newAccessLoggedRouter(t *testing.T) (*Router, string) {
	t.Helper()

	logFile := filepath.Join(t.TempDir(), "tcp-access.log")

	logger, err := accesslogMiddle.NewLogger(t.Context(), &otypes.TCPAccessLog{FilePath: logFile})
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, logger.Close()) })

	router, err := NewRouter(nil)
	require.NoError(t, err)

	router.SetAccessLogger(accesslogMiddle.NewEntryPointLogger(logger, "web"))

	return router, logFile
}

func readAccessLogEntries(t *testing.T, logFile string) []map[string]any {
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

// pipeWriteCloser adds the CloseWrite method required by tcp.WriteCloser to a net.Pipe connection,
// and gives it a TCP remote address, as the routing reads the client IP from it.
type pipeWriteCloser struct {
	net.Conn
}

func (p pipeWriteCloser) CloseWrite() error {
	return p.Close()
}

func (p pipeWriteCloser) RemoteAddr() net.Addr {
	return &net.TCPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 4242}
}
