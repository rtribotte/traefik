package server

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func promptText(t *testing.T, res *mcp.GetPromptResult) string {
	t.Helper()
	require.Len(t, res.Messages, 1)
	tc, ok := res.Messages[0].Content.(*mcp.TextContent)
	require.True(t, ok)
	return tc.Text
}

func TestDiagnose(t *testing.T) {
	req := &mcp.GetPromptRequest{Params: &mcp.GetPromptParams{Arguments: map[string]string{
		"problem": "billing returns 502",
		"target":  "billing@docker",
	}}}
	res, err := diagnose(context.Background(), req)
	require.NoError(t, err)

	text := promptText(t, res)
	assert.Contains(t, text, "billing returns 502")
	assert.Contains(t, text, `"billing@docker"`)
	// All three symptom paths and the grounding step are present in the one playbook.
	assert.Contains(t, text, "list_routers")
	assert.Contains(t, text, "get_service_health")
	assert.Contains(t, text, "search_traces")
	assert.Contains(t, text, "validate_traefik_config")
}

func TestDiagnose_NoArg(t *testing.T) {
	req := &mcp.GetPromptRequest{Params: &mcp.GetPromptParams{}}
	res, err := diagnose(context.Background(), req)
	require.NoError(t, err)
	assert.Contains(t, promptText(t, res), "list_routers")
}

func TestServer_ListPrompts(t *testing.T) {
	session := connect(t, newRawdataTarget(t))

	res, err := session.ListPrompts(context.Background(), nil)
	require.NoError(t, err)

	// Prompts are disabled while we measure how the model resolves the scenarios
	// without the guided diagnosis playbook. buildDiagnose stays unit-tested.
	assert.Empty(t, res.Prompts)
}
