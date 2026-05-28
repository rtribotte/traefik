package logs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func openAppLog(t *testing.T) *os.File {
	t.Helper()
	f, err := os.Open("testdata/traefik.log")
	require.NoError(t, err)
	t.Cleanup(func() { f.Close() })
	return f
}

func TestTailApp(t *testing.T) {
	entries, err := TailApp(openAppLog(t), 10, AppFilter{})
	require.NoError(t, err)
	// Five valid JSON lines; the malformed line is skipped.
	require.Len(t, entries, 5)

	assert.Equal(t, "info", entries[0].Level)
	assert.Equal(t, "3.7.1", entries[0].Fields["version"])

	last := entries[4]
	assert.Equal(t, "error", last.Level)
	assert.Contains(t, last.Error, "middleware")
	assert.Equal(t, "broken@file", last.Fields["routerName"])
}

func TestTailApp_Filters(t *testing.T) {
	testCases := []struct {
		desc   string
		filter AppFilter
		want   int
	}{
		{"exact level error", AppFilter{Level: "error"}, 2},
		{"exact level info", AppFilter{Level: "info"}, 2},
		{"min level warn", AppFilter{MinLevel: "warn"}, 3},
		{"contains service", AppFilter{Contains: "service"}, 1},
		{"contains middleware", AppFilter{Contains: "middleware"}, 2},
		{"contains case-insensitive", AppFilter{Contains: "MIDDLEWARE"}, 2},
		{"combined", AppFilter{Level: "error", Contains: "does not exist"}, 2},
		{"no match", AppFilter{Level: "fatal"}, 0},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			entries, err := TailApp(openAppLog(t), 100, test.filter)
			require.NoError(t, err)
			assert.Len(t, entries, test.want)
		})
	}
}
