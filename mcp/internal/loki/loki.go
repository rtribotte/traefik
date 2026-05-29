// Package loki queries Traefik's access logs from a Loki instance (as bundled in
// otel-lgtm) when Traefik ships them over OTLP. The log line bodies are the same
// JSON access-log records Traefik writes to a file, so callers can parse them
// with the internal/logs access parser.
package loki

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"time"
)

// Client is a minimal Loki query API client.
type Client struct {
	baseURL string
	client  *http.Client
}

// New returns a Client for the given Loki base URL (e.g. http://localhost:3100).
func New(baseURL string, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{baseURL: strings.TrimRight(baseURL, "/"), client: client}
}

// DefaultQuery selects Traefik's logs; Traefik ships them under the OTLP
// resource attribute service.name=traefik, which Loki exposes as service_name.
const DefaultQuery = `{service_name="traefik"}`

type queryRangeResponse struct {
	Data struct {
		Result []struct {
			Values [][2]string `json:"values"`
		} `json:"result"`
	} `json:"data"`
}

// QueryRange returns the line bodies of up to limit log entries matching the
// LogQL query within the last lookback window, in chronological order (oldest
// first). An empty query selects Traefik's logs; a non-positive limit defaults
// to 100 and a non-positive lookback to one hour.
func (c *Client) QueryRange(ctx context.Context, query string, limit int, lookback time.Duration) ([]string, error) {
	if query == "" {
		query = DefaultQuery
	}
	if limit <= 0 {
		limit = 100
	}
	if lookback <= 0 {
		lookback = time.Hour
	}

	end := time.Now()
	start := end.Add(-lookback)

	q := url.Values{}
	q.Set("query", query)
	q.Set("limit", strconv.Itoa(limit))
	q.Set("start", strconv.FormatInt(start.UnixNano(), 10))
	q.Set("end", strconv.FormatInt(end.UnixNano(), 10))
	q.Set("direction", "backward") // newest first, so limit keeps the most recent.

	body, err := c.get(ctx, "/loki/api/v1/query_range?"+q.Encode())
	if err != nil {
		return nil, err
	}

	var resp queryRangeResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("decoding Loki response: %w", err)
	}

	// Merge all streams and sort by nanosecond timestamp ascending so the result
	// reads oldest-to-newest regardless of Loki's per-stream grouping.
	type tsLine struct {
		ts   int64
		line string
	}
	var entries []tsLine
	for _, stream := range resp.Data.Result {
		for _, v := range stream.Values {
			ts, _ := strconv.ParseInt(v[0], 10, 64)
			entries = append(entries, tsLine{ts: ts, line: v[1]})
		}
	}
	slices.SortFunc(entries, func(a, b tsLine) int { return int(a.ts - b.ts) })

	lines := make([]string, len(entries))
	for i, e := range entries {
		lines[i] = e.line
	}
	return lines, nil
}

func (c *Client) get(ctx context.Context, path string) ([]byte, error) {
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
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Loki returned %s: %s", resp.Status, body[:min(len(body), 1024)])
	}
	return body, nil
}
