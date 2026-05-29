package loki

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

func TestQueryRange(t *testing.T) {
	var gotQuery url.Values
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/loki/api/v1/query_range", r.URL.Path)
		gotQuery = r.URL.Query()
		// Two streams, out of timestamp order, to verify merge + chronological sort.
		_, _ = w.Write([]byte(`{"status":"success","data":{"resultType":"streams","result":[
			{"stream":{"a":"1"},"values":[["300","third"],["100","first"]]},
			{"stream":{"a":"2"},"values":[["200","second"]]}
		]}}`))
	}))
	defer srv.Close()

	lines, err := New(srv.URL, nil).QueryRange(context.Background(), `{service_name="traefik"}`, 50, time.Hour)
	require.NoError(t, err)
	assert.Equal(t, []string{"first", "second", "third"}, lines)

	assert.Equal(t, `{service_name="traefik"}`, gotQuery.Get("query"))
	assert.Equal(t, "50", gotQuery.Get("limit"))
	assert.Equal(t, "backward", gotQuery.Get("direction"))
	assert.NotEmpty(t, gotQuery.Get("start"))
	assert.NotEmpty(t, gotQuery.Get("end"))
}

func TestQueryRangeDefaults(t *testing.T) {
	var gotQuery url.Values
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.Query()
		_, _ = w.Write([]byte(`{"status":"success","data":{"result":[]}}`))
	}))
	defer srv.Close()

	_, err := New(srv.URL, nil).QueryRange(context.Background(), "", 0, 0)
	require.NoError(t, err)
	assert.Equal(t, DefaultQuery, gotQuery.Get("query"))
	assert.Equal(t, "100", gotQuery.Get("limit"))
}

func TestQueryRangeErrorStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("parse error"))
	}))
	defer srv.Close()

	_, err := New(srv.URL, nil).QueryRange(context.Background(), "{bad", 10, time.Hour)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "parse error")
}
