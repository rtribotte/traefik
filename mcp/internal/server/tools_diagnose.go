package server

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Diagnostic playbooks are also exposed as prompts (user-invoked). These tools
// expose the same scripts to the model directly, so they surface in tool-search
// clients where prompts never load. Output is the step-by-step procedure the
// model should then execute with the live read tools.

type diagnoseRouterMissingInput struct {
	Router string `json:"router,omitempty" jsonschema:"fully qualified router name if known (e.g. api@docker), optional"`
}

type diagnosePlaybookOutput struct {
	Playbook string `json:"playbook"`
}

func diagnoseRouterMissingTool(_ context.Context, _ *mcp.CallToolRequest, in diagnoseRouterMissingInput) (*mcp.CallToolResult, diagnosePlaybookOutput, error) {
	return nil, diagnosePlaybookOutput{Playbook: buildRouterMissing(in.Router)}, nil
}

type diagnose5xxInput struct {
	Service string `json:"service,omitempty" jsonschema:"service name if known (e.g. billing@docker), optional"`
	Host    string `json:"host,omitempty" jsonschema:"request host if known (e.g. billing.localhost), optional"`
}

func diagnose5xxTool(_ context.Context, _ *mcp.CallToolRequest, in diagnose5xxInput) (*mcp.CallToolResult, diagnosePlaybookOutput, error) {
	return nil, diagnosePlaybookOutput{Playbook: buildDiagnose5xx(in.Service, in.Host)}, nil
}
