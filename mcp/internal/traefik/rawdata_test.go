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

func TestFetchRawData(t *testing.T) {
	body, err := os.ReadFile("testdata/rawdata.json")
	require.NoError(t, err)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/rawdata", r.URL.Path)
		_, _ = w.Write(body)
	}))
	defer srv.Close()

	target := NewHTTPTarget("primary", srv.URL, srv.Client())

	raw, err := FetchRawData(context.Background(), target)
	require.NoError(t, err)

	t.Run("routers decoded with rule and status", func(t *testing.T) {
		require.Contains(t, raw.Routers, "web@docker")
		web := raw.Routers["web@docker"]
		assert.Equal(t, "Host(`whoami.localhost`)", web.Rule)
		assert.Equal(t, "whoami", web.Service)
		assert.Equal(t, "enabled", web.Status)
	})

	t.Run("router errors decoded", func(t *testing.T) {
		broken := raw.Routers["broken@docker"]
		require.NotNil(t, broken)
		assert.Equal(t, "warning", broken.Status)
		assert.Contains(t, broken.Err, `the service "missing@docker" does not exist`)
	})

	t.Run("service serverStatus decoded", func(t *testing.T) {
		svc := raw.Services["whoami@docker"]
		require.NotNil(t, svc)
		assert.Equal(t, "UP", svc.ServerStatus["http://172.17.0.2:80"])
		require.NotNil(t, svc.LoadBalancer)
		require.Len(t, svc.LoadBalancer.Servers, 1)
		assert.Equal(t, "http://172.17.0.2:80", svc.LoadBalancer.Servers[0].URL)
	})

	t.Run("middleware decoded", func(t *testing.T) {
		mw := raw.Middlewares["auth@docker"]
		require.NotNil(t, mw)
		assert.Equal(t, "enabled", mw.Status)
		require.NotNil(t, mw.BasicAuth)
	})
}
