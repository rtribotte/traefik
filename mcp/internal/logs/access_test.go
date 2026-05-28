package logs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTailAccess(t *testing.T) {
	f, err := os.Open("testdata/access.log")
	require.NoError(t, err)
	defer f.Close()

	entries, err := TailAccess(f, 10, AccessFilter{})
	require.NoError(t, err)

	// Four valid JSON lines; the malformed line is skipped.
	require.Len(t, entries, 4)

	assert.Equal(t, 200, entries[0].Status)
	assert.Equal(t, "whoami@docker", entries[0].Router)
	assert.Equal(t, "whoami-demo@docker", entries[0].Service)
	assert.Equal(t, "GET", entries[0].Method)
	assert.Equal(t, "/", entries[0].Path)
	assert.Equal(t, "whoami.localhost", entries[0].Host)
	assert.InDelta(t, 1.2, entries[0].DurationMs, 0.001)
}

func TestTailAccess_LastN(t *testing.T) {
	f, err := os.Open("testdata/access.log")
	require.NoError(t, err)
	defer f.Close()

	entries, err := TailAccess(f, 2, AccessFilter{})
	require.NoError(t, err)
	require.Len(t, entries, 2)
	// Last two valid lines in file order.
	assert.Equal(t, 502, entries[0].Status)
	assert.Equal(t, 404, entries[1].Status)
}

func TestTailAccess_FilterMinStatus(t *testing.T) {
	f, err := os.Open("testdata/access.log")
	require.NoError(t, err)
	defer f.Close()

	entries, err := TailAccess(f, 10, AccessFilter{MinStatus: 500})
	require.NoError(t, err)
	require.Len(t, entries, 2)
	for _, e := range entries {
		assert.GreaterOrEqual(t, e.Status, 500)
	}
}

func TestTailAccess_FilterService(t *testing.T) {
	f, err := os.Open("testdata/access.log")
	require.NoError(t, err)
	defer f.Close()

	entries, err := TailAccess(f, 10, AccessFilter{Service: "billing"})
	require.NoError(t, err)
	require.Len(t, entries, 2)
	for _, e := range entries {
		assert.Contains(t, e.Service, "billing")
	}
}
