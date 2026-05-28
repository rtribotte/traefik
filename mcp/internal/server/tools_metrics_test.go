package server

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetMetrics(t *testing.T) {
	raw := "traefik_config_reloads_total 3\n"
	target := &fakeTarget{responses: map[string]json.RawMessage{metricsPath: json.RawMessage(raw)}}

	_, out, err := getMetrics(target)(context.Background(), nil, getMetricsInput{})
	require.NoError(t, err)
	assert.Equal(t, raw, out.Metrics)
}
