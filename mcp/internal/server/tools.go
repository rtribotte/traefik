package server

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/traefik/traefik-mcp/internal/traefik"
	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

// addReadTools registers the deterministic, read-only tools backed by the
// Traefik API on the server.
func addReadTools(s *mcp.Server, target traefik.Target) {
	mcp.AddTool(s, &mcp.Tool{
		Name:        "ping",
		Description: "Health check for the MCP server. Returns \"pong\".",
	}, handlePing)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "list_routers",
		Description: "List all HTTP routers with their rule, service, status and errors.",
	}, listRouters(target))

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_router",
		Description: "Get a single HTTP router by its fully qualified name (e.g. my-router@docker).",
	}, getRouter(target))
}

type pingInput struct{}

type pingOutput struct {
	Message string `json:"message"`
}

func handlePing(_ context.Context, _ *mcp.CallToolRequest, _ pingInput) (*mcp.CallToolResult, pingOutput, error) {
	return nil, pingOutput{Message: "pong"}, nil
}

// RouterSummary is the structured projection of a router returned to the model.
type RouterSummary struct {
	Name        string   `json:"name"`
	Rule        string   `json:"rule,omitempty"`
	Service     string   `json:"service,omitempty"`
	Status      string   `json:"status,omitempty"`
	EntryPoints []string `json:"entryPoints,omitempty"`
	Middlewares []string `json:"middlewares,omitempty"`
	Errors      []string `json:"errors,omitempty"`
}

type listRoutersInput struct{}

type listRoutersOutput struct {
	Routers []RouterSummary `json:"routers"`
}

func listRouters(target traefik.Target) mcp.ToolHandlerFor[listRoutersInput, listRoutersOutput] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, _ listRoutersInput) (*mcp.CallToolResult, listRoutersOutput, error) {
		raw, err := traefik.FetchRawData(ctx, target)
		if err != nil {
			return nil, listRoutersOutput{}, err
		}

		out := listRoutersOutput{Routers: make([]RouterSummary, 0, len(raw.Routers))}
		for name, info := range raw.Routers {
			out.Routers = append(out.Routers, routerSummary(name, info))
		}
		return nil, out, nil
	}
}

type getRouterInput struct {
	Name string `json:"name" jsonschema:"fully qualified router name, e.g. my-router@docker"`
}

func getRouter(target traefik.Target) mcp.ToolHandlerFor[getRouterInput, RouterSummary] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, in getRouterInput) (*mcp.CallToolResult, RouterSummary, error) {
		raw, err := traefik.FetchRawData(ctx, target)
		if err != nil {
			return nil, RouterSummary{}, err
		}

		info, ok := raw.Routers[in.Name]
		if !ok {
			return nil, RouterSummary{}, fmt.Errorf("router %q not found", in.Name)
		}
		return nil, routerSummary(in.Name, info), nil
	}
}

func routerSummary(name string, info *runtime.RouterInfo) RouterSummary {
	s := RouterSummary{
		Name:   name,
		Status: info.Status,
		Errors: info.Err,
	}
	if info.Router != nil {
		s.Rule = info.Rule
		s.Service = info.Service
		s.EntryPoints = info.EntryPoints
		s.Middlewares = info.Middlewares
	}
	return s
}
