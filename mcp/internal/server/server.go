// Package server wires the Traefik-backed tools, resources and prompts onto an
// MCP server.
package server

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/traefik/traefik-mcp/internal/loki"
	"github.com/traefik/traefik-mcp/internal/prom"
	"github.com/traefik/traefik-mcp/internal/reference"
	"github.com/traefik/traefik-mcp/internal/staticconf"
	"github.com/traefik/traefik-mcp/internal/tempo"
	"github.com/traefik/traefik-mcp/internal/traefik"
)

// Deps holds the collaborators the MCP surface is built on.
type Deps struct {
	Target traefik.Target
	// AccessLogPath is the host path to Traefik's JSON access log, enabling the
	// access-log tool. Empty disables it at call time with a helpful error.
	AccessLogPath string
	// AppLogPath is the host path to Traefik's JSON application log, enabling the
	// app-log tool. Empty disables it at call time with a helpful error.
	AppLogPath string
	// Tempo queries distributed traces. Nil disables the trace tools at call
	// time with a helpful error.
	Tempo *tempo.Client
	// Loki queries access logs Traefik ships over OTLP. Nil disables the
	// access-log query tool at call time with a helpful error.
	Loki *loki.Client
	// Prom queries metrics Traefik ships over OTLP. Nil disables the metrics
	// query tool at call time with a helpful error.
	Prom *prom.Client
	// Reference is the embedded Traefik configuration reference
	// (github.com/traefik/reference). It backs the grounding tools (search,
	// concept, schema, doc) and schema validation. Nil disables those tools at
	// call time with a helpful error.
	Reference *reference.Catalogue
	// Caps are the observability capabilities deduced from Traefik's static
	// configuration. When set, data-source-backed tools (metrics, traces, logs)
	// are registered only if the matching source is configured. Nil means the
	// static configuration is unknown, so every tool is registered.
	Caps *staticconf.Capabilities
}

// has reports whether a capability-gated tool should be registered. With no
// known capabilities every tool is registered; otherwise have decides.
func (d Deps) has(have func(staticconf.Capabilities) bool) bool {
	if d.Caps == nil {
		return true
	}
	return have(*d.Caps)
}

// instructions guide the client to treat Traefik's configuration as live state.
// Without this, models answer follow-up questions from earlier tool results even
// after the configuration changed under them.
const instructions = `This server exposes a live, read-only view of a running Traefik instance.

The configuration is dynamic: routers, services and middlewares can change between
your turns (a new container, an edited config file). Never answer a question about
the current state from earlier tool results — call the relevant tool again.

Every list tool returns a "configHash" fingerprint of the configuration. If it
differs from the one you saw previously, the configuration changed and any cached
understanding is stale. Use get_config_hash for a cheap explicit check.

When diagnosing a router or service, always re-fetch its current state first.`

