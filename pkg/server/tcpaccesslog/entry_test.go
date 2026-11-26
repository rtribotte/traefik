package tcpaccesslog

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
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
	otypes "github.com/traefik/traefik/v3/pkg/observability/types"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/plog/plogotlp"
)

func TestLogger_LogConnection(t *testing.T) {
	testCases := []struct {
		desc     string
		entry    Entry
		expected map[string]any
	}{
		{
			desc: "connection handled by a router",
			entry: Entry{
				EntryPointName: "web",
				RouterName:     "router@file",
				ServiceName:    "service@file",
				TLSMode:        "none",
				Stats:          &ConnStats{ReceivedBytes: 4, SentBytes: 5, Duration: time.Second},
			},
			expected: map[string]any{
				"traefik.entrypoint.name":           "web",
				"traefik.router.name":               "router@file",
				"traefik.service.name":              "service@file",
				"traefik.tls.mode":                  "none",
				"traefik.connection.received_bytes": float64(4),
				"traefik.connection.sent_bytes":     float64(5),
				"traefik.connection.duration":       float64(1),
				"client.address":                    "10.0.0.1",
				"client.port":                       float64(4242),
			},
		},
		{
			desc: "connection closed without a router handling it",
			entry: Entry{
				EntryPointName: "web",
				ServerName:     "foo.localhost",
				ErrorType:      ErrorTypeNoMatchingRouter,
			},
			expected: map[string]any{
				"traefik.entrypoint.name": "web",
				"server.address":          "foo.localhost",
				"error.type":              "no_matching_router",
				"client.address":          "10.0.0.1",
				"client.port":             float64(4242),
			},
		},
		{
			desc: "connection closed by the client",
			entry: Entry{
				EntryPointName: "web",
				ErrorType:      ErrorTypeInvalidClientHello,
				Err:            io.ErrUnexpectedEOF,
			},
			expected: map[string]any{
				"traefik.entrypoint.name": "web",
				"error.type":              "connection_closed",
				"client.address":          "10.0.0.1",
				"client.port":             float64(4242),
			},
		},
		{
			desc: "connection timed out",
			entry: Entry{
				EntryPointName: "web",
				ErrorType:      ErrorTypeConnectionClosed,
				Err:            timeoutError{},
			},
			expected: map[string]any{
				"traefik.entrypoint.name": "web",
				"error.type":              "timeout",
				"client.address":          "10.0.0.1",
				"client.port":             float64(4242),
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			logger, logFile := newLogger(t)

			logger.LogConnection(addrConn{remoteAddr: &net.TCPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 4242}}, test.entry)

			entries := readEntries(t, logFile)
			require.Len(t, entries, 1)

			entry := entries[0]
			assert.NotEmpty(t, entry["time"])
			delete(entry, "time")
			delete(entry, "level")
			delete(entry, "msg")

			assert.Equal(t, test.expected, entry)
		})
	}
}

func TestLogger_LogConnection_disabled(t *testing.T) {
	var logger *Logger

	assert.NotPanics(t, func() {
		logger.LogConnection(addrConn{}, Entry{EntryPointName: "web"})
	})
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

func TestLogger_LogConnection_otlp(t *testing.T) {
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

			logger.LogConnection(addrConn{}, Entry{
				EntryPointName: "web",
				RouterName:     "router@file",
				Stats:          &ConnStats{ReceivedBytes: 4},
			})

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

// addrConn is a connection which only has a remote address.
type addrConn struct {
	net.Conn

	remoteAddr net.Addr
}

func (c addrConn) RemoteAddr() net.Addr {
	return c.remoteAddr
}

type stringAddr string

func (s stringAddr) Network() string { return "tcp" }

func (s stringAddr) String() string { return string(s) }

type timeoutError struct{}

func (timeoutError) Error() string { return "i/o timeout" }

func (timeoutError) Timeout() bool { return true }

func (timeoutError) Temporary() bool { return true }
