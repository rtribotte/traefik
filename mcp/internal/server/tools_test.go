package server

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// fakeTarget serves a recorded payload per path, decoding it into out.
type fakeTarget struct {
	responses map[string]json.RawMessage
	err       error
}

func (f *fakeTarget) Name() string { return "fake" }

func (f *fakeTarget) Get(_ context.Context, path string, out any) error {
	if f.err != nil {
		return f.err
	}
	return json.Unmarshal(f.responses[path], out)
}

func newRawdataTarget(t *testing.T) *fakeTarget {
	t.Helper()
	body, err := os.ReadFile("../traefik/testdata/rawdata.json")
	require.NoError(t, err)
	return &fakeTarget{responses: map[string]json.RawMessage{"/api/rawdata": body}}
}

func TestHandlePing(t *testing.T) {
	_, out, err := handlePing(context.Background(), nil, pingInput{})
	require.NoError(t, err)
	assert.Equal(t, "pong", out.Message)
}

func TestListRouters(t *testing.T) {
	handler := listRouters(newRawdataTarget(t))

	_, out, err := handler(context.Background(), nil, listRoutersInput{})
	require.NoError(t, err)

	byName := map[string]RouterSummary{}
	for _, r := range out.Routers {
		byName[r.Name] = r
	}

	require.Contains(t, byName, "web@docker")
	web := byName["web@docker"]
	assert.Equal(t, "Host(`whoami.localhost`)", web.Rule)
	assert.Equal(t, "whoami", web.Service)
	assert.Equal(t, "enabled", web.Status)
	assert.Equal(t, []string{"web"}, web.EntryPoints)

	require.Contains(t, byName, "broken@docker")
	assert.Equal(t, "warning", byName["broken@docker"].Status)
	assert.Contains(t, byName["broken@docker"].Errors, `the service "missing@docker" does not exist`)
}

func TestGetRouter(t *testing.T) {
	handler := getRouter(newRawdataTarget(t))

	_, out, err := handler(context.Background(), nil, getRouterInput{Name: "web@docker"})
	require.NoError(t, err)
	assert.Equal(t, "web@docker", out.Name)
	assert.Equal(t, "whoami", out.Service)

	_, _, err = handler(context.Background(), nil, getRouterInput{Name: "nope@docker"})
	require.Error(t, err)
}
