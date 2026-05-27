# Scenario 04 — add a secured route (advisory)

No extra infrastructure. The MCP server is read-only, so this scenario is about
*generating* a correct config, not applying it. The base stack is enough.

## Run

```bash
./demo.sh up
```

## Demo in Claude Desktop

- "Generate a router exposing api.example.com to my billing service, with HTTPS
  and basic auth."

Expected: the assistant produces a valid dynamic-configuration snippet (router +
TLS + a basicauth middleware), explains where to put it, and reminds you to apply
it yourself (the server never writes config). Once config validation lands, the
generated snippet can be checked automatically before you paste it.

## Reset

```bash
./demo.sh down
```
