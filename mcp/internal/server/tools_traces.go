package server

import (
	"context"
	"errors"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/traefik/traefik-mcp/internal/tempo"
)

// errNoTempo is returned by the trace tools when no Tempo URL was configured.
var errNoTempo = errors.New("trace querying is not configured: start traefik-mcp with --tempo.url pointing at a Tempo (otel-lgtm) instance")

type searchTracesInput struct {
	Query string `json:"query,omitempty" jsonschema:"A TraceQL filter, e.g. {duration>1s} for slow traces, {status=error} for failed ones, or {resource.service.name=\"traefik\"}. Empty matches all recent traces."`
	Limit int    `json:"limit,omitempty" jsonschema:"Maximum number of traces to return (default 20)."`
}

type searchTracesOutput struct {
	Traces []tempo.TraceSummary `json:"traces"`
}

func searchTraces(client *tempo.Client) mcp.ToolHandlerFor[searchTracesInput, searchTracesOutput] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, in searchTracesInput) (*mcp.CallToolResult, searchTracesOutput, error) {
		if client == nil {
			return nil, searchTracesOutput{}, errNoTempo
		}
		traces, err := client.Search(ctx, in.Query, in.Limit)
		if err != nil {
			return nil, searchTracesOutput{}, err
		}
		return nil, searchTracesOutput{Traces: traces}, nil
	}
}

type getTraceInput struct {
	TraceID string `json:"traceID" jsonschema:"The trace ID to fetch, as returned by search_traces."`
}

type getTraceOutput struct {
	Spans []tempo.Span `json:"spans"`
}

func getTrace(client *tempo.Client) mcp.ToolHandlerFor[getTraceInput, getTraceOutput] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, in getTraceInput) (*mcp.CallToolResult, getTraceOutput, error) {
		if client == nil {
			return nil, getTraceOutput{}, errNoTempo
		}
		spans, err := client.Trace(ctx, in.TraceID)
		if err != nil {
			return nil, getTraceOutput{}, err
		}
		return nil, getTraceOutput{Spans: spans}, nil
	}
}
