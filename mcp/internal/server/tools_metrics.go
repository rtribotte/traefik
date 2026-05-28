package server

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/traefik/traefik-mcp/internal/metrics"
	"github.com/traefik/traefik-mcp/internal/traefik"
)

// metricsPath is where Traefik serves Prometheus metrics. With api.insecure the
// metrics endpoint shares the API entrypoint; a dedicated metrics URL can be
// wired later if the deployment separates them.
const metricsPath = "/metrics"

func fetchMetrics(ctx context.Context, target traefik.Target) (map[string]*metrics.MetricFamily, error) {
	body, err := target.GetRaw(ctx, metricsPath)
	if err != nil {
		return nil, fmt.Errorf("fetching metrics: %w", err)
	}
	return metrics.Parse(bytes.NewReader(body))
}

type getReloadStatusInput struct{}

type getReloadStatusOutput struct {
	Success        bool   `json:"success"`
	LastReload     string `json:"lastReload,omitempty"`
	LastReloadAgeS int    `json:"lastReloadAgeSeconds,omitempty"`
	Reloads        int    `json:"reloads"`
}

func getReloadStatus(target traefik.Target) mcp.ToolHandlerFor[getReloadStatusInput, getReloadStatusOutput] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, _ getReloadStatusInput) (*mcp.CallToolResult, getReloadStatusOutput, error) {
		fams, err := fetchMetrics(ctx, target)
		if err != nil {
			return nil, getReloadStatusOutput{}, err
		}

		rs := metrics.ReloadStatusFrom(fams)
		out := getReloadStatusOutput{Success: rs.Success, Reloads: rs.Reloads}
		if !rs.LastReload.IsZero() {
			out.LastReload = rs.LastReload.Format(time.RFC3339)
			out.LastReloadAgeS = int(time.Since(rs.LastReload).Seconds())
		}
		return nil, out, nil
	}
}

type getRequestMetricsInput struct {
	Scope     string `json:"scope,omitempty" jsonschema:"limit to a scope: service or entrypoint (default both)"`
	Name      string `json:"name,omitempty" jsonschema:"only counts whose service/entrypoint name contains this (case-insensitive)"`
	MinStatus int    `json:"minStatus,omitempty" jsonschema:"only counts with HTTP status code >= this (e.g. 500 for server errors)"`
}

type getRequestMetricsOutput struct {
	Counts []metrics.RequestCount `json:"counts"`
}

func getRequestMetrics(target traefik.Target) mcp.ToolHandlerFor[getRequestMetricsInput, getRequestMetricsOutput] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, in getRequestMetricsInput) (*mcp.CallToolResult, getRequestMetricsOutput, error) {
		fams, err := fetchMetrics(ctx, target)
		if err != nil {
			return nil, getRequestMetricsOutput{}, err
		}

		var counts []metrics.RequestCount
		for _, c := range metrics.RequestCounts(fams) {
			if in.Scope != "" && c.Scope != in.Scope {
				continue
			}
			if in.Name != "" && !strings.Contains(strings.ToLower(c.Name), strings.ToLower(in.Name)) {
				continue
			}
			if in.MinStatus > 0 && statusCode(c.Code) < in.MinStatus {
				continue
			}
			counts = append(counts, c)
		}
		return nil, getRequestMetricsOutput{Counts: counts}, nil
	}
}

func statusCode(code string) int {
	n := 0
	for _, r := range code {
		if r < '0' || r > '9' {
			return 0
		}
		n = n*10 + int(r-'0')
	}
	return n
}
