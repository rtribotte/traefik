package server

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDiagnoseRouterMissingTool(t *testing.T) {
	_, out, err := diagnoseRouterMissingTool(context.Background(), nil, diagnoseRouterMissingInput{Router: "api@docker"})
	require.NoError(t, err)
	assert.Contains(t, out.Playbook, `router "api@docker"`)
	assert.Contains(t, out.Playbook, "list_routers")
}

func TestDiagnose5xxTool(t *testing.T) {
	_, out, err := diagnose5xxTool(context.Background(), nil, diagnose5xxInput{Service: "billing@docker"})
	require.NoError(t, err)
	assert.Contains(t, out.Playbook, `service "billing@docker"`)
	assert.Contains(t, out.Playbook, "get_service_health")
}

func TestDiagnoseTools_SharePromptText(t *testing.T) {
	_, out, err := diagnose5xxTool(context.Background(), nil, diagnose5xxInput{Service: "billing@docker", Host: "billing.localhost"})
	require.NoError(t, err)
	assert.Equal(t, buildDiagnose5xx("billing@docker", "billing.localhost"), out.Playbook)
}
