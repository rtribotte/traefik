package server

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// connect starts the server and a client over in-memory transports and returns
// the connected client session.
func connect(t *testing.T, target *fakeTarget) *mcp.ClientSession {
	t.Helper()

	srv := New("traefik-mcp", "test", Deps{Target: target})

	serverTransport, clientTransport := mcp.NewInMemoryTransports()

	ctx := context.Background()
	_, err := srv.Connect(ctx, serverTransport, nil)
	require.NoError(t, err)

	client := mcp.NewClient(&mcp.Implementation{Name: "test-client", Version: "test"}, nil)
	session, err := client.Connect(ctx, clientTransport, nil)
	require.NoError(t, err)
	t.Cleanup(func() { _ = session.Close() })

	return session
}

func TestServer_ListTools(t *testing.T) {
	session := connect(t, newRawdataTarget(t))

	res, err := session.ListTools(context.Background(), nil)
	require.NoError(t, err)

	names := map[string]bool{}
	for _, tool := range res.Tools {
		names[tool.Name] = true
	}
	assert.True(t, names["ping"])
	assert.True(t, names["list_routers"])
	assert.True(t, names["get_router"])
}

func TestServer_CallPing(t *testing.T) {
	session := connect(t, newRawdataTarget(t))

	res, err := session.CallTool(context.Background(), &mcp.CallToolParams{Name: "ping"})
	require.NoError(t, err)
	require.False(t, res.IsError)

	var out pingOutput
	require.NoError(t, json.Unmarshal(mustJSON(t, res.StructuredContent), &out))
	assert.Equal(t, "pong", out.Message)
}

func TestServer_CallListRouters(t *testing.T) {
	session := connect(t, newRawdataTarget(t))

	res, err := session.CallTool(context.Background(), &mcp.CallToolParams{Name: "list_routers"})
	require.NoError(t, err)
	require.False(t, res.IsError)

	var out listRoutersOutput
	require.NoError(t, json.Unmarshal(mustJSON(t, res.StructuredContent), &out))
	assert.Len(t, out.Routers, 2)
}

func mustJSON(t *testing.T, v any) []byte {
	t.Helper()
	b, err := json.Marshal(v)
	require.NoError(t, err)
	return b
}
