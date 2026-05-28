package server

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTailAccessLogs(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "access.log")
	line := `{"DownstreamStatus":502,"ServiceName":"billing@docker","RequestPath":"/pay","RequestMethod":"GET","RequestHost":"billing.localhost"}` + "\n" +
		`{"DownstreamStatus":200,"ServiceName":"whoami@docker","RequestPath":"/","RequestMethod":"GET","RequestHost":"whoami.localhost"}` + "\n"
	require.NoError(t, os.WriteFile(path, []byte(line), 0o644))

	_, out, err := tailAccessLogs(path)(context.Background(), nil, tailAccessLogsInput{MinStatus: 500})
	require.NoError(t, err)
	require.Len(t, out.Entries, 1)
	assert.Equal(t, 502, out.Entries[0].Status)
	assert.Equal(t, "billing@docker", out.Entries[0].Service)
}

func TestTailAccessLogs_NoPath(t *testing.T) {
	_, _, err := tailAccessLogs("")(context.Background(), nil, tailAccessLogsInput{})
	require.Error(t, err)
}

func TestTailAccessLogs_MissingFile(t *testing.T) {
	_, _, err := tailAccessLogs("/nonexistent/access.log")(context.Background(), nil, tailAccessLogsInput{})
	require.Error(t, err)
}
