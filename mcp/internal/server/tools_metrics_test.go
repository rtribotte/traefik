package server

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newMetricsTarget(t *testing.T) *fakeTarget {
	t.Helper()
	body, err := os.ReadFile("../metrics/testdata/metrics.txt")
	require.NoError(t, err)
	return &fakeTarget{responses: map[string]json.RawMessage{metricsPath: body}}
}

func TestGetReloadStatus(t *testing.T) {
	_, out, err := getReloadStatus(newMetricsTarget(t))(context.Background(), nil, getReloadStatusInput{})
	require.NoError(t, err)
	assert.True(t, out.Success)
	assert.Equal(t, 3, out.Reloads)
	assert.NotEmpty(t, out.LastReload)
}

func TestGetRequestMetrics_MinStatus(t *testing.T) {
	_, out, err := getRequestMetrics(newMetricsTarget(t))(context.Background(), nil, getRequestMetricsInput{MinStatus: 500})
	require.NoError(t, err)
	require.NotEmpty(t, out.Counts)
	for _, c := range out.Counts {
		assert.GreaterOrEqual(t, statusCode(c.Code), 500)
	}
}

func TestGetRequestMetrics_ScopeAndName(t *testing.T) {
	_, out, err := getRequestMetrics(newMetricsTarget(t))(context.Background(), nil, getRequestMetricsInput{Scope: "service", Name: "billing"})
	require.NoError(t, err)
	require.Len(t, out.Counts, 1)
	assert.Equal(t, "billing@docker", out.Counts[0].Name)
	assert.Equal(t, "502", out.Counts[0].Code)
	assert.Equal(t, 7, out.Counts[0].Value)
}
