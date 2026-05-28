# Scenario 06 — "reports.localhost doesn't work and I don't know why"

A dynamic config file is loaded, but its router references a service
(`reports-backend`) and a middleware (`reports-auth`) that are never defined.
Traefik logs an error for each unresolved reference and the route never serves
traffic. This is the scenario that can only be diagnosed from the **application
log** — the access log has nothing useful (no request ever succeeds) and the
router list alone doesn't make the cause obvious.

## Run

```bash
./demo.sh up 06-config-error   # brings up the stack and sends the failing traffic

./demo.sh traffic 06-config-error   # replay the failing traffic any time
```

## What's wrong

`scenarios/06-config-error/dynamic/broken.yml` defines router `reports@file`
pointing at service `reports-backend` and middleware `reports-auth`, neither of
which exists. Traefik loads the file but cannot apply the router, logging:

- `the service "reports-backend@file" does not exist`
- `middleware "reports-auth@file" does not exist`

## Demo in Claude Desktop

- "reports.localhost isn't working — can you find out why?"

Expected: the assistant checks the access log and finds nothing explanatory,
then reads the application log (`tail_traefik_logs`, filtering to errors) and
surfaces the unresolved service and middleware references. It concludes the
dynamic config references names that don't exist, and the fix is to define them
(or correct the references).

> Point the MCP server at the application log so the assistant can read it:
> add `--traefik.app-log=/Users/romain/go/src/github.com/traefik/traefik/mcp/demo/logs/traefik.log`
> to its args in the Claude Desktop config.

## Reset

```bash
./demo.sh down 06-config-error
```
