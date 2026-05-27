package server

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/traefik/traefik-mcp/internal/traefik"
)

type getEntryPointsInput struct{}

type getEntryPointsOutput struct {
	EntryPoints []traefik.EntryPoint `json:"entryPoints"`
}

func getEntryPoints(target traefik.Target) mcp.ToolHandlerFor[getEntryPointsInput, getEntryPointsOutput] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, _ getEntryPointsInput) (*mcp.CallToolResult, getEntryPointsOutput, error) {
		eps, err := traefik.FetchEntryPoints(ctx, target)
		if err != nil {
			return nil, getEntryPointsOutput{}, err
		}
		return nil, getEntryPointsOutput{EntryPoints: eps}, nil
	}
}

type getOverviewInput struct{}

func getOverview(target traefik.Target) mcp.ToolHandlerFor[getOverviewInput, traefik.Overview] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, _ getOverviewInput) (*mcp.CallToolResult, traefik.Overview, error) {
		ov, err := traefik.FetchOverview(ctx, target)
		if err != nil {
			return nil, traefik.Overview{}, err
		}
		return nil, *ov, nil
	}
}
