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
	url := t.baseURL + path

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("building request for %s: %w", path, err)
	}

	resp, err := t.client.Do(req)
	if err != nil {
		return fmt.Errorf("requesting %s: %w", path, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		return fmt.Errorf("%s: unexpected status %d: %s", path, resp.StatusCode, strings.TrimSpace(string(body)))
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("decoding response from %s: %w", path, err)
	}

	return nil
}
