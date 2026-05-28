package server

import (
	"context"
	"fmt"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/traefik/traefik-mcp/internal/logs"
)

type tailAccessLogsInput struct {
	Count         int     `json:"count,omitempty" jsonschema:"max entries to return, newest last (default 50)"`
	Status        int     `json:"status,omitempty" jsonschema:"only entries with this exact HTTP status (e.g. 404)"`
	MinStatus     int     `json:"minStatus,omitempty" jsonschema:"only entries with HTTP status >= this (e.g. 500 for server errors)"`
	MaxStatus     int     `json:"maxStatus,omitempty" jsonschema:"only entries with HTTP status <= this (e.g. 299 with minStatus 200 for successes)"`
	Service       string  `json:"service,omitempty" jsonschema:"only entries whose service name contains this (case-insensitive)"`
	Router        string  `json:"router,omitempty" jsonschema:"only entries whose router name contains this (case-insensitive)"`
	Host          string  `json:"host,omitempty" jsonschema:"only entries whose request host contains this (case-insensitive)"`
	Method        string  `json:"method,omitempty" jsonschema:"only entries with this HTTP method (e.g. POST)"`
	Path          string  `json:"path,omitempty" jsonschema:"only entries whose request path contains this (case-insensitive)"`
	MinDurationMs float64 `json:"minDurationMs,omitempty" jsonschema:"only entries that took at least this many milliseconds, for finding slow requests"`
}

type tailAccessLogsOutput struct {
	Entries []logs.AccessEntry `json:"entries"`
}

func tailAccessLogs(path string) mcp.ToolHandlerFor[tailAccessLogsInput, tailAccessLogsOutput] {
	return func(_ context.Context, _ *mcp.CallToolRequest, in tailAccessLogsInput) (*mcp.CallToolResult, tailAccessLogsOutput, error) {
		if path == "" {
			return nil, tailAccessLogsOutput{}, fmt.Errorf("access log path not configured; start traefik-mcp with --traefik.access-log pointing at the JSON access log file")
		}

		f, err := os.Open(path)
		if err != nil {
			return nil, tailAccessLogsOutput{}, fmt.Errorf("opening access log %s: %w", path, err)
		}
		defer f.Close()

		count := in.Count
		if count <= 0 {
			count = 50
		}

		entries, err := logs.TailAccess(f, count, logs.AccessFilter{
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
			return nil, tailAccessLogsOutput{}, err
		}
		return nil, tailAccessLogsOutput{Entries: entries}, nil
	}
}
