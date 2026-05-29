package server

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"

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

	mcp.AddTool(s, &mcp.Tool{
		Name:        "list_services",
		Description: "List all HTTP services with their type, backend servers, per-server health and status.",
	}, listServices(target))

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_service",
		Description: "Get a single HTTP service by its fully qualified name (e.g. whoami@docker).",
	}, getService(target))

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_service_health",
		Description: "Report the per-server health of an HTTP service: how many backends are UP vs DOWN.",
	}, getServiceHealth(target))

	mcp.AddTool(s, &mcp.Tool{
		Name:        "list_middlewares",
		Description: "List all HTTP middlewares with their type, status and the routers using them.",
	}, listMiddlewares(target))

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_middleware",
		Description: "Get a single HTTP middleware by its fully qualified name (e.g. auth@docker).",
	}, getMiddleware(target))

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_entrypoints",
		Description: "List the configured entry points (listeners): name, address and whether TLS is the default.",
	}, getEntryPoints(target))

	mcp.AddTool(s, &mcp.Tool{
		Name: "list_certificates",
		Description: "List the TLS certificates Traefik is currently serving: common name, SANs, " +
			"validity window, days until expiry, issuer, key type and the expiry status Traefik " +
			"computes (enabled, warning when under 30 days left, or expired). Use it for any " +
			"certificate or expiry question — it is the authoritative live view, richer than the " +
			"traefik_tls_certs_not_after metric. Returns an empty list when no TLS is configured.",
	}, listCertificates(target))

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_overview",
		Description: "Summary counts of routers, services and middlewares per protocol, plus enabled features and providers.",
	}, getOverview(target))

	mcp.AddTool(s, &mcp.Tool{
		Name: "get_config_hash",
		Description: "Return a fingerprint (sha256) of Traefik's current runtime configuration, " +
			"plus resource counts. Call this before relying on previously fetched routers, " +
			"services or middlewares: if the hash differs from the last one you saw, the " +
			"configuration changed and you must re-fetch it. Cheap to call repeatedly.",
	}, getConfigHash(target))
}

type getConfigHashInput struct{}

type configHashOutput struct {
	Hash        string `json:"hash"`
	Routers     int    `json:"routers"`
	Services    int    `json:"services"`
	Middlewares int    `json:"middlewares"`
}

func getConfigHash(target traefik.Target) mcp.ToolHandlerFor[getConfigHashInput, configHashOutput] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, _ getConfigHashInput) (*mcp.CallToolResult, configHashOutput, error) {
		raw, err := traefik.FetchRawData(ctx, target)
		if err != nil {
			return nil, configHashOutput{}, err
		}

		return nil, configHashOutput{
			Hash:        raw.Hash(),
			Routers:     len(raw.Routers) + len(raw.TCPRouters) + len(raw.UDPRouters),
			Services:    len(raw.Services) + len(raw.TCPServices) + len(raw.UDPServices),
			Middlewares: len(raw.Middlewares) + len(raw.TCPMiddlewares),
		}, nil
	}
}

// configKeys marshals a dynamic config object and returns its top-level keys,
// which identify the configured kind(s) — e.g. "loadBalancer" for a service or
// "basicAuth" for a middleware. It avoids enumerating every dynamic field by
// hand and stays correct as upstream adds new ones.
func configKeys(v any) []string {
	b, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return nil
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
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
	ConfigHash string          `json:"configHash"`
	Routers    []RouterSummary `json:"routers"`
}

func listRouters(target traefik.Target) mcp.ToolHandlerFor[listRoutersInput, listRoutersOutput] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, _ listRoutersInput) (*mcp.CallToolResult, listRoutersOutput, error) {
		raw, err := traefik.FetchRawData(ctx, target)
		if err != nil {
			return nil, listRoutersOutput{}, err
		}

		out := listRoutersOutput{ConfigHash: raw.Hash(), Routers: make([]RouterSummary, 0, len(raw.Routers))}
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
