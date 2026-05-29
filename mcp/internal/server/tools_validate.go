package server

import (
	"context"
	"errors"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/traefik/traefik-mcp/internal/configschema"
)

var errNoValidator = errors.New("configuration validation is unavailable: the embedded schemas failed to load")

type validateConfigInput struct {
	Config string `json:"config" jsonschema:"the Traefik configuration to validate, as YAML or JSON"`
}

type validateConfigOutput struct {
	Valid    bool                   `json:"valid"`
	Problems []configschema.Problem `json:"problems,omitempty"`
}

func validateConfig(v *configschema.Validator, kind configschema.Kind) mcp.ToolHandlerFor[validateConfigInput, validateConfigOutput] {
	return func(_ context.Context, _ *mcp.CallToolRequest, in validateConfigInput) (*mcp.CallToolResult, validateConfigOutput, error) {
		if v == nil {
			return nil, validateConfigOutput{}, errNoValidator
		}
		problems, err := v.Validate(kind, []byte(in.Config))
		if err != nil {
			return nil, validateConfigOutput{}, err
		}
		return nil, validateConfigOutput{Valid: len(problems) == 0, Problems: problems}, nil
	}
}
