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
}

// New builds an MCP server with the read-only Traefik tools registered.
func New(name, version string, deps Deps) *mcp.Server {
	s := mcp.NewServer(&mcp.Implementation{Name: name, Version: version}, nil)
	addReadTools(s, deps.Target)
	return s
}
