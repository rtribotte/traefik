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
  scripts/gen-certs.sh      short-lived TLS cert for the audit scenario
  demo.sh                   launch helper
  scenarios/
    01-router-missing/      router in warning + router never created
    02-service-5xx/         backend that goes down -> 502
    03-security-audit/      public-no-TLS route + expiring cert
    04-secured-route/       advisory config generation (no extra infra)
    05-slow-backend/        slow backend, diagnosed from access-log durations
    06-config-error/        unresolved refs, diagnosed from the app log
    07-cert-expiry/         expiring cert, surfaced via the metrics endpoint
    08-trace-latency/       slow request, found in the Tempo distributed traces
```

## Usage

```bash
./demo.sh up                     # base stack only
./demo.sh up 01-router-missing   # base + a scenario overlay
./demo.sh ps  01-router-missing
./demo.sh logs
./demo.sh down 01-router-missing
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
        "--tempo.url=http://localhost:3200"
      ]
    }
  }
}
```

Each scenario's README lists the prompts to try and what the assistant should
conclude.
