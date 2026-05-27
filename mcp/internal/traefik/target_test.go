package traefik

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHTTPTarget_Get(t *testing.T) {
	testCases := []struct {
		desc       string
		path       string
		status     int
		body       string
		wantErr    bool
		assertFunc func(t *testing.T, out map[string]any)
	}{
		{
			desc:   "decodes router JSON",
			path:   "/api/http/routers/my-router@docker",
			status: http.StatusOK,
			body:   `{"name":"my-router@docker","status":"enabled","rule":"Host(` + "`x`" + `)"}`,
			assertFunc: func(t *testing.T, out map[string]any) {
				t.Helper()
				assert.Equal(t, "my-router@docker", out["name"])
				assert.Equal(t, "enabled", out["status"])
			},
		},
		{
			desc:    "404 returns error",
			path:    "/api/http/routers/missing",
			status:  http.StatusNotFound,
			body:    `{"error":"not found"}`,
			wantErr: true,
		},
		{
			desc:    "500 returns error",
			path:    "/api/overview",
			status:  http.StatusInternalServerError,
			body:    `boom`,
			wantErr: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, test.path, r.URL.Path)
				assert.Equal(t, http.MethodGet, r.Method)
				w.WriteHeader(test.status)
				_, _ = w.Write([]byte(test.body))
			}))
			defer srv.Close()

			target := NewHTTPTarget("primary", srv.URL, srv.Client())

			var out map[string]any
			err := target.Get(context.Background(), test.path, &out)

			if test.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			test.assertFunc(t, out)
		})
	}
}

func TestHTTPTarget_Name(t *testing.T) {
	target := NewHTTPTarget("primary", "http://localhost:8080", http.DefaultClient)
	assert.Equal(t, "primary", target.Name())
}

func TestHTTPTarget_DecodesIntoTypedStruct(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{"version": "v3.7.1"})
	}))
	defer srv.Close()

	target := NewHTTPTarget("primary", srv.URL, srv.Client())

	var out struct {
		Version string `json:"version"`
	}
	require.NoError(t, target.Get(context.Background(), "/api/version", &out))
	assert.Equal(t, "v3.7.1", out.Version)
}
