# Scenario 02 — "I'm getting 502s"

A backend dies but its router and service stay configured. Classic incident.

## Run

```bash
./demo.sh up 02-service-5xx

# confirm it works
curl -H 'Host: billing.localhost' http://localhost/

# break it: kill the backend, leave routing intact
docker compose -f docker-compose.yml -f scenarios/02-service-5xx/compose.yml stop billing

# now it 502s
curl -H 'Host: billing.localhost' http://localhost/
```

## Demo in Claude Desktop

- "billing.localhost is returning 502, what's going on?"

Expected: the assistant checks `get_service_health` for `billing@docker` (servers
DOWN), confirms the router/service still exist via `get_service`, scans recent
access logs for the 502s, and concludes the backend is down rather than a routing
or TLS problem.

## Reset

```bash
./demo.sh down 02-service-5xx
```
