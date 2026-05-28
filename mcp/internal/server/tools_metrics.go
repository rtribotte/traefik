package server

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/traefik/traefik-mcp/internal/traefik"
)

// metricsPath is where Traefik serves Prometheus metrics. With api.insecure the
// metrics endpoint shares the API entrypoint; a dedicated metrics URL can be
// wired later if the deployment separates them.
const metricsPath = "/metrics"

type getMetricsInput struct{}

type getMetricsOutput struct {
	Metrics string `json:"metrics"`
}

// getMetrics returns Traefik's raw Prometheus exposition text. Parsing and
// per-metric filtering are deliberately left to the caller for now.
func getMetrics(target traefik.Target) mcp.ToolHandlerFor[getMetricsInput, getMetricsOutput] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, _ getMetricsInput) (*mcp.CallToolResult, getMetricsOutput, error) {
		body, err := target.GetRaw(ctx, metricsPath)
		if err != nil {
			return nil, getMetricsOutput{}, fmt.Errorf("fetching metrics: %w", err)
		}
		return nil, getMetricsOutput{Metrics: string(body)}, nil
	}
}
