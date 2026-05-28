// Package server wires the Traefik-backed tools, resources and prompts onto an
// MCP server.
package server

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/traefik/traefik-mcp/internal/traefik"
)

// Deps holds the collaborators the MCP surface is built on.
type Deps struct {
	Target traefik.Target
	// AccessLogPath is the host path to Traefik's JSON access log, enabling the
	// access-log tool. Empty disables it at call time with a helpful error.
	AccessLogPath string
	// AppLogPath is the host path to Traefik's JSON application log, enabling the
	// app-log tool. Empty disables it at call time with a helpful error.
	AppLogPath string
}

// instructions guide the client to treat Traefik's configuration as live state.
// Without this, models answer follow-up questions from earlier tool results even
// after the configuration changed under them.
const instructions = `This server exposes a live, read-only view of a running Traefik instance.

The configuration is dynamic: routers, services and middlewares can change between
your turns (a new container, an edited config file). Never answer a question about
the current state from earlier tool results — call the relevant tool again.

Every list tool returns a "configHash" fingerprint of the configuration. If it
differs from the one you saw previously, the configuration changed and any cached
understanding is stale. Use get_config_hash for a cheap explicit check.

When diagnosing a router or service, always re-fetch its current state first.`

// New builds an MCP server with the read-only Traefik tools registered.
func New(name, version string, deps Deps) *mcp.Server {
	s := mcp.NewServer(
		&mcp.Implementation{Name: name, Version: version},
		&mcp.ServerOptions{Instructions: instructions},
	)
	addReadTools(s, deps.Target)

	mcp.AddTool(s, &mcp.Tool{
		Name: "tail_access_logs",
		Description: "Return recent entries from Traefik's access log, newest last. All filters are " +
			"optional and combine with AND: exact status, status range (minStatus/maxStatus), " +
			"service, router, host, method, path substring, and minDurationMs. Use it for any " +
			"traffic question — errors, slow requests, which router/service served a host, " +
			"traffic to a path, requests by method, etc.",
	}, tailAccessLogs(deps.AccessLogPath))

	mcp.AddTool(s, &mcp.Tool{
		Name: "tail_traefik_logs",
		Description: "Return recent entries from Traefik's application log, newest last. This is " +
			"where Traefik reports configuration and runtime errors (unresolved middleware/service " +
			"references, TLS/certificate problems, provider connection failures, invalid config). " +
			"Filters are optional and combine with AND: level (exact), minLevel (e.g. warn for " +
			"warnings and errors), and contains (substring in the message or error).",
	}, tailAppLogs(deps.AppLogPath))

	mcp.AddTool(s, &mcp.Tool{
		Name: "diagnose_router_missing",
		Description: "Return a step-by-step playbook for diagnosing why a Traefik router is " +
			"missing or not routing traffic. Follow the returned steps using the live read tools.",
	}, diagnoseRouterMissingTool)

	mcp.AddTool(s, &mcp.Tool{
		Name: "diagnose_5xx",
		Description: "Return a step-by-step playbook for diagnosing 5xx errors and determining " +
			"whether the fault is Traefik routing, the service config, or the backend app. " +
			"Follow the returned steps using the live read tools.",
	}, diagnose5xxTool)

	mcp.AddTool(s, &mcp.Tool{
		Name: "get_metrics",
		Description: "Return Traefik's raw Prometheus /metrics output as text: config reload " +
			"status, request/response counts and durations by entrypoint and service, open " +
			"connections, TLS certificate expiry, and more.",
	}, getMetrics(deps.Target))

	addPrompts(s)

	return s
}
