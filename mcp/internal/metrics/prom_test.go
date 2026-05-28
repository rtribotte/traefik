package metrics

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func parseFixture(t *testing.T) map[string]*MetricFamily {
	t.Helper()
	f, err := os.Open("testdata/metrics.txt")
	require.NoError(t, err)
	defer f.Close()

	fams, err := Parse(f)
	require.NoError(t, err)
	return fams
}

func TestReloadStatus(t *testing.T) {
	rs := ReloadStatusFrom(parseFixture(t))
	assert.True(t, rs.Success)
	assert.Equal(t, 3, rs.Reloads)
	assert.Equal(t, int64(1779978056), rs.LastReload.Unix())
}

func TestRequestCounts(t *testing.T) {
	counts := RequestCounts(parseFixture(t))

	// Two service entries and two entrypoint entries.
	var svc, ep int
	for _, c := range counts {
		switch c.Scope {
		case "service":
			svc++
		case "entrypoint":
			ep++
		}
	}
	assert.Equal(t, 3, svc)
	assert.Equal(t, 2, ep)

	// The billing 502 must be present with the right count.
	var found bool
	for _, c := range counts {
		if c.Scope == "service" && c.Name == "billing@docker" && c.Code == "502" {
			found = true
			assert.Equal(t, 7, c.Value)
			assert.Equal(t, "GET", c.Method)
		}
	}
	assert.True(t, found, "billing@docker 502 count not found")
}

func TestParse_Empty(t *testing.T) {
	fams, err := Parse(strings.NewReader(""))
	require.NoError(t, err)
	rs := ReloadStatusFrom(fams)
	assert.False(t, rs.Success)
	assert.True(t, rs.LastReload.IsZero())
	assert.Empty(t, RequestCounts(fams))
}
