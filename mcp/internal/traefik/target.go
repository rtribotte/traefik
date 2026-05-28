// Package traefik provides a typed, read-only client over the Traefik HTTP API.
package traefik

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Target is a read-only handle to a single Traefik instance's API.
// One implementation wraps a single base URL today; a fan-out implementation
// can satisfy the same interface later without changing tool code.
type Target interface {
	// Get performs GET base+path and decodes the JSON response into out.
	Get(ctx context.Context, path string, out any) error
	// GetRaw performs GET base+path and returns the raw response body, for
	// non-JSON endpoints such as Prometheus /metrics.
	GetRaw(ctx context.Context, path string) ([]byte, error)
	// Name identifies the target in multi-instance output.
	Name() string
}

// HTTPTarget is a Target backed by one Traefik API base URL.
type HTTPTarget struct {
	name    string
	baseURL string
	client  *http.Client
}

// NewHTTPTarget builds a Target for a single Traefik instance.
func NewHTTPTarget(name, baseURL string, client *http.Client) *HTTPTarget {
	if client == nil {
		client = http.DefaultClient
	}
	return &HTTPTarget{
		name:    name,
		baseURL: strings.TrimRight(baseURL, "/"),
		client:  client,
	}
}

func (t *HTTPTarget) Name() string { return t.name }

func (t *HTTPTarget) Get(ctx context.Context, path string, out any) error {
	body, err := t.GetRaw(ctx, path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("decoding response from %s: %w", path, err)
	}

	return nil
}

func (t *HTTPTarget) GetRaw(ctx context.Context, path string) ([]byte, error) {
	url := t.baseURL + path

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("building request for %s: %w", path, err)
	}

	resp, err := t.client.Do(req)
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
