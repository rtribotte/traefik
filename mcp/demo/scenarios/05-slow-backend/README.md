# Scenario 05 — "catalog.localhost is slow — what's going on?"

Nothing is broken. The router and service exist, the backend is UP, every
request returns 200. The backend application is simply slow to respond. This is
the one scenario that cannot be diagnosed from the configuration or health alone
— the only evidence is the request duration in the access log.

## Run

```bash
./demo.sh up 05-slow-backend   # brings up the stack and runs the slow traffic

./demo.sh traffic 05-slow-backend   # replay the slow traffic any time
```

`up` runs this scenario's `traffic.sh` automatically once Traefik is ready: it
sends 20 requests with `?wait=2s` (whoami delays its reply when asked), so each
takes ~2s and the access log records the high duration. The base stack publishes
the web entrypoint on host port 8081.

## What's "wrong"

Nothing in Traefik or the service config. The application takes ~2s to answer.
Traefik routes correctly, the backend is healthy, the status is 200 — but the
access log shows each request's `durationMs` around 2000.

## Demo in Claude Desktop

- "catalog.localhost feels slow, can you find out why?"

Expected: the assistant confirms the router and service exist and are enabled and
the backend servers are UP (so routing and health are fine), then tails the
access log filtered by host `catalog.localhost` — ideally with `minDurationMs` to
surface slow requests — and sees the requests succeeding (200) but taking ~2s.
It concludes the latency is in the backend application, not Traefik: the proxy
overhead is negligible and the time is spent waiting on the upstream. The fix is
in the application, not the routing.

> Point the MCP server at the access log so the assistant can see the durations:
> add `--traefik.access-log=/Users/romain/go/src/github.com/traefik/traefik/mcp/demo/logs/access.log`
> to its args in the Claude Desktop config.

## Reset

```bash
./demo.sh down 05-slow-backend
```
