package server

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/traefik/traefik-mcp/internal/traefik"
)

type listCertificatesInput struct{}

type listCertificatesOutput struct {
	Certificates []traefik.Certificate `json:"certificates"`
}

func listCertificates(target traefik.Target) mcp.ToolHandlerFor[listCertificatesInput, listCertificatesOutput] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, _ listCertificatesInput) (*mcp.CallToolResult, listCertificatesOutput, error) {
		certs, err := traefik.FetchCertificates(ctx, target)
		if err != nil {
			return nil, listCertificatesOutput{}, err
		}
		return nil, listCertificatesOutput{Certificates: certs}, nil
	}
}
