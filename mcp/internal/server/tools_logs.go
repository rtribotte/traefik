package server

import (
	"context"
	"fmt"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/traefik/traefik-mcp/internal/logs"
)

type tailAccessLogsInput struct {
	Count     int    `json:"count,omitempty" jsonschema:"max entries to return (default 50)"`
	MinStatus int    `json:"minStatus,omitempty" jsonschema:"only entries with HTTP status >= this (e.g. 500 for server errors)"`
	Service   string `json:"service,omitempty" jsonschema:"only entries whose service name contains this"`
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

		entries, err := logs.TailAccess(f, count, logs.AccessFilter{MinStatus: in.MinStatus, Service: in.Service})
		if err != nil {
			return nil, tailAccessLogsOutput{}, err
		}
		return nil, tailAccessLogsOutput{Entries: entries}, nil
	}
}
