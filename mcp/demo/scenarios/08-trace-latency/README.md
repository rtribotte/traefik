# Scenario 08 — "find the slow checkout request in the traces"

Nothing is broken: `checkout.localhost` routes correctly, the backend is UP, and
every request returns 200. Most requests are fast; a few are slow because the
backend application is slow to answer those. This scenario is diagnosed through
the **distributed traces** Traefik ships to Tempo (otel-lgtm), not the access
log — `search_traces` finds the slow trace and `get_trace` shows the span tree.

## Run

```bash
./demo.sh up 08-trace-latency      # brings up the stack and runs the traffic

./demo.sh traffic 08-trace-latency # replay the traffic any time
```

`up` sends 15 fast and 5 slow (`?wait=2s`) requests once Traefik is ready, so
Tempo holds a mix of fast and slow traces. Allow a few seconds for ingestion
before querying.

## Demo in Claude Desktop

- "Some checkout requests are slow — find a slow one in the traces and tell me
  where the time is going."

Expected: the assistant calls `search_traces` with `{duration>1s}` (or
`{resource.service.name="traefik" && duration>1s}`), picks a slow trace, then
`get_trace` on its ID. It sees the entrypoint span lasting ~2s with the
reverse-proxy/service span consuming almost all of it, while Traefik's own
processing is negligible — so the latency is in the backend application, not the
proxy. Compare against a fast trace to make the point.

> Point the MCP server at Tempo so the assistant can query traces: add
> `--tempo.url=http://localhost:3200` to its args in the Claude Desktop config.

## Reset

```bash
./demo.sh down 08-trace-latency
```
