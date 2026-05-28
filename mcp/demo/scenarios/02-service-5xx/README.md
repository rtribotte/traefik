# Scenario 02 — "I'm getting 502s — is it Traefik or my service?"

A backend that's misconfigured, not down. The router and service stay in the
config, so this is a real 502 (not a 404 from a vanished route). The fault is in
the service definition — the wrong backend port — not in Traefik.

## Run

```bash
./demo.sh up 02-service-5xx   # brings up the stack and sends the 502 traffic

./demo.sh traffic 02-service-5xx   # replay the 502 traffic any time
```

`up` runs this scenario's `traffic.sh` automatically once Traefik is ready: it
sends requests to `billing.localhost`, which 502 because Traefik routes
correctly but can't reach the backend on port 9999. The base stack publishes the
web entrypoint on host port 8081.

## What's wrong

`billing` listens on port 80 (whoami's default), but its service is configured
with `loadbalancer.server.port=9999`. Traefik resolves the route, picks the
service, dials port 9999, gets refused, and returns 502.

## Demo in Claude Desktop

- "billing.localhost returns 502, is the problem Traefik or my service?"

Expected: the assistant confirms the router and service exist and are enabled
(routing is fine), tails the access log to see the 502s on `billing@docker`,
notes the backend is unreachable, and concludes the service is misconfigured —
the configured backend port doesn't match where the app listens. The fix is in
the service config, not Traefik.

> Point the MCP server at the access log so the assistant can see the 502s:
> add `--traefik.access-log=/Users/romain/go/src/github.com/traefik/traefik/mcp/demo/logs/access.log`
> to its args in the Claude Desktop config.

## Reset

```bash
./demo.sh down 02-service-5xx
```
