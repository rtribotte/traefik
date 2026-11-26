package server

import (
	"bufio"
	"context"
	"crypto/tls"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/static"
	"github.com/traefik/traefik/v3/pkg/middlewares/tcp/accesslog"
	otypes "github.com/traefik/traefik/v3/pkg/observability/types"
	tcprouter "github.com/traefik/traefik/v3/pkg/server/router/tcp"
	"github.com/traefik/traefik/v3/pkg/tls/generate"
)

func TestTCPEntryPoint_unservedHTTPConnAccessLog(t *testing.T) {
	cert, err := generate.DefaultCertificate()
	require.NoError(t, err)

	testCases := []struct {
		desc              string
		client            func(t *testing.T, addr string)
		expectedErrorType string
	}{
		{
			desc: "bytes which are not an HTTP request",
			client: func(t *testing.T, addr string) {
				t.Helper()

				conn, err := (&net.Dialer{}).DialContext(t.Context(), "tcp", addr)
				require.NoError(t, err)
				defer conn.Close()

				_, err = conn.Write([]byte("not an HTTP request\r\n\r\n"))
				require.NoError(t, err)

				// The HTTP server replies with a 400 status code, and closes the connection.
				_, _ = io.Copy(io.Discard, conn)
			},
			expectedErrorType: "no_http_request",
		},
		{
			desc: "failed TLS handshake",
			client: func(t *testing.T, addr string) {
				t.Helper()

				conn, err := (&net.Dialer{}).DialContext(t.Context(), "tcp", addr)
				require.NoError(t, err)
				defer conn.Close()

				// The client rejects the self-signed certificate of the HTTPS server, which aborts the handshake.
				tlsConn := tls.Client(conn, &tls.Config{ServerName: "foo.localhost"})
				require.Error(t, tlsConn.Handshake())
			},
			expectedErrorType: "tls_handshake_error",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			addr, logFile := startTCPAccessLoggedEntryPoint(t, &tls.Config{Certificates: []tls.Certificate{*cert}})

			test.client(t, addr)

			// The connection is logged once the HTTP server reports it closed.
			var entries []map[string]any
			require.Eventually(t, func() bool {
				entries = readTCPAccessLogEntries(t, logFile)
				return len(entries) > 0
			}, 5*time.Second, 10*time.Millisecond)
			require.Len(t, entries, 1)

			entry := entries[0]
			assert.Equal(t, "web", entry["traefik.entrypoint.name"])
			assert.Equal(t, test.expectedErrorType, entry["error.type"])
			assert.Equal(t, "127.0.0.1", entry["client.address"])
			assert.NotContains(t, entry, "traefik.connection.received_bytes")
			assert.NotContains(t, entry, "traefik.router.name")
		})
	}
}

func TestTCPEntryPoint_servedHTTPConnNotInTCPAccessLog(t *testing.T) {
	addr, logFile := startTCPAccessLoggedEntryPoint(t, nil)

	client := &http.Client{Transport: &http.Transport{DisableKeepAlives: true}}

	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "http://"+addr, nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// The served connections are only logged in the HTTP access logs,
	// so there is no TCP access log line to wait for.
	time.Sleep(200 * time.Millisecond)
	assert.Empty(t, readTCPAccessLogEntries(t, logFile))
}

// startTCPAccessLoggedEntryPoint starts an entry point with the TCP access logs enabled,
// whose HTTP and HTTPS routers reply with a 200 status code,
// and returns its address and the TCP access log file.
func startTCPAccessLoggedEntryPoint(t *testing.T, httpsTLSConfig *tls.Config) (string, string) {
	t.Helper()

	epConfig := &static.EntryPointsTransport{}
	epConfig.SetDefaults()

	logFile := filepath.Join(t.TempDir(), "tcp-access.log")

	logger, err := accesslog.NewLogger(t.Context(), &otypes.TCPAccessLog{FilePath: logFile})
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, logger.Close()) })

	entryPoint, err := NewTCPEntryPoint(t.Context(), "web", &static.EntryPoint{
		Address:          "127.0.0.1:0",
		Transport:        epConfig,
		ForwardedHeaders: &static.ForwardedHeaders{},
		HTTP2:            &static.HTTP2Config{},
	}, nil, nil, logger)
	require.NoError(t, err)

	router, err := tcprouter.NewRouter(nil)
	require.NoError(t, err)

	router.SetAccessLogger(accesslog.NewEntryPointLogger(logger, "web"))

	okHandler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})
	router.SetHTTPHandler(okHandler)
	if httpsTLSConfig != nil {
		router.SetHTTPSHandler(okHandler, httpsTLSConfig)
	}

	// The entry point is not probed by a connection, which would be logged too,
	// as its listener accepts connections already.
	entryPoint.SwitchRouter(router)
	go entryPoint.Start(t.Context())

	t.Cleanup(func() { entryPoint.Shutdown(context.Background()) })

	return entryPoint.listener.Addr().String(), logFile
}

func readTCPAccessLogEntries(t *testing.T, logFile string) []map[string]any {
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
