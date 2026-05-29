# traefik-mcp demo environment

A local Traefik + whoami stack for exercising the MCP server, plus per-scenario
overlays that introduce the specific breakage each demo needs.

## Layout

```
demo/
  docker-compose.yml        base stack: Traefik (api.insecure, prometheus,
                            file access log, OTLP tracing) + whoami + otel-lgtm
                            (Tempo trace store, query API on :3200)
  config/
    traefik.yml             static (install) configuration
    dynamic/                file-provider directory (watched)
  scripts/gen-certs.sh      short-lived TLS cert for the security scenario
  demo.sh                   launch helper
  scenarios/
    01-broken-route/        v2-syntax router fails to register -> read the parse
                            error off the router, look up the v3 rule, fix, validate
    02-security-posture/    public-no-TLS route + expiring cert -> audit with
                            list_certificates, harden with the reference, validate
    03-trace-latency/       slow request, found in the Tempo distributed traces
    04-shadowed-route/      valid router served by the wrong backend -> a
                            high-priority catch-all wins; needs the reference on
                            router priority, then validate
    05-degraded-backend/    intermittent 502s while health says UP -> the failure
                            rate and trend live in the metrics (query_metrics)
```

The scenarios each force a different evidence source — the application log
(01), `list_certificates` (02), distributed traces (03), router priority +
reference (04), metrics history (05) — past the point where reading a single
router status or log line answers the question. The config-fault ones (01, 02,
04) close the loop through the embedded reference (look up the fix) and
`validate_traefik_config` (prove it before applying).

## Usage

```bash
./demo.sh up                     # base stack only
./demo.sh up 01-broken-route     # base + a scenario overlay
./demo.sh ps  01-broken-route
./demo.sh logs
./demo.sh down 01-broken-route
./demo.sh list                   # list scenarios
```

Traefik API/dashboard: http://localhost:8080 — this is what the MCP server reads.

`*.localhost` hosts resolve to loopback; reach a service with e.g.
`curl -H 'Host: whoami.localhost' http://localhost/`.

## Point the MCP server at it

```json
{
  "mcpServers": {
    "traefik": {
      "command": "/Users/romain/go/src/github.com/traefik/traefik/mcp/bin/traefik-mcp",
      "args": [
        "--traefik.api-url=http://localhost:8088",
        "--traefik.access-log=/Users/romain/go/src/github.com/traefik/traefik/mcp/demo/logs/access.log",
        "--traefik.app-log=/Users/romain/go/src/github.com/traefik/traefik/mcp/demo/logs/traefik.log",
        "--tempo.url=http://localhost:3200",
        "--loki.url=http://localhost:3100",
        "--prometheus.url=http://localhost:9090"
      ]
    }
  }
}
```

On startup the server reads `--traefik.app-log` for the line where Traefik dumps
its static configuration (`Static configuration loaded [json]`) and registers
only the tools whose data sources are configured — metrics, traces and the log
tools. Traefik logs that line only at `log.level=DEBUG` (the demo config sets
it); without it the server falls back to registering every tool.

The demo Traefik exports metrics and access logs over OTLP to otel-lgtm *as well
as* the Prometheus pull endpoint and the file logs, so two extra tools surface:
`query_metrics` (PromQL against lgtm's Prometheus, `--prometheus.url`) and
`query_access_logs` (Traefik's access logs from lgtm's Loki, `--loki.url`).
These complement `get_metrics` (raw current scrape) and `tail_access_logs` (file
tail) with history, aggregation and time-windowed queries.

Five more tools need no live Traefik and are always available, backed by the whole
[github.com/traefik/reference](https://github.com/traefik/reference) catalogue vendored and
embedded in the binary. `search_traefik_docs` finds the concept behind a question and returns
its id; `get_traefik_concept` returns that concept's full field contract (types, defaults,
descriptions); `get_traefik_schema` returns its JSON Schema; `get_traefik_doc` resolves its
narrative documentation page. `validate_traefik_config` checks a YAML/JSON configuration
against the official JSON Schemas by matching it to the right one in the registry — a whole
traefik.yaml (static or dynamic), a CRD, an annotated manifest, or a single concept fragment —
and reports each violation with its location, for vetting generated or hand-written
configuration before applying it.

Each scenario's README lists the prompts to try and what the assistant should
conclude.
