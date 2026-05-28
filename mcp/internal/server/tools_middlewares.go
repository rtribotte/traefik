package server

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/traefik/traefik-mcp/internal/traefik"
)

// MiddlewareSummary is the structured projection of a middleware returned to the model.
type MiddlewareSummary struct {
	Name   string   `json:"name"`
	Status string   `json:"status,omitempty"`
	Types  []string `json:"types,omitempty"`
	UsedBy []string `json:"usedBy,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

type listMiddlewaresInput struct{}

type listMiddlewaresOutput struct {
	ConfigHash  string              `json:"configHash"`
	Middlewares []MiddlewareSummary `json:"middlewares"`
}

func listMiddlewares(target traefik.Target) mcp.ToolHandlerFor[listMiddlewaresInput, listMiddlewaresOutput] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, _ listMiddlewaresInput) (*mcp.CallToolResult, listMiddlewaresOutput, error) {
		raw, err := traefik.FetchRawData(ctx, target)
		if err != nil {
			return nil, listMiddlewaresOutput{}, err
		}

		out := listMiddlewaresOutput{ConfigHash: raw.Hash(), Middlewares: make([]MiddlewareSummary, 0, len(raw.Middlewares))}
		for name, info := range raw.Middlewares {
			summary := MiddlewareSummary{
				Name:   name,
				Status: info.Status,
				UsedBy: info.UsedBy,
				Errors: info.Err,
				Types:  configKeys(info.Middleware),
			}
			out.Middlewares = append(out.Middlewares, summary)
		}
		return nil, out, nil
	}
}

type getMiddlewareInput struct {
	Name string `json:"name" jsonschema:"fully qualified middleware name, e.g. auth@docker"`
}

func getMiddleware(target traefik.Target) mcp.ToolHandlerFor[getMiddlewareInput, MiddlewareSummary] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, in getMiddlewareInput) (*mcp.CallToolResult, MiddlewareSummary, error) {
		raw, err := traefik.FetchRawData(ctx, target)
		if err != nil {
			return nil, MiddlewareSummary{}, err
		}

		info, ok := raw.Middlewares[in.Name]
		if !ok {
			return nil, MiddlewareSummary{}, fmt.Errorf("middleware %q not found", in.Name)
		}
		return nil, MiddlewareSummary{
			Name:   in.Name,
			Status: info.Status,
			UsedBy: info.UsedBy,
			Errors: info.Err,
			Types:  configKeys(info.Middleware),
		}, nil
	}
}
