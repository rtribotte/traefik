package tempo

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	var gotQuery, gotLimit string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/search", r.URL.Path)
		gotQuery = r.URL.Query().Get("q")
		gotLimit = r.URL.Query().Get("limit")
		_, _ = w.Write([]byte(`{"traces":[
			{"traceID":"abc","rootServiceName":"traefik","rootTraceName":"EntryPoint web","startTimeUnixNano":"1","durationMs":2003},
			{"traceID":"def","rootServiceName":"traefik","rootTraceName":"EntryPoint web","startTimeUnixNano":"2","durationMs":5}
		]}`))
	}))
	defer srv.Close()

	out, err := New(srv.URL, nil).Search(context.Background(), `{duration>1s}`, 5)
	require.NoError(t, err)
	assert.Equal(t, `{duration>1s}`, gotQuery)
	assert.Equal(t, "5", gotLimit)
	require.Len(t, out, 2)
	assert.Equal(t, "abc", out[0].TraceID)
	assert.Equal(t, 2003, out[0].DurationMs)
}

func TestSearchDefaults(t *testing.T) {
	var gotQuery, gotLimit string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.Query().Get("q")
		gotLimit = r.URL.Query().Get("limit")
		_, _ = w.Write([]byte(`{"traces":[]}`))
	}))
	defer srv.Close()

	_, err := New(srv.URL, nil).Search(context.Background(), "", 0)
	require.NoError(t, err)
	assert.Equal(t, "{}", gotQuery)
	assert.Equal(t, "20", gotLimit)
}

func TestTrace(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/traces/abc", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Accept"))
		_, _ = w.Write([]byte(`{"batches":[
			{
			  "resource":{"attributes":[{"key":"service.name","value":{"stringValue":"traefik"}}]},
			  "scopeSpans":[{"spans":[
				{
				  "traceId":"abc","spanId":"s1","name":"EntryPoint web",
				  "kind":"SPAN_KIND_SERVER",
				  "startTimeUnixNano":"1000000","endTimeUnixNano":"2003000000",
				  "attributes":[
					{"key":"http.response.status_code","value":{"intValue":"502"}},
					{"key":"http.request.method","value":{"stringValue":"GET"}}
				  ],
				  "status":{"code":"STATUS_CODE_ERROR","message":"backend down"}
				},
				{
				  "traceId":"abc","spanId":"s2","parentSpanId":"s1","name":"reverse proxy",
				  "kind":3,
				  "startTimeUnixNano":"1000000","endTimeUnixNano":"1500000",
				  "status":{"code":2}
				}
			  ]}]
			}
		]}`))
	}))
	defer srv.Close()

	spans, err := New(srv.URL, nil).Trace(context.Background(), "abc")
	require.NoError(t, err)
	require.Len(t, spans, 2)

	assert.Equal(t, "s1", spans[0].SpanID)
	assert.Equal(t, "traefik", spans[0].Service)
	assert.Equal(t, "SPAN_KIND_SERVER", spans[0].Kind)
	assert.Equal(t, "STATUS_CODE_ERROR", spans[0].StatusCode)
	assert.Equal(t, "backend down", spans[0].StatusMessage)
	assert.InDelta(t, 2002.0, spans[0].DurationMs, 0.001)
	assert.Equal(t, "502", spans[0].Attributes["http.response.status_code"])
	assert.Equal(t, "GET", spans[0].Attributes["http.request.method"])

	assert.Equal(t, "s2", spans[1].SpanID)
	assert.Equal(t, "s1", spans[1].ParentSpanID)
	assert.Equal(t, "SPAN_KIND_CLIENT", spans[1].Kind)
	assert.Equal(t, "STATUS_CODE_ERROR", spans[1].StatusCode)
}

func TestTraceDecodesBase64IDsToHex(t *testing.T) {
	// Real Tempo OTLP-JSON encodes trace/span IDs as base64; the flattened
	// output should render them as hex so they match the search IDs.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"batches":[{"scopeSpans":[{"spans":[
			{"traceId":"ZtPOecsgrgn5qy9v5UWpow==","spanId":"+qim68W75sw=","parentSpanId":"zUVre3nNKJk=","name":"ReverseProxy","startTimeUnixNano":"0","endTimeUnixNano":"0"}
		]}]}]}`))
	}))
	defer srv.Close()

	spans, err := New(srv.URL, nil).Trace(context.Background(), "x")
	require.NoError(t, err)
	require.Len(t, spans, 1)
	assert.Equal(t, "66d3ce79cb20ae09f9ab2f6fe545a9a3", spans[0].TraceID)
	assert.Equal(t, "faa8a6ebc5bbe6cc", spans[0].SpanID)
	assert.Equal(t, "cd456b7b79cd2899", spans[0].ParentSpanID)
}

func TestSearchErrorStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("bad query"))
	}))
	defer srv.Close()

	_, err := New(srv.URL, nil).Search(context.Background(), "{", 5)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "bad query")
}
