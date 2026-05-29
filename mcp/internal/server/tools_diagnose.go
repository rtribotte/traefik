package server

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// The diagnostic playbook is also exposed as a prompt (user-invoked). This tool
// exposes the same script to the model directly, so it surfaces in tool-search
// clients where prompts never load. Output is the step-by-step procedure the
// model should then execute with the live read tools.

type diagnoseInput struct {
	Problem string `json:"problem,omitempty" jsonschema:"what's wrong, in the user's words (e.g. 'api.localhost 404s', 'billing returns 502', 'checkout is slow'), optional"`
	Target  string `json:"target,omitempty" jsonschema:"the router, service or host involved if known (e.g. api@docker, billing.localhost), optional"`
}

type diagnosePlaybookOutput struct {
	Playbook string `json:"playbook"`
}

func diagnoseTool(_ context.Context, _ *mcp.CallToolRequest, in diagnoseInput) (*mcp.CallToolResult, diagnosePlaybookOutput, error) {
	return nil, diagnosePlaybookOutput{Playbook: buildDiagnose(in.Problem, in.Target)}, nil
}
