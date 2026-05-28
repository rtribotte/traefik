package staticconf

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/static"
	otypes "github.com/traefik/traefik/v3/pkg/observability/types"
)

func TestDetect(t *testing.T) {
	cfg := &static.Configuration{
		Metrics:   &otypes.Metrics{Prometheus: &otypes.Prometheus{}, OTLP: &otypes.OTLP{}},
		Tracing:   &static.Tracing{OTLP: &otypes.OTelTracing{}},
		AccessLog: &otypes.AccessLog{FilePath: "/logs/access.log", OTLP: &otypes.OTelLog{}},
		Log:       &otypes.TraefikLog{FilePath: "/logs/traefik.log"},
	}

	got := Detect(cfg)
	assert.Equal(t, Capabilities{
		PrometheusMetrics: true,
		OTLPMetrics:       true,
		OTLPTracing:       true,
		FileAccessLog:     true,
		OTLPAccessLog:     true,
		FileAppLog:        true,
		OTLPAppLog:        false,
	}, got)
}

func TestDetectEmpty(t *testing.T) {
	assert.Equal(t, Capabilities{}, Detect(&static.Configuration{}))
	assert.Equal(t, Capabilities{}, Detect(nil))
}

func TestFromLog(t *testing.T) {
	log := strings.Join([]string{
		`{"level":"info","time":"t1","message":"Traefik version 3.6.0"}`,
		`{"level":"debug","time":"t2","staticConfiguration":{"metrics":{"prometheus":{}},"tracing":{"otlp":{}},"accessLog":{"filePath":"/logs/access.log"},"log":{"filePath":"/logs/traefik.log"}},"message":"Static configuration loaded [json]"}`,
		`{"level":"info","time":"t3","message":"Server configuration reloaded"}`,
	}, "\n")

	cfg, err := FromLog(strings.NewReader(log))
	require.NoError(t, err)

	caps := Detect(cfg)
	assert.True(t, caps.PrometheusMetrics)
	assert.True(t, caps.OTLPTracing)
	assert.True(t, caps.FileAccessLog)
	assert.True(t, caps.FileAppLog)
	assert.False(t, caps.OTLPMetrics)
}

func TestFromLogLatestWins(t *testing.T) {
	log := strings.Join([]string{
		`{"staticConfiguration":{"metrics":{"prometheus":{}}},"message":"Static configuration loaded [json]"}`,
		`{"staticConfiguration":{"tracing":{"otlp":{}}},"message":"Static configuration loaded [json]"}`,
	}, "\n")

	cfg, err := FromLog(strings.NewReader(log))
	require.NoError(t, err)

	caps := Detect(cfg)
	assert.False(t, caps.PrometheusMetrics) // superseded by the later restart.
	assert.True(t, caps.OTLPTracing)
}

func TestFromLogNotLogged(t *testing.T) {
	log := `{"level":"info","message":"Traefik version 3.6.0"}`

	_, err := FromLog(strings.NewReader(log))
	assert.ErrorIs(t, err, ErrNotLogged)
}
