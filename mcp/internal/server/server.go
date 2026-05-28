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
	// log-tailing tools. Empty disables them at call time with a helpful error.
	AccessLogPath string
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
		Description: "Return recent entries from Traefik's access log, newest last. Filter by " +
			"minStatus (e.g. 500 to see only server errors) or service name. Use this to " +
			"investigate 5xx errors, latency or which router/service served a request.",
	}, tailAccessLogs(deps.AccessLogPath))

	addPrompts(s)

	return s
}
