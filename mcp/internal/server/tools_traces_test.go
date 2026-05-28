package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/traefik/traefik-mcp/internal/tempo"
)

func TestSearchTraces(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"traces":[{"traceID":"abc","rootServiceName":"traefik","durationMs":2003}]}`))
	}))
	defer srv.Close()

	_, out, err := searchTraces(tempo.New(srv.URL, nil))(context.Background(), nil, searchTracesInput{Query: "{duration>1s}"})
	require.NoError(t, err)
	require.Len(t, out.Traces, 1)
	assert.Equal(t, "abc", out.Traces[0].TraceID)
}

func TestGetTrace(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"batches":[{"resource":{"attributes":[{"key":"service.name","value":{"stringValue":"traefik"}}]},"scopeSpans":[{"spans":[{"spanId":"s1","name":"EntryPoint web","startTimeUnixNano":"0","endTimeUnixNano":"5000000"}]}]}]}`))
	}))
	defer srv.Close()

	_, out, err := getTrace(tempo.New(srv.URL, nil))(context.Background(), nil, getTraceInput{TraceID: "abc"})
	require.NoError(t, err)
	require.Len(t, out.Spans, 1)
	assert.Equal(t, "s1", out.Spans[0].SpanID)
	assert.Equal(t, "traefik", out.Spans[0].Service)
}

func TestTraceToolsNotConfigured(t *testing.T) {
	_, _, err := searchTraces(nil)(context.Background(), nil, searchTracesInput{})
	require.ErrorIs(t, err, errNoTempo)

	_, _, err = getTrace(nil)(context.Background(), nil, getTraceInput{TraceID: "abc"})
	require.ErrorIs(t, err, errNoTempo)
}
