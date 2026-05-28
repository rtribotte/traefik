// Package tempo provides a read-only client over a Tempo (otel-lgtm) HTTP API
// for searching distributed traces and fetching a trace's spans.
package tempo

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Client queries a single Tempo instance over HTTP.
type Client struct {
	baseURL string
	client  *http.Client
}

// New builds a Tempo client for one base URL (e.g. http://localhost:3200).
func New(baseURL string, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{baseURL: strings.TrimRight(baseURL, "/"), client: client}
}

// TraceSummary is one hit from a trace search.
type TraceSummary struct {
	TraceID           string `json:"traceID"`
	RootServiceName   string `json:"rootServiceName"`
	RootTraceName     string `json:"rootTraceName"`
	StartTimeUnixNano string `json:"startTimeUnixNano"`
	DurationMs        int    `json:"durationMs"`
}

// Span is a flattened view of one OTLP span, with its resource service name
// and attributes folded in as strings for easy consumption by a model.
type Span struct {
	TraceID       string            `json:"traceID"`
	SpanID        string            `json:"spanID"`
	ParentSpanID  string            `json:"parentSpanID,omitempty"`
	Name          string            `json:"name"`
	Service       string            `json:"service"`
	Kind          string            `json:"kind,omitempty"`
	StartUnixNano uint64            `json:"startUnixNano"`
	EndUnixNano   uint64            `json:"endUnixNano"`
	DurationMs    float64           `json:"durationMs"`
	StatusCode    string            `json:"statusCode,omitempty"`
	StatusMessage string            `json:"statusMessage,omitempty"`
	Attributes    map[string]string `json:"attributes,omitempty"`
}

