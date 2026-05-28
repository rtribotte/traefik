// Package staticconf recovers a running Traefik's static (install) configuration
// from its application log and derives which observability data sources it
// exposes. Traefik has no API endpoint for its static configuration, but it logs
// the whole redacted configuration as JSON at startup (debug level); reading that
// line lets the MCP server register only the tools whose backing data exists.
package staticconf

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"

	"github.com/traefik/traefik/v3/pkg/config/static"
)

// logMessage is the zerolog message Traefik attaches to the static-configuration
// dump (cmd/traefik/traefik.go). It carries the redacted configuration in the
// "staticConfiguration" field and is emitted only at debug level.
const logMessage = "Static configuration loaded [json]"

// ErrNotLogged means the static-configuration line was not found in the log.
// Traefik emits it only at debug level, so this usually means log.level is not
// DEBUG or the log does not reach back to startup.
var ErrNotLogged = errors.New("static configuration not found in log: Traefik logs it only at debug level (set log.level=DEBUG)")

// Capabilities reports which observability data sources a Traefik instance is
// configured to expose. The MCP server registers a data-source-backed tool only
// when the matching capability is present.
type Capabilities struct {
	PrometheusMetrics bool `json:"prometheusMetrics"`
	OTLPMetrics       bool `json:"otlpMetrics"`
	OTLPTracing       bool `json:"otlpTracing"`
	FileAccessLog     bool `json:"fileAccessLog"`
	OTLPAccessLog     bool `json:"otlpAccessLog"`
	FileAppLog        bool `json:"fileAppLog"`
	OTLPAppLog        bool `json:"otlpAppLog"`
}

// Detect derives the capabilities from a parsed static configuration. A nil
// configuration yields the zero value (no capabilities).
func Detect(cfg *static.Configuration) Capabilities {
	var c Capabilities
	if cfg == nil {
		return c
	}
	if cfg.Metrics != nil {
		c.PrometheusMetrics = cfg.Metrics.Prometheus != nil
		c.OTLPMetrics = cfg.Metrics.OTLP != nil
	}
	if cfg.Tracing != nil {
		c.OTLPTracing = cfg.Tracing.OTLP != nil
	}
	if cfg.AccessLog != nil {
		c.FileAccessLog = cfg.AccessLog.FilePath != ""
		c.OTLPAccessLog = cfg.AccessLog.OTLP != nil
	}
	if cfg.Log != nil {
		c.FileAppLog = cfg.Log.FilePath != ""
		c.OTLPAppLog = cfg.Log.OTLP != nil
	}
	return c
}

// FromLog scans a Traefik JSON application log for the startup line that carries
// the redacted static configuration and returns it parsed. The last occurrence
// wins, so a restarted instance reflects its newest configuration. Returns
// ErrNotLogged if no such line is present.
func FromLog(r io.Reader) (*static.Configuration, error) {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 0, 64*1024), 8*1024*1024)

	var latest json.RawMessage
	for scanner.Scan() {
		var line struct {
			Message string          `json:"message"`
			Static  json.RawMessage `json:"staticConfiguration"`
		}
		if err := json.Unmarshal(scanner.Bytes(), &line); err != nil {
			continue
		}
		if line.Message == logMessage && len(line.Static) > 0 {
			latest = line.Static
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if latest == nil {
		return nil, ErrNotLogged
	}

	var cfg static.Configuration
	if err := json.Unmarshal(latest, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
