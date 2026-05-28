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

func TestTailAccess_Filters(t *testing.T) {
	testCases := []struct {
		desc   string
		filter AccessFilter
		want   int
	}{
		{"exact status", AccessFilter{Status: 200}, 1},
		{"exact status 502", AccessFilter{Status: 502}, 2},
		{"status range 2xx", AccessFilter{MinStatus: 200, MaxStatus: 299}, 1},
		{"status range 4xx-5xx", AccessFilter{MinStatus: 400}, 3},
		{"max status only", AccessFilter{MaxStatus: 399}, 1},
		{"host substring", AccessFilter{Host: "billing"}, 2},
		{"host case-insensitive", AccessFilter{Host: "BILLING"}, 2},
		{"router substring", AccessFilter{Router: "whoami"}, 1},
		{"method", AccessFilter{Method: "get"}, 4},
		{"path substring", AccessFilter{Path: "/pay"}, 1},
		{"min duration", AccessFilter{MinDurationMs: 2.6}, 1},
		{"combined", AccessFilter{MinStatus: 500, Host: "billing", Method: "GET"}, 2},
		{"no match", AccessFilter{Status: 418}, 0},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			f, err := os.Open("testdata/access.log")
			require.NoError(t, err)
			defer f.Close()

			entries, err := TailAccess(f, 100, test.filter)
			require.NoError(t, err)
			assert.Len(t, entries, test.want)
		})
	}
}
