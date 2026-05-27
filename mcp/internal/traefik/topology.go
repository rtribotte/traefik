package traefik

import (
	"context"
	"encoding/json"
)

// EntryPoint is a lean projection of an entry point from GET /api/entrypoints.
// The full static.EntryPoint is intentionally not reused here: its schema is
// large and most fields are noise for an operator asking "what is listening".
type EntryPoint struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	AsDefault  bool   `json:"asDefault"`
	DefaultTLS bool   `json:"defaultTLS"`
}

// FetchEntryPoints retrieves the configured entry points.
func FetchEntryPoints(ctx context.Context, target Target) ([]EntryPoint, error) {
	var raw []struct {
		Name      string `json:"name"`
		Address   string `json:"address"`
		AsDefault bool   `json:"asDefault"`
		HTTP      struct {
			TLS *json.RawMessage `json:"tls"`
		} `json:"http"`
	}
	if err := target.Get(ctx, "/api/entrypoints", &raw); err != nil {
		return nil, err
	}

	eps := make([]EntryPoint, 0, len(raw))
	for _, r := range raw {
		eps = append(eps, EntryPoint{
			Name:       r.Name,
			Address:    r.Address,
			AsDefault:  r.AsDefault,
			DefaultTLS: r.HTTP.TLS != nil,
		})
	}
	return eps, nil
}

// Section counts resources by health for a kind (routers/services/middlewares).
type Section struct {
	Total    int `json:"total"`
	Warnings int `json:"warnings"`
	Errors   int `json:"errors"`
}

// SchemeOverview groups the counts for one protocol scheme.
type SchemeOverview struct {
	Routers     Section `json:"routers"`
	Services    Section `json:"services"`
	Middlewares Section `json:"middlewares"`
}

// Features reports which observability features are enabled.
type Features struct {
	Tracing   string `json:"tracing"`
	Metrics   string `json:"metrics"`
	AccessLog bool   `json:"accessLog"`
}

// Overview is the decoded GET /api/overview payload.
type Overview struct {
	HTTP      SchemeOverview `json:"http"`
	TCP       SchemeOverview `json:"tcp"`
	UDP       SchemeOverview `json:"udp"`
	Features  Features       `json:"features"`
	Providers []string       `json:"providers,omitempty"`
}

// FetchOverview retrieves the dashboard overview counters.
func FetchOverview(ctx context.Context, target Target) (*Overview, error) {
	var ov Overview
	if err := target.Get(ctx, "/api/overview", &ov); err != nil {
		return nil, err
	}
	return &ov, nil
}
