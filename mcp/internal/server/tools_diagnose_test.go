package server

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDiagnoseTool(t *testing.T) {
	_, out, err := diagnoseTool(context.Background(), nil, diagnoseInput{Problem: "api.localhost 404s", Target: "api@docker"})
	require.NoError(t, err)
	assert.Contains(t, out.Playbook, "api.localhost 404s")
	assert.Contains(t, out.Playbook, "list_routers")
}

// The tool and the prompt render the exact same playbook, so they never drift.
func TestDiagnoseTool_SharesPromptText(t *testing.T) {
	_, out, err := diagnoseTool(context.Background(), nil, diagnoseInput{Problem: "billing returns 502", Target: "billing@docker"})
	require.NoError(t, err)
	assert.Equal(t, buildDiagnose("billing returns 502", "billing@docker"), out.Playbook)
}
