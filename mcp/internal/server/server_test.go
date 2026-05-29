package server

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik-mcp/internal/configschema"
	"github.com/traefik/traefik-mcp/internal/rag"
	"github.com/traefik/traefik-mcp/internal/staticconf"
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

func TestServer_GatesToolsOnCapabilities(t *testing.T) {
	tools := func(deps Deps) map[string]bool {
		srv := New("traefik-mcp", "test", deps)
		serverTransport, clientTransport := mcp.NewInMemoryTransports()
		ctx := context.Background()
		_, err := srv.Connect(ctx, serverTransport, nil)
		require.NoError(t, err)
		client := mcp.NewClient(&mcp.Implementation{Name: "c", Version: "t"}, nil)
		session, err := client.Connect(ctx, clientTransport, nil)
		require.NoError(t, err)
		t.Cleanup(func() { _ = session.Close() })

		res, err := session.ListTools(ctx, nil)
		require.NoError(t, err)
		names := map[string]bool{}
		for _, tool := range res.Tools {
			names[tool.Name] = true
		}
		return names
	}

	target := newRawdataTarget(t)

	// Nil caps: every data-source tool is registered (back-compat).
	all := tools(Deps{Target: target})
	for _, name := range []string{"tail_access_logs", "tail_traefik_logs", "get_metrics", "search_traces", "get_trace", "query_access_logs", "query_metrics"} {
		assert.True(t, all[name], name)
	}

	// Caps with only tracing: trace tools present, the rest gated out.
	traceOnly := tools(Deps{Target: target, Caps: &staticconf.Capabilities{OTLPTracing: true}})
	assert.True(t, traceOnly["search_traces"])
	assert.True(t, traceOnly["get_trace"])
	assert.False(t, traceOnly["get_metrics"])
	assert.False(t, traceOnly["tail_access_logs"])
	assert.False(t, traceOnly["tail_traefik_logs"])
	assert.False(t, traceOnly["query_access_logs"])
	assert.False(t, traceOnly["query_metrics"])

	// OTLP logs/metrics caps register exactly the otel-lgtm query tools.
	otlp := tools(Deps{Target: target, Caps: &staticconf.Capabilities{OTLPAccessLog: true, OTLPMetrics: true}})
	assert.True(t, otlp["query_access_logs"])
	assert.True(t, otlp["query_metrics"])
	assert.False(t, otlp["tail_access_logs"])
	assert.False(t, otlp["get_metrics"])
	assert.False(t, otlp["search_traces"])

	// Core read tools and prompts-as-tools are always present.
	assert.True(t, traceOnly["list_routers"])
	assert.True(t, traceOnly["diagnose_5xx"])

	// Validation and docs tools are capability-independent: always present.
	for _, name := range []string{"validate_static_config", "validate_dynamic_config", "search_traefik_docs"} {
		assert.True(t, traceOnly[name], name)
		assert.True(t, otlp[name], name)
	}
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

func connectWith(t *testing.T, deps Deps) *mcp.ClientSession {
	t.Helper()
	srv := New("traefik-mcp", "test", deps)
	serverTransport, clientTransport := mcp.NewInMemoryTransports()
	ctx := context.Background()
	_, err := srv.Connect(ctx, serverTransport, nil)
	require.NoError(t, err)
	client := mcp.NewClient(&mcp.Implementation{Name: "c", Version: "t"}, nil)
	session, err := client.Connect(ctx, clientTransport, nil)
	require.NoError(t, err)
	t.Cleanup(func() { _ = session.Close() })
	return session
}

func TestServer_ValidateDynamicConfig(t *testing.T) {
	v, err := configschema.New()
	require.NoError(t, err)
	session := connectWith(t, Deps{Target: newRawdataTarget(t), Validator: v})

	res, err := session.CallTool(context.Background(), &mcp.CallToolParams{
		Name: "validate_dynamic_config",
		Arguments: map[string]any{
			"config": "http:\n  services:\n    s:\n      loadBalancer:\n        servers: \"not-an-array\"\n",
		},
	})
	require.NoError(t, err)
	require.False(t, res.IsError)

	var out validateConfigOutput
	require.NoError(t, json.Unmarshal(mustJSON(t, res.StructuredContent), &out))
	assert.False(t, out.Valid)
	assert.NotEmpty(t, out.Problems)
}

func TestServer_SearchTraefikDocs(t *testing.T) {
	r, err := rag.NewEmbedded()
	require.NoError(t, err)
	session := connectWith(t, Deps{Target: newRawdataTarget(t), Retriever: r})

	res, err := session.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "search_traefik_docs",
		Arguments: map[string]any{"query": "forwardauth middleware", "limit": 3},
	})
	require.NoError(t, err)
	require.False(t, res.IsError)

	var out searchDocsOutput
	require.NoError(t, json.Unmarshal(mustJSON(t, res.StructuredContent), &out))
	require.NotEmpty(t, out.Results)
	assert.Contains(t, out.Results[0].ID, "forwardauth")
}

func mustJSON(t *testing.T, v any) []byte {
	t.Helper()
	b, err := json.Marshal(v)
	require.NoError(t, err)
	return b
}
