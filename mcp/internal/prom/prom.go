// Package prom queries Traefik's metrics from a Prometheus instance (as bundled
// in otel-lgtm) when Traefik exports them over OTLP. It speaks the Prometheus
// HTTP query API and projects the results into flat structures for the model.
package prom

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Client is a minimal Prometheus query API client.
type Client struct {
	baseURL string
	client  *http.Client
}

// New returns a Client for the given Prometheus base URL (e.g.
// http://localhost:9090).
func New(baseURL string, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{baseURL: strings.TrimRight(baseURL, "/"), client: client}
}

// Sample is a single value of an instant query result.
type Sample struct {
	Metric   map[string]string `json:"metric"`
	Value    float64           `json:"value"`
	TimeUnix float64           `json:"timeUnix"`
}

// Point is one (time, value) pair of a range query series.
type Point struct {
	TimeUnix float64 `json:"timeUnix"`
	Value    float64 `json:"value"`
}

// Series is one labelled time series of a range query result.
type Series struct {
	Metric map[string]string `json:"metric"`
	Points []Point           `json:"points"`
}

type apiResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
	Data   struct {
		ResultType string          `json:"resultType"`
		Result     json.RawMessage `json:"result"`
	} `json:"data"`
}

type rawVector struct {
	Metric map[string]string  `json:"metric"`
	Value  [2]json.RawMessage `json:"value"`
}

type rawMatrix struct {
	Metric map[string]string    `json:"metric"`
	Values [][2]json.RawMessage `json:"values"`
}

// Query runs an instant PromQL query and returns the resulting samples.
func (c *Client) Query(ctx context.Context, expr string) ([]Sample, error) {
	q := url.Values{}
	q.Set("query", expr)

	result, err := c.query(ctx, "/api/v1/query?"+q.Encode())
	if err != nil {
		return nil, err
	}

	var vectors []rawVector
	if err := json.Unmarshal(result, &vectors); err != nil {
		return nil, fmt.Errorf("decoding instant result: %w", err)
	}

	samples := make([]Sample, 0, len(vectors))
	for _, v := range vectors {
		ts, val := parsePair(v.Value)
		samples = append(samples, Sample{Metric: v.Metric, TimeUnix: ts, Value: val})
	}
	return samples, nil
}

// QueryRange runs a PromQL query over [start, end] at the given step and returns
// the resulting series. A non-positive step defaults to 15s.
func (c *Client) QueryRange(ctx context.Context, expr string, start, end time.Time, step time.Duration) ([]Series, error) {
	if step <= 0 {
		step = 15 * time.Second
	}

	q := url.Values{}
	q.Set("query", expr)
	q.Set("start", strconv.FormatInt(start.Unix(), 10))
	q.Set("end", strconv.FormatInt(end.Unix(), 10))
	q.Set("step", strconv.FormatFloat(step.Seconds(), 'f', -1, 64))

	result, err := c.query(ctx, "/api/v1/query_range?"+q.Encode())
	if err != nil {
		return nil, err
	}

	var matrices []rawMatrix
	if err := json.Unmarshal(result, &matrices); err != nil {
		return nil, fmt.Errorf("decoding range result: %w", err)
	}

	series := make([]Series, 0, len(matrices))
	for _, m := range matrices {
		s := Series{Metric: m.Metric, Points: make([]Point, 0, len(m.Values))}
		for _, v := range m.Values {
			ts, val := parsePair(v)
			s.Points = append(s.Points, Point{TimeUnix: ts, Value: val})
		}
		series = append(series, s)
	}
	return series, nil
}

// parsePair decodes a Prometheus [<unix float>, "<string value>"] pair.
func parsePair(pair [2]json.RawMessage) (ts, val float64) {
	_ = json.Unmarshal(pair[0], &ts)
	var s string
	if json.Unmarshal(pair[1], &s) == nil {
		val, _ = strconv.ParseFloat(s, 64)
	}
	return ts, val
}

func (c *Client) query(ctx context.Context, path string) (json.RawMessage, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var out apiResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("Prometheus returned %s: %s", resp.Status, body[:min(len(body), 1024)])
	}
	if out.Status != "success" {
		return nil, fmt.Errorf("Prometheus query failed: %s", out.Error)
	}
	return out.Data.Result, nil
}
