# Scenario 04 — "shop.localhost serves the wrong page, but nothing is broken"

A valid-but-wrong configuration: the `shop` router is enabled, its rule is
correct and its backend is healthy, yet `shop.localhost` answers from the
*maintenance* backend. The cause is **router priority** — a catch-all router
with an explicit `priority: 1000` outranks the more specific `shop` router — and
it shows up nowhere as an error. This is the scenario that the live read tools
alone diagnose only halfway: they surface the *what*, but the *fix* needs the
reference (how v3 derives and compares priority) and `validate_traefik_config`.

## Run

```bash
./demo.sh up 04-shadowed-route      # brings up the stack and sends the traffic

./demo.sh traffic 04-shadowed-route # replay the traffic any time
```

## What's wrong

`scenarios/04-shadowed-route/dynamic/shadow.yml` defines two routers on the `web`
entrypoint, both `enabled`:

- `shop@file` — `Host(`shop.localhost`)`, no explicit priority, so Traefik
  derives it from the rule length (**22**).
- `catchall@file` — `PathPrefix(`/`)` with `priority: 1000`.

A higher priority wins, so `catchall` outranks `shop` (and every other web
router) and `shop.localhost` is served by the `maintenance` backend. No router is
disabled, no rule fails to parse, the backend is up — the only evidence is the
priority gap.

## Demo in Claude Desktop

- "shop.localhost returns the maintenance page instead of the shop, but the shop
  router looks fine and the backend is up. Why, and how do I fix it?"

Expected flow:

1. **Diagnose (live).** `tail_access_logs` (or `query_access_logs`) for
   `shop.localhost` shows the requests attributed to `catchall@file` /
   `maintenance@file`, not `shop@file` — a *different* router is winning.
   `get_router("shop@file")` and `get_router("catchall@file")` both show
   `status: enabled`; the tell is `catchall`'s `priority: 1000` against `shop`'s
   `22`. The backend is healthy, so this is a routing-precedence problem, not a
   broken router or service.
2. **Look up (reference).** `search_traefik_docs("router priority")` →
   `get_traefik_concept("http.routers")`: when several routers match, the one with
   the highest priority is selected, and an unset priority defaults to the length
   of the rule — so a hand-set `priority: 1000` on a catch-all will always beat a
   specific host router.
3. **Fix.** Remove the explicit priority from the catch-all (let rule length
   order them so the specific `shop` rule wins), or scope the catch-all so it no
   longer matches `shop.localhost`.
4. **Validate (before applying).** `validate_traefik_config` on the corrected
   dynamic config proves it is well-formed against the official schema before you
   paste it back — the server never writes the file itself.

> Point the MCP server at the access log so the assistant can see which router
> served the request: add
> `--traefik.access-log=/Users/romain/go/src/github.com/traefik/traefik/mcp/demo/logs/access.log`
> to its args in the Claude Desktop config.

## Reset

```bash
./demo.sh down 04-shadowed-route
```
