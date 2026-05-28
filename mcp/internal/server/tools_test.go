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
	return newFixtureTarget(t, map[string]string{"/api/rawdata": "rawdata.json"})
}

func newFixtureTarget(t *testing.T, files map[string]string) *fakeTarget {
	t.Helper()
	responses := make(map[string]json.RawMessage, len(files))
	for path, file := range files {
		body, err := os.ReadFile("../traefik/testdata/" + file)
		require.NoError(t, err)
		responses[path] = body
	}
	return &fakeTarget{responses: responses}
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

func TestListServices(t *testing.T) {
	handler := listServices(newRawdataTarget(t))

	_, out, err := handler(context.Background(), nil, listServicesInput{})
	require.NoError(t, err)
	require.Len(t, out.Services, 1)

	svc := out.Services[0]
	assert.Equal(t, "whoami@docker", svc.Name)
	assert.Equal(t, "enabled", svc.Status)
	assert.Contains(t, svc.Types, "loadBalancer")
	assert.Equal(t, []string{"http://172.17.0.2:80"}, svc.Servers)
	assert.Equal(t, "UP", svc.ServerStatus["http://172.17.0.2:80"])
}

func TestGetServiceHealth(t *testing.T) {
	handler := getServiceHealth(newRawdataTarget(t))

	_, out, err := handler(context.Background(), nil, getServiceHealthInput{Name: "whoami@docker"})
	require.NoError(t, err)
	assert.Equal(t, "whoami@docker", out.Name)
	assert.Equal(t, 1, out.Up)
	assert.Equal(t, 0, out.Down)
	assert.True(t, out.Healthy)

	_, _, err = handler(context.Background(), nil, getServiceHealthInput{Name: "ghost@docker"})
	require.Error(t, err)
}

func TestListMiddlewares(t *testing.T) {
	handler := listMiddlewares(newRawdataTarget(t))

	_, out, err := handler(context.Background(), nil, listMiddlewaresInput{})
	require.NoError(t, err)
	require.Len(t, out.Middlewares, 1)

	mw := out.Middlewares[0]
	assert.Equal(t, "auth@docker", mw.Name)
	assert.Equal(t, "enabled", mw.Status)
	assert.Contains(t, mw.Types, "basicAuth")
	assert.Equal(t, []string{"web@docker"}, mw.UsedBy)
}

func TestGetConfigHash(t *testing.T) {
	handler := getConfigHash(newRawdataTarget(t))

	_, out, err := handler(context.Background(), nil, getConfigHashInput{})
	require.NoError(t, err)

	assert.Len(t, out.Hash, 64) // sha256 hex
	assert.Equal(t, 2, out.Routers)
	assert.Equal(t, 1, out.Services)
	assert.Equal(t, 1, out.Middlewares)

	// Same snapshot -> same hash.
	_, out2, err := handler(context.Background(), nil, getConfigHashInput{})
	require.NoError(t, err)
	assert.Equal(t, out.Hash, out2.Hash)
}

func TestGetConfigHash_ChangesWithConfig(t *testing.T) {
	base := getConfigHash(newRawdataTarget(t))
	_, out, err := base(context.Background(), nil, getConfigHashInput{})
	require.NoError(t, err)

	// A snapshot with one extra router must produce a different hash.
	mutated := &fakeTarget{responses: map[string]json.RawMessage{
		"/api/rawdata": json.RawMessage(`{"routers":{"x@docker":{"rule":"Host(` + "`x`" + `)","service":"x","status":"enabled"}}}`),
	}}
	_, out2, err := getConfigHash(mutated)(context.Background(), nil, getConfigHashInput{})
	require.NoError(t, err)

	assert.NotEqual(t, out.Hash, out2.Hash)
}

func TestGetEntryPoints(t *testing.T) {
	target := newFixtureTarget(t, map[string]string{"/api/entrypoints": "entrypoints.json"})
	handler := getEntryPoints(target)

	_, out, err := handler(context.Background(), nil, getEntryPointsInput{})
	require.NoError(t, err)
	require.Len(t, out.EntryPoints, 2)
	assert.Equal(t, "web", out.EntryPoints[0].Name)
	assert.True(t, out.EntryPoints[1].DefaultTLS)
}

func TestGetOverview(t *testing.T) {
	target := newFixtureTarget(t, map[string]string{"/api/overview": "overview.json"})
	handler := getOverview(target)

	_, out, err := handler(context.Background(), nil, getOverviewInput{})
	require.NoError(t, err)
	assert.Equal(t, 2, out.HTTP.Routers.Total)
	assert.Equal(t, []string{"docker", "file"}, out.Providers)
}
