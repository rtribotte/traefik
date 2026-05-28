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

func TestDiagnoseRouterMissing(t *testing.T) {
	req := &mcp.GetPromptRequest{Params: &mcp.GetPromptParams{Arguments: map[string]string{"router": "api@docker"}}}
	res, err := diagnoseRouterMissing(context.Background(), req)
	require.NoError(t, err)

	text := promptText(t, res)
	assert.Contains(t, text, `router "api@docker"`)
	assert.Contains(t, text, "list_routers")
	assert.Contains(t, text, "get_router")
	assert.Contains(t, text, "tail_access_logs")
}

func TestDiagnoseRouterMissing_NoArg(t *testing.T) {
	req := &mcp.GetPromptRequest{Params: &mcp.GetPromptParams{}}
	res, err := diagnoseRouterMissing(context.Background(), req)
	require.NoError(t, err)
	assert.Contains(t, promptText(t, res), "list_routers")
}

func TestDiagnose5xx(t *testing.T) {
	req := &mcp.GetPromptRequest{Params: &mcp.GetPromptParams{Arguments: map[string]string{
		"service": "billing@docker",
		"host":    "billing.localhost",
	}}}
	res, err := diagnose5xx(context.Background(), req)
	require.NoError(t, err)

	text := promptText(t, res)
	assert.Contains(t, text, `service "billing@docker"`)
	assert.Contains(t, text, "billing.localhost")
	assert.Contains(t, text, "get_service_health")
	assert.Contains(t, text, "tail_access_logs")
}

func TestServer_ListPrompts(t *testing.T) {
	session := connect(t, newRawdataTarget(t))

	res, err := session.ListPrompts(context.Background(), nil)
	require.NoError(t, err)

	names := map[string]bool{}
	for _, p := range res.Prompts {
		names[p.Name] = true
	}
	assert.True(t, names["diagnose_router_missing"])
	assert.True(t, names["diagnose_5xx"])
}
