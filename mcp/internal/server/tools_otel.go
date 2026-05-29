package server

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/traefik/traefik-mcp/internal/logs"
	"github.com/traefik/traefik-mcp/internal/loki"
	"github.com/traefik/traefik-mcp/internal/prom"
)

var errNoLoki = errors.New("access log querying is not configured: start traefik-mcp with --loki.url pointing at a Loki (otel-lgtm) instance")

var errNoProm = errors.New("metrics querying is not configured: start traefik-mcp with --prometheus.url pointing at a Prometheus (otel-lgtm) instance")

// lokiFetch caps how many recent lines are pulled from Loki before client-side
// filtering. Generous so selective filters still have material to match.
const lokiFetch = 1000

type queryAccessLogsInput struct {
	Count           int     `json:"count,omitempty" jsonschema:"max entries to return, newest last (default 50)"`
	LookbackMinutes int     `json:"lookbackMinutes,omitempty" jsonschema:"how far back to search, in minutes (default 60)"`
	Status          int     `json:"status,omitempty" jsonschema:"only entries with this exact HTTP status (e.g. 404)"`
	MinStatus       int     `json:"minStatus,omitempty" jsonschema:"only entries with HTTP status >= this (e.g. 500 for server errors)"`
	MaxStatus       int     `json:"maxStatus,omitempty" jsonschema:"only entries with HTTP status <= this"`
	Service         string  `json:"service,omitempty" jsonschema:"only entries whose service name contains this (case-insensitive)"`
	Router          string  `json:"router,omitempty" jsonschema:"only entries whose router name contains this (case-insensitive)"`
	Host            string  `json:"host,omitempty" jsonschema:"only entries whose request host contains this (case-insensitive)"`
	Method          string  `json:"method,omitempty" jsonschema:"only entries with this HTTP method (e.g. POST)"`
	Path            string  `json:"path,omitempty" jsonschema:"only entries whose request path contains this (case-insensitive)"`
	MinDurationMs   float64 `json:"minDurationMs,omitempty" jsonschema:"only entries that took at least this many milliseconds, for finding slow requests"`
	LogQL           string  `json:"logQL,omitempty" jsonschema:"optional raw LogQL stream selector (default {service_name=\"traefik\"})"`
}

type queryAccessLogsOutput struct {
	Entries []logs.AccessEntry `json:"entries"`
}

// queryAccessLogs reads Traefik's access logs from Loki (shipped over OTLP). The
// Loki line bodies are the same JSON records Traefik writes to a file, so the
// file parser and filter are reused verbatim.
func queryAccessLogs(client *loki.Client) mcp.ToolHandlerFor[queryAccessLogsInput, queryAccessLogsOutput] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, in queryAccessLogsInput) (*mcp.CallToolResult, queryAccessLogsOutput, error) {
		if client == nil {
			return nil, queryAccessLogsOutput{}, errNoLoki
		}

		count := in.Count
		if count <= 0 {
			count = 50
		}
		lookback := time.Duration(in.LookbackMinutes) * time.Minute

		lines, err := client.QueryRange(ctx, in.LogQL, lokiFetch, lookback)
		if err != nil {
			return nil, queryAccessLogsOutput{}, err
		}

		entries, err := logs.TailAccess(strings.NewReader(strings.Join(lines, "\n")), count, logs.AccessFilter{
			Status:        in.Status,
			MinStatus:     in.MinStatus,
			MaxStatus:     in.MaxStatus,
			Service:       in.Service,
			Router:        in.Router,
			Host:          in.Host,
			Method:        in.Method,
			Path:          in.Path,
			MinDurationMs: in.MinDurationMs,
		})
		if err != nil {
			return nil, queryAccessLogsOutput{}, err
		}
		return nil, queryAccessLogsOutput{Entries: entries}, nil
	}
}

type queryMetricsInput struct {
	Query        string `json:"query" jsonschema:"PromQL expression, e.g. sum by (service,code) (rate(traefik_service_requests_total[5m]))"`
	RangeMinutes int    `json:"rangeMinutes,omitempty" jsonschema:"if set, run a range query over the last N minutes (returns series) instead of an instant query"`
	StepSeconds  int    `json:"stepSeconds,omitempty" jsonschema:"range query resolution in seconds (default 15)"`
}

type queryMetricsOutput struct {
	Samples []prom.Sample `json:"samples,omitempty"`
	Series  []prom.Series `json:"series,omitempty"`
}

// queryMetrics runs a PromQL query against the Prometheus that otel-lgtm fills
// from Traefik's OTLP metrics. Unlike get_metrics (a raw scrape of the current
// values), this exposes history and aggregation (rates, sums, quantiles).
func queryMetrics(client *prom.Client) mcp.ToolHandlerFor[queryMetricsInput, queryMetricsOutput] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, in queryMetricsInput) (*mcp.CallToolResult, queryMetricsOutput, error) {
		if client == nil {
			return nil, queryMetricsOutput{}, errNoProm
		}

		if in.RangeMinutes > 0 {
			step := time.Duration(in.StepSeconds) * time.Second
			end := time.Now()
			start := end.Add(-time.Duration(in.RangeMinutes) * time.Minute)
			series, err := client.QueryRange(ctx, in.Query, start, end, step)
			if err != nil {
				return nil, queryMetricsOutput{}, err
			}
			return nil, queryMetricsOutput{Series: series}, nil
		}

		samples, err := client.Query(ctx, in.Query)
		if err != nil {
			return nil, queryMetricsOutput{}, err
		}
		return nil, queryMetricsOutput{Samples: samples}, nil
	}
}