// Search runs a TraceQL query (default "{}" matches everything) and returns up
// to limit trace summaries, most recent first as Tempo orders them.
func (c *Client) Search(ctx context.Context, query string, limit int) ([]TraceSummary, error) {
	if query == "" {
		query = "{}"
	}
	if limit <= 0 {
		limit = 20
	}

	q := url.Values{}
	q.Set("q", query)
	q.Set("limit", strconv.Itoa(limit))

	body, err := c.get(ctx, "/api/search?"+q.Encode(), false)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Traces []TraceSummary `json:"traces"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("decoding search response: %w", err)
	}
	return resp.Traces, nil
}

// Trace fetches one trace by ID and flattens it into a list of spans.
func (c *Client) Trace(ctx context.Context, id string) ([]Span, error) {
	body, err := c.get(ctx, "/api/traces/"+url.PathEscape(id), true)
	if err != nil {
		return nil, err
	}
	return flatten(body)
}

func (c *Client) get(ctx context.Context, path string, asJSON bool) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+path, nil)
	if err != nil {
		return nil, fmt.Errorf("building request for %s: %w", path, err)
	}
	// Tempo serves a trace as protobuf unless JSON is explicitly requested.
	if asJSON {
		req.Header.Set("Accept", "application/json")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("requesting %s: %w", path, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response from %s: %w", path, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: unexpected status %d: %s", path, resp.StatusCode, strings.TrimSpace(string(body[:min(len(body), 1024)])))
	}
	return body, nil
}

// otlpTrace mirrors the subset of Tempo's OTLP-JSON trace response we read.
type otlpTrace struct {
	Batches []struct {
		Resource struct {
			Attributes []otlpAttr `json:"attributes"`
		} `json:"resource"`
		ScopeSpans []struct {
			Spans []otlpSpan `json:"spans"`
		} `json:"scopeSpans"`
	} `json:"batches"`
}

type otlpSpan struct {
	TraceID           string          `json:"traceId"`
	SpanID            string          `json:"spanId"`
	ParentSpanID      string          `json:"parentSpanId"`
	Name              string          `json:"name"`
	Kind              json.RawMessage `json:"kind"`
	StartTimeUnixNano string          `json:"startTimeUnixNano"`
	EndTimeUnixNano   string          `json:"endTimeUnixNano"`
	Attributes        []otlpAttr      `json:"attributes"`
	Status            struct {
		Code    json.RawMessage `json:"code"`
		Message string          `json:"message"`
	} `json:"status"`
}

type otlpAttr struct {
	Key   string `json:"key"`
	Value struct {
		StringValue *string  `json:"stringValue"`
		IntValue    *string  `json:"intValue"`
		BoolValue   *bool    `json:"boolValue"`
		DoubleValue *float64 `json:"doubleValue"`
	} `json:"value"`
}

func (a otlpAttr) string() string {
	switch {
	case a.Value.StringValue != nil:
		return *a.Value.StringValue
	case a.Value.IntValue != nil:
		return *a.Value.IntValue
	case a.Value.BoolValue != nil:
		return strconv.FormatBool(*a.Value.BoolValue)
	case a.Value.DoubleValue != nil:
		return strconv.FormatFloat(*a.Value.DoubleValue, 'f', -1, 64)
	default:
		return ""
	}
}

func flatten(body []byte) ([]Span, error) {
	var trace otlpTrace
	if err := json.Unmarshal(body, &trace); err != nil {
		return nil, fmt.Errorf("decoding trace: %w", err)
	}

	var spans []Span
	for _, batch := range trace.Batches {
		service := ""
		for _, attr := range batch.Resource.Attributes {
			if attr.Key == "service.name" {
				service = attr.string()
			}
		}

		for _, scope := range batch.ScopeSpans {
			for _, s := range scope.Spans {
				start := parseUint(s.StartTimeUnixNano)
				end := parseUint(s.EndTimeUnixNano)

				attrs := make(map[string]string, len(s.Attributes))
				for _, attr := range s.Attributes {
					attrs[attr.Key] = attr.string()
				}

				spans = append(spans, Span{
					TraceID:       decodeID(s.TraceID),
					SpanID:        decodeID(s.SpanID),
					ParentSpanID:  decodeID(s.ParentSpanID),
					Name:          s.Name,
					Service:       service,
					Kind:          spanKind(s.Kind),
					StartUnixNano: start,
					EndUnixNano:   end,
					DurationMs:    float64(end-start) / 1e6,
					StatusCode:    statusCode(s.Status.Code),
					StatusMessage: s.Status.Message,
					Attributes:    attrs,
				})
			}
		}
	}
	return spans, nil
}

// decodeID renders an OTLP trace/span ID as hex. Tempo's OTLP-JSON encodes
// these IDs as base64; converting to hex makes them match the IDs from search
// and keeps parent/child links comparable. Non-base64 input is returned as is.
func decodeID(id string) string {
	if id == "" {
		return ""
	}
	raw, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		return id
	}
	return hex.EncodeToString(raw)
}

func parseUint(s string) uint64 {
	v, _ := strconv.ParseUint(s, 10, 64)
	return v
}

// spanKind renders the OTLP span kind, which Tempo may emit either as the enum
// name ("SPAN_KIND_SERVER") or its numeric value.
func spanKind(raw json.RawMessage) string {
	if len(raw) == 0 {
		return ""
	}
	var name string
	if err := json.Unmarshal(raw, &name); err == nil {
		return name
	}
	var n int
	if err := json.Unmarshal(raw, &n); err == nil {
		switch n {
		case 1:
			return "SPAN_KIND_INTERNAL"
		case 2:
			return "SPAN_KIND_SERVER"
		case 3:
			return "SPAN_KIND_CLIENT"
		case 4:
			return "SPAN_KIND_PRODUCER"
		case 5:
			return "SPAN_KIND_CONSUMER"
		}
	}
	return ""
}

// statusCode renders the OTLP status code, emitted either as the enum name
// ("STATUS_CODE_ERROR") or its numeric value.
func statusCode(raw json.RawMessage) string {
	if len(raw) == 0 {
		return ""
	}
	var name string
	if err := json.Unmarshal(raw, &name); err == nil {
		return name
	}
	var n int
	if err := json.Unmarshal(raw, &n); err == nil {
		switch n {
		case 1:
			return "STATUS_CODE_OK"
		case 2:
			return "STATUS_CODE_ERROR"
		default:
			return "STATUS_CODE_UNSET"
		}
	}
	return ""
}
