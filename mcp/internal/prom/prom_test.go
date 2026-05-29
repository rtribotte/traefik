package prom

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQuery(t *testing.T) {
	var gotQuery url.Values
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/query", r.URL.Path)
		gotQuery = r.URL.Query()
		_, _ = w.Write([]byte(`{"status":"success","data":{"resultType":"vector","result":[
			{"metric":{"__name__":"traefik_service_requests_total","code":"200"},"value":[1780000000,"42"]}
		]}}`))
	}))
	defer srv.Close()

	samples, err := New(srv.URL, nil).Query(context.Background(), "traefik_service_requests_total")
	require.NoError(t, err)
	require.Len(t, samples, 1)
	assert.Equal(t, "200", samples[0].Metric["code"])
	assert.Equal(t, 42.0, samples[0].Value)
	assert.Equal(t, 1780000000.0, samples[0].TimeUnix)
	assert.Equal(t, "traefik_service_requests_total", gotQuery.Get("query"))
}

func TestQueryRange(t *testing.T) {
	var gotQuery url.Values
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/query_range", r.URL.Path)
		gotQuery = r.URL.Query()
		_, _ = w.Write([]byte(`{"status":"success","data":{"resultType":"matrix","result":[
			{"metric":{"code":"200"},"values":[[1780000000,"1"],[1780000015,"3"]]}
		]}}`))
	}))
	defer srv.Close()

	start := time.Unix(1780000000, 0)
	end := time.Unix(1780000015, 0)
	series, err := New(srv.URL, nil).QueryRange(context.Background(), "rate(x[1m])", start, end, 15*time.Second)
	require.NoError(t, err)
	require.Len(t, series, 1)
	require.Len(t, series[0].Points, 2)
	assert.Equal(t, 3.0, series[0].Points[1].Value)
	assert.Equal(t, "15", gotQuery.Get("step"))
}

func TestQueryError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"status":"error","error":"bad expression"}`))
	}))
	defer srv.Close()

	_, err := New(srv.URL, nil).Query(context.Background(), "((")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "bad expression")
}
