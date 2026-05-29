# Scenario 01 — "shop.localhost is broken and I don't know why"

The full broken-route loop: **diagnose → look up → fix → validate**. A router
silently fails to register because its rule still uses Traefik v2 syntax. This is
the most common Traefik question, and it exercises every layer of the server at
once — live introspection, the embedded reference, and config validation.

## Run

```bash
./demo.sh up 01-broken-route      # brings up the stack and sends the failing traffic

./demo.sh traffic 01-broken-route # replay the failing traffic any time
```

## What's wrong

`scenarios/01-broken-route/dynamic/broken.yml` defines router `shop@file` with the
rule:

```
Host(`shop.localhost`) && PathPrefix(`/api`, `/admin`)
```

`PathPrefix` with several arguments is the **Traefik v2** form. In v3 every matcher
takes a single argument, so the rule fails to parse, the router is never added to
the runtime, and `shop.localhost` 404s. Traefik logs the parse error; the `shop`
*service* is valid and shows up unused in `list_services`.

## Demo in Claude Desktop

- "shop.localhost returns 404 but I configured a router for it — what's wrong, and
  how do I fix it?"

Expected flow:

1. **Diagnose (live).** `get_router("shop@file")` shows `status: disabled` with the
   rule parse error attached (`PathPrefix: unexpected number of parameters; got 2,
   expected one of [1]`); `list_services` shows `shop@file` defined and healthy.
   `tail_traefik_logs` filtered to errors carries the same message. The router, not
   the backend, is broken.
2. **Look up (reference).** `search_traefik_docs("routing rule syntax")` →
   `get_traefik_concept("http.routers")` for the v3 rule grammar: matchers take one
   argument; the v2 multi-argument `PathPrefix` becomes
   `PathPrefix(`/api`) || PathPrefix(`/admin`)`.
3. **Fix.** The assistant rewrites the rule to the v3 form.
4. **Validate (before applying).** `validate_traefik_config` on the corrected dynamic
   config proves it is well-formed against the official schema before you paste it
   back — the server is read-only, so it never writes the file itself.

> Point the MCP server at the application log so the assistant can read the parse
> error: add
> `--traefik.app-log=/Users/romain/go/src/github.com/traefik/traefik/mcp/demo/logs/traefik.log`
> to its args in the Claude Desktop config.

## Reset

```bash
./demo.sh down 01-broken-route
```
