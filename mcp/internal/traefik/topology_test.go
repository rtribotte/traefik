package traefik

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func fixtureTarget(t *testing.T, path, file string) Target {
	t.Helper()
	body, err := os.ReadFile(file)
	require.NoError(t, err)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, path, r.URL.Path)
		_, _ = w.Write(body)
	}))
	t.Cleanup(srv.Close)

	return NewHTTPTarget("primary", srv.URL, srv.Client())
}

func TestFetchEntryPoints(t *testing.T) {
	target := fixtureTarget(t, "/api/entrypoints", "testdata/entrypoints.json")

	eps, err := FetchEntryPoints(context.Background(), target)
	require.NoError(t, err)
	require.Len(t, eps, 2)

	assert.Equal(t, "web", eps[0].Name)
	assert.Equal(t, ":80", eps[0].Address)
	assert.True(t, eps[0].AsDefault)

	assert.Equal(t, "websecure", eps[1].Name)
	assert.True(t, eps[1].DefaultTLS)
}

func TestFetchOverview(t *testing.T) {
	target := fixtureTarget(t, "/api/overview", "testdata/overview.json")

	ov, err := FetchOverview(context.Background(), target)
	require.NoError(t, err)

	assert.Equal(t, 2, ov.HTTP.Routers.Total)
	assert.Equal(t, 1, ov.HTTP.Routers.Warnings)
	assert.Equal(t, "Prometheus", ov.Features.Metrics)
	assert.True(t, ov.Features.AccessLog)
	assert.Equal(t, []string{"docker", "file"}, ov.Providers)
}
