# Scenario 04 — "shop.localhost serves the wrong page, but nothing is broken"

A valid-but-wrong configuration: the `shop` router is enabled, its rule is
correct and its backend is healthy, yet `shop.localhost` answers from the
*maintenance* backend. The cause is **router priority** — a second router for the
same host, with an explicit `priority: 1000`, outranks the intended `shop`
router — and it shows up nowhere as an error. This is the scenario that the live
read tools alone diagnose only halfway: they surface the *what*, but the *fix*
needs the reference (how v3 derives and compares priority) and
`validate_traefik_config`.

## Run

```bash
./demo.sh up 04-shadowed-route      # brings up the stack and sends the traffic

./demo.sh traffic 04-shadowed-route # replay the traffic any time
```

## What's wrong

`scenarios/04-shadowed-route/dynamic/shadow.yml` defines two routers that both
match `Host(`shop.localhost`)` on the `web` entrypoint, both `enabled`:

- `shop@file` — no explicit priority, so Traefik derives it from the rule length
  (**22**).
- `shop-maintenance@file` — same rule, but with `priority: 1000` (a leftover from
  pointing the host at a maintenance page during an outage).

A higher priority wins, so `shop-maintenance` outranks `shop` and `shop.localhost`
is served by the `maintenance` backend. No router is disabled, no rule fails to
parse, the backend is up — the only evidence is the priority gap.

## Demo in Claude Desktop

- "shop.localhost returns the maintenance page instead of the shop page, but the shop
  router looks fine, what is happening and how can I fix it ?"

Expected flow:

1. **Diagnose (live).** `tail_access_logs` (or `query_access_logs`) for
   `shop.localhost` shows the requests attributed to `shop-maintenance@file` /
   `maintenance@file`, not `shop@file` — a *different* router is winning.
   `get_router("shop@file")` and `get_router("shop-maintenance@file")` both show
   `status: enabled`; the tell is `shop-maintenance`'s `priority: 1000` against
   `shop`'s `22`. The backend is healthy, so this is a routing-precedence problem,
   not a broken router or service.
2. **Look up (reference).** `search_traefik_docs("router priority")` →
   `get_traefik_concept("http.routers")`: when several routers match, the one with
   the highest priority is selected, and an unset priority defaults to the length
   of the rule — so a hand-set `priority: 1000` will always beat a router that
   relies on the default.
3. **Fix.** Remove the leftover `shop-maintenance` router, or — if it is still
   needed — scope its rule to a maintenance path and drop the manual priority so
   it no longer shadows `shop`.
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
