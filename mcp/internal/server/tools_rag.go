package server

import (
	"context"
	"errors"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/traefik/traefik-mcp/internal/rag"
)

var errNoRetriever = errors.New("documentation search is unavailable: the embedded reference corpus failed to load")

type searchDocsInput struct {
	Query  string `json:"query" jsonschema:"what to look up, e.g. 'forward auth middleware' or 'kubernetes ingress provider'"`
	Source string `json:"source,omitempty" jsonschema:"restrict to a product: 'oss' for Traefik Proxy or 'hub' for Traefik Hub (default: both)"`
	Limit  int    `json:"limit,omitempty" jsonschema:"max results to return (default 5)"`
}

type searchDocsOutput struct {
	Results []rag.Result `json:"results"`
}

func searchDocs(r rag.Retriever) mcp.ToolHandlerFor[searchDocsInput, searchDocsOutput] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, in searchDocsInput) (*mcp.CallToolResult, searchDocsOutput, error) {
		if r == nil {
			return nil, searchDocsOutput{}, errNoRetriever
		}
		results, err := r.Search(ctx, in.Query, in.Source, in.Limit)
		if err != nil {
			return nil, searchDocsOutput{}, err
		}
		return nil, searchDocsOutput{Results: results}, nil
	}
}
