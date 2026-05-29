# Scenario 03 — "find the slow checkout request in the traces"

Nothing looks broken: `checkout.localhost` routes correctly, the backend
(whoami) is UP and fast, and every request returns 200. Yet requests are slow,
by a random amount. The cause is a local Yaegi plugin middleware
(`pluginA@file`) that sleeps a random duration on every request.

The delay is **not configurable** — the plugin takes no options — and its name
gives nothing away, so it cannot be inferred from the middleware configuration.
It can only be found in the **distributed traces** Traefik ships to Tempo
(otel-lgtm): the time appears in Traefik's middleware chain, not the backend
service.

## Run

```bash
./demo.sh up 03-trace-latency      # brings up the stack and runs the traffic

./demo.sh traffic 03-trace-latency # replay the traffic any time
```

`up` sends 20 requests once Traefik is ready; each is delayed a random amount by
the plugin, so Tempo holds a spread of durations. Allow a few seconds for
ingestion before querying.

## Demo in Claude Desktop

- "checkout.localhost is slow but I can't see why — the config and the backend
  look fine. Can you find a slow request in the traces and tell me where the
  time is going?"

Expected: the assistant inspects the router/service/middleware config and the
backend health and finds nothing wrong — in particular the `pluginA`
middleware has no options and its name says nothing, so the config gives no
hint. It calls `search_traces` with
`{duration>1s}`, picks a slow trace, and `get_trace` on its ID. It sees two
spans: the entrypoint span (`GET`) lasting ~2s, and the backend span
(`ReverseProxy`) lasting only a few milliseconds and starting ~2s into the
request. The gap means the time is spent in Traefik's middleware chain before
the backend is ever called — i.e. a middleware, not the backend or the network.
The culprit is the `pluginA` middleware, invisible in its configuration but
plain in the trace.

> Point the MCP server at Tempo so the assistant can query traces: add
> `--tempo.url=http://localhost:3200` to its args in the Claude Desktop config.

## Reset

```bash
./demo.sh down 03-trace-latency
```
