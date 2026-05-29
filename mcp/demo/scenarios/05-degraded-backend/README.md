# Scenario 05 — "api.localhost fails about half the time, but everything looks UP"

An intermittent failure that the obvious surfaces get wrong. The `api` service
load-balances across two servers; one points at a dead port, and with no health
check configured Traefik never evicts it, so roughly half the requests return
502 — intermittently, with no router or service in an error state. This is the
scenario where `get_service_health` actively *misleads* and the answer lives in
the **metrics**.

## Run

```bash
./demo.sh up 05-degraded-backend      # brings up the stack and sends the traffic

./demo.sh traffic 05-degraded-backend # replay the traffic any time
```

## What's wrong

`scenarios/05-degraded-backend/dynamic/api.yml` defines service `api@file` with two
load-balanced servers:

```
http://whoami:80   # healthy
http://whoami:81   # nothing listens here -> 502
```

There is no health check, so Traefik keeps round-robining onto the dead server
and ~50% of requests to `api.localhost` fail. Because server health is only
tracked when a health check is configured, `serverStatus` reports **both servers
`UP`** — so `get_service_health` says the service is healthy. It is not.

## Demo in Claude Desktop

- "Users report api.localhost fails intermittently. How often is it failing, and
  is it one backend or all of them?"

Expected flow:

1. **The obvious tools mislead or under-inform.** `get_service("api@file")` shows
   `status: enabled`; `get_service_health("api@file")` reports both servers `UP`
   and the service healthy (no health check means Traefik assumes up). A
   `tail_access_logs` snapshot shows a *few* scattered 502s mixed with 200s but
   can't tell you the rate, or whether it is steady or worsening.
2. **Quantify with metrics.** `query_metrics` over
   `traefik_service_requests_total{service="api@file"}` by `code` shows the split
   — about half `200`, half `502` — and a range query shows it is a steady ~50%,
   not a transient blip. (`get_metrics` carries the same counters in the current
   scrape; `query_metrics` adds the rate and the trend over time.)
3. **Locate the cause.** With the failure quantified and pinned to `api@file`,
   `get_service("api@file")` shows the two-server load balancer; one server URL
   (`http://whoami:81`) is the dead one. The fix is to correct or remove that
   server and add a health check so a future bad backend is evicted instead of
   silently serving 502s.

> Point the MCP server at otel-lgtm's Prometheus so the assistant can run the
> metric queries: the demo's `--prometheus.url=http://localhost:9090` already
> covers `query_metrics`, and `get_metrics` reads the Traefik scrape directly.

## Reset

```bash
./demo.sh down 05-degraded-backend
```
