package tcpaccesslog

import (
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
)

func TestLogger_bufferingSize(t *testing.T) {
	logFile := filepath.Join(t.TempDir(), "tcp-access.log")

	logger, err := NewLogger(t.Context(), &otypes.TCPAccessLog{FilePath: logFile, BufferingSize: 2})
	require.NoError(t, err)

	end := time.Date(2026, time.November, 25, 12, 0, 0, 0, time.UTC)
	for range 10 {
		logger.logChan <- logLine{fields: logrus.Fields{"traefik.router.name": "router@file"}, time: end}
	}

	// Close writes the buffered lines before closing the file.
	require.NoError(t, logger.Close())

	entries := readEntries(t, logFile)
	require.Len(t, entries, 10)

	// The line keeps the time the connection ended, not the time it is written.
	assert.Equal(t, end.Format(time.RFC3339Nano), entries[0]["time"])
}

func TestLogger_Rotate(t *testing.T) {
	logFile := filepath.Join(t.TempDir(), "tcp-access.log")
	rotatedFile := logFile + ".1"

	logger, err := NewLogger(t.Context(), &otypes.TCPAccessLog{FilePath: logFile})
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, logger.Close()) })

	logger.log(logrus.Fields{"traefik.router.name": "before"})

	require.NoError(t, os.Rename(logFile, rotatedFile))
	require.NoError(t, logger.Rotate())

	logger.log(logrus.Fields{"traefik.router.name": "after"})

	rotated := readEntries(t, rotatedFile)
	require.Len(t, rotated, 1)
	assert.Equal(t, "before", rotated[0]["traefik.router.name"])

	current := readEntries(t, logFile)
	require.Len(t, current, 1)
	assert.Equal(t, "after", current[0]["traefik.router.name"])
}

func TestLogger_Rotate_otlpOnly(t *testing.T) {
	collector := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	t.Cleanup(collector.Close)

	logFile := filepath.Join(t.TempDir(), "tcp-access.log")

	logger, err := NewLogger(t.Context(), &otypes.TCPAccessLog{
		FilePath: logFile,
		OTLP: &otypes.OTelLog{
			HTTP: &otypes.OTelHTTP{Endpoint: collector.URL},
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, logger.Close()) })

	require.NoError(t, logger.Rotate())

	logger.log(logrus.Fields{"traefik.router.name": "router@file"})

	// The regular output stays discarded after a rotation when only OTLP is used.
	assert.Empty(t, readEntries(t, logFile))
}

func TestLogger_Rotate_stdout(t *testing.T) {
	logger, err := NewLogger(t.Context(), &otypes.TCPAccessLog{})
	require.NoError(t, err)

	require.NoError(t, logger.Rotate())

	// Closing the logger does not close the standard output.
	require.NoError(t, logger.Close())
	_, err = os.Stdout.Stat()
	assert.NoError(t, err)
}