// New builds an MCP server with the read-only Traefik tools registered.
func New(name, version string, deps Deps) *mcp.Server {
	s := mcp.NewServer(
		&mcp.Implementation{Name: name, Version: version},
		&mcp.ServerOptions{Instructions: instructions},
	)
	addReadTools(s, deps.Target)

	if deps.has(func(c staticconf.Capabilities) bool { return c.FileAccessLog }) {
		mcp.AddTool(s, &mcp.Tool{
			Name: "tail_access_logs",
			Description: "Return recent entries from Traefik's access log, newest last. All filters are " +
				"optional and combine with AND: exact status, status range (minStatus/maxStatus), " +
				"service, router, host, method, path substring, and minDurationMs. Use it for any " +
				"traffic question — errors, slow requests, which router/service served a host, " +
				"traffic to a path, requests by method, etc.",
		}, tailAccessLogs(deps.AccessLogPath))
	}

	if deps.has(func(c staticconf.Capabilities) bool { return c.FileAppLog }) {
		mcp.AddTool(s, &mcp.Tool{
			Name: "tail_traefik_logs",
			Description: "Return recent entries from Traefik's application log, newest last. This is " +
				"where Traefik reports configuration and runtime errors (unresolved middleware/service " +
				"references, TLS/certificate problems, provider connection failures, invalid config). " +
				"Filters are optional and combine with AND: level (exact), minLevel (e.g. warn for " +
				"warnings and errors), and contains (substring in the message or error).",
		}, tailAppLogs(deps.AppLogPath))
	}

	mcp.AddTool(s, &mcp.Tool{
		Name: "diagnose_router_missing",
		Description: "Return a step-by-step playbook for diagnosing why a Traefik router is " +
			"missing or not routing traffic. Follow the returned steps using the live read tools.",
	}, diagnoseRouterMissingTool)

	mcp.AddTool(s, &mcp.Tool{
		Name: "diagnose_5xx",
		Description: "Return a step-by-step playbook for diagnosing 5xx errors and determining " +
			"whether the fault is Traefik routing, the service config, or the backend app. " +
			"Follow the returned steps using the live read tools.",
	}, diagnose5xxTool)

	if deps.has(func(c staticconf.Capabilities) bool { return c.PrometheusMetrics }) {
		mcp.AddTool(s, &mcp.Tool{
			Name: "get_metrics",
			Description: "Return Traefik's raw Prometheus /metrics output as text: config reload " +
				"status, request/response counts and durations by entrypoint and service, open " +
				"connections, TLS certificate expiry, and more.",
		}, getMetrics(deps.Target))
	}

	if deps.has(func(c staticconf.Capabilities) bool { return c.OTLPTracing }) {
		mcp.AddTool(s, &mcp.Tool{
			Name: "search_traces",
			Description: "Search Traefik's distributed traces in Tempo with an optional TraceQL " +
				"filter, returning matching trace summaries (trace ID, root service and operation, " +
				"duration). Use it to find slow requests ({duration>1s}), failed ones ({status=error}), " +
				"or recent traffic, then fetch a specific trace with get_trace.",
		}, searchTraces(deps.Tempo))

		mcp.AddTool(s, &mcp.Tool{
			Name: "get_trace",
			Description: "Return the full span tree of one trace by ID (from search_traces): each " +
				"span's service, operation, parent, duration, status and attributes. Use it to see " +
				"where time went or which hop failed across the entrypoint, router, middleware and " +
				"service spans of a request.",
		}, getTrace(deps.Tempo))
	}

	if deps.has(func(c staticconf.Capabilities) bool { return c.OTLPAccessLog }) {
		mcp.AddTool(s, &mcp.Tool{
			Name: "query_access_logs",
			Description: "Return recent Traefik access-log entries from Loki (otel-lgtm), newest " +
				"last, when Traefik ships access logs over OTLP. Same filters as tail_access_logs " +
				"(status, status range, service, router, host, method, path, minDurationMs), plus " +
				"lookbackMinutes for the time window. Use it for any traffic question — errors, slow " +
				"requests, which router/service served a host — when logs go to otel-lgtm.",
		}, queryAccessLogs(deps.Loki))
	}

	if deps.has(func(c staticconf.Capabilities) bool { return c.OTLPMetrics }) {
		mcp.AddTool(s, &mcp.Tool{
			Name: "query_metrics",
			Description: "Run a PromQL query against the Prometheus that otel-lgtm fills from " +
				"Traefik's OTLP metrics. Unlike get_metrics (a raw current scrape), this gives " +
				"history and aggregation: rates, sums and quantiles over time. Common metrics: " +
				"traefik_service_requests_total, traefik_service_request_duration_seconds_bucket, " +
				"traefik_entrypoint_requests_total, traefik_config_reloads_total, " +
				"traefik_open_connections. Set rangeMinutes for a range query (series); omit it " +
				"for an instant query (samples).",
		}, queryMetrics(deps.Prom))
	}

	mcp.AddTool(s, &mcp.Tool{
		Name: "search_traefik_docs",
		Description: "Search the Traefik configuration reference (github.com/traefik/reference) for the " +
			"concept behind a question — middlewares, routers, services, providers, CRDs, annotations and " +
			"static configuration sections — returning each match's concept id, type name and summary. " +
			"This is the discovery step: take the id of the best match and pass it to get_traefik_concept, " +
			"get_traefik_schema or get_traefik_doc. Covers Traefik Proxy (oss), Traefik Hub (hub) and " +
			"vendored Gateway API / core Kubernetes (external).",
	}, searchDocs(deps.Reference))

	mcp.AddTool(s, &mcp.Tool{
		Name: "get_traefik_concept",
		Description: "Return the full reference page for a concept id (from search_traefik_docs): every " +
			"field with its type, default and description — the exact contract for that configuration " +
			"object. Ground on this before authoring or correcting configuration; field names, types and " +
			"enums are exact for the pinned Traefik version.",
	}, getConcept(deps.Reference))

	mcp.AddTool(s, &mcp.Tool{
		Name: "get_traefik_schema",
		Description: "Return the JSON Schema for a concept id (from search_traefik_docs), suitable for " +
			"structured generation or as a tool input schema, so a generated fragment cannot be " +
			"structurally invalid.",
	}, getSchema(deps.Reference))

	mcp.AddTool(s, &mcp.Tool{
		Name: "get_traefik_doc",
		Description: "Return the narrative documentation page for a concept id (from search_traefik_docs): " +
			"the prose guide with examples, resolved via the reference's DOC_INDEX. Use it when the field " +
			"contract from get_traefik_concept is not enough and you need explanation or examples.",
	}, getDoc(deps.Reference))

	mcp.AddTool(s, &mcp.Tool{
		Name: "validate_traefik_config",
		Description: "Validate Traefik configuration against the official JSON Schemas " +
			"(github.com/traefik/reference) and report each violation with its location. Accepts YAML or " +
			"JSON, including multi-document YAML. It auto-detects what you passed — a whole traefik.yaml " +
			"(static install or dynamic routing config), a Traefik or Gateway API CRD, or an annotated " +
			"Kubernetes manifest — by matching it against the schema registry. Pass 'concept' to validate " +
			"a single fragment (e.g. one middleware) against that concept's schema. Use it to vet " +
			"generated or hand-written configuration before applying it.",
	}, validateConfig(deps.Reference))

	//addPrompts(s)

	return s
}
