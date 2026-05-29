package server

import (
	"context"
	"errors"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/traefik/traefik-mcp/internal/reference"
)

var errNoReference = errors.New("the Traefik reference is unavailable: the embedded catalogue failed to load")

type searchDocsInput struct {
	Query  string `json:"query" jsonschema:"what to look up, e.g. 'forward auth middleware' or 'kubernetes ingress provider'"`
	Source string `json:"source,omitempty" jsonschema:"restrict to a product: 'oss' for Traefik Proxy, 'hub' for Traefik Hub, 'external' for Gateway API / core Kubernetes (default: all)"`
	Limit  int    `json:"limit,omitempty" jsonschema:"max results to return (default 5)"`
}

type searchDocsOutput struct {
	Results []reference.SearchResult `json:"results"`
}

// searchDocs is the discovery step: it returns concept ids to feed
// get_traefik_concept, get_traefik_schema and get_traefik_doc.
func searchDocs(cat *reference.Catalogue) mcp.ToolHandlerFor[searchDocsInput, searchDocsOutput] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, in searchDocsInput) (*mcp.CallToolResult, searchDocsOutput, error) {
		if cat == nil {
			return nil, searchDocsOutput{}, errNoReference
		}
		results, err := cat.Search(ctx, in.Query, in.Source, in.Limit)
		if err != nil {
			return nil, searchDocsOutput{}, err
		}
		return nil, searchDocsOutput{Results: results}, nil
	}
}

type conceptInput struct {
	ID string `json:"id" jsonschema:"the concept id from search_traefik_docs, e.g. 'http.middlewares.forwardauth'"`
}

type conceptOutput struct {
	ID       string `json:"id"`
	Markdown string `json:"markdown"`
}

// getConcept returns the full reference contract (fields, types, defaults) for a
// concept id — what the model should ground on before authoring configuration.
func getConcept(cat *reference.Catalogue) mcp.ToolHandlerFor[conceptInput, conceptOutput] {
	return func(_ context.Context, _ *mcp.CallToolRequest, in conceptInput) (*mcp.CallToolResult, conceptOutput, error) {
		if cat == nil {
			return nil, conceptOutput{}, errNoReference
		}
		md, err := cat.Concept(in.ID)
		if err != nil {
			return nil, conceptOutput{}, err
		}
		return nil, conceptOutput{ID: in.ID, Markdown: md}, nil
	}
}

type schemaOutput struct {
	ID     string `json:"id"`
	Schema string `json:"schema"`
}

// getSchema returns the JSON Schema for a concept id, for validated generation
// or structured output.
func getSchema(cat *reference.Catalogue) mcp.ToolHandlerFor[conceptInput, schemaOutput] {
	return func(_ context.Context, _ *mcp.CallToolRequest, in conceptInput) (*mcp.CallToolResult, schemaOutput, error) {
		if cat == nil {
			return nil, schemaOutput{}, errNoReference
		}
		schema, err := cat.Schema(in.ID)
		if err != nil {
			return nil, schemaOutput{}, err
		}
		return nil, schemaOutput{ID: in.ID, Schema: schema}, nil
	}
}

// getDoc resolves a concept id to its narrative documentation page.
func getDoc(cat *reference.Catalogue) mcp.ToolHandlerFor[conceptInput, reference.DocResult] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, in conceptInput) (*mcp.CallToolResult, reference.DocResult, error) {
		if cat == nil {
			return nil, reference.DocResult{}, errNoReference
		}
		doc, err := cat.Doc(ctx, in.ID)
		if err != nil {
			return nil, reference.DocResult{}, err
		}
		return nil, doc, nil
	}
}

type validateConfigInput struct {
	Config  string `json:"config" jsonschema:"the Traefik configuration to validate, as YAML or JSON; may contain multiple YAML documents"`
	Concept string `json:"concept,omitempty" jsonschema:"optional concept id (e.g. 'http.middlewares.ratelimit') to validate a single fragment against that concept's schema instead of auto-detecting"`
}

// validateConfig validates configuration against the reference schema registry
// by catalogue rotation: a whole traefik.yaml (static or dynamic), a Traefik or
// Gateway API CRD, an annotated Kubernetes manifest, or — with a concept hint —
// a single fragment.
func validateConfig(cat *reference.Catalogue) mcp.ToolHandlerFor[validateConfigInput, reference.Validation] {
	return func(_ context.Context, _ *mcp.CallToolRequest, in validateConfigInput) (*mcp.CallToolResult, reference.Validation, error) {
		if cat == nil {
			return nil, reference.Validation{}, errNoReference
		}
		res, err := cat.Validate([]byte(in.Config), in.Concept)
		if err != nil {
			return nil, reference.Validation{}, err
		}
		return nil, res, nil
	}
}
