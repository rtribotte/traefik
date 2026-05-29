# Scenario 02 — "is anything exposed?"

The full hardening loop: **audit → harden → validate**. Surface the recurring,
high-stakes findings from the live runtime, then use the embedded reference to
write the fix and validate it before applying. There is no built-in Traefik
audit today, so this composes what the server already has into one answer.

## Run

```bash
# generate a cert that expires in 7 days (so its expiry metric stands out)
./scripts/gen-certs.sh

./demo.sh up 02-security-posture
```

## What's flagged

| Finding | Source the assistant reads |
|---------|----------------------------|
| `admin.localhost` exposed on plain HTTP — no TLS, no auth middleware | `/api/rawdata` / `list_routers` (public route on the `web` entrypoint, empty `tls`, no `middlewares`) |
| TLS certificate expiring in <30 days | `get_metrics` → `traefik_tls_certs_not_after` (Unix timestamp per cert) |

## Demo in Claude Desktop

- "Audit my Traefik setup for security issues and show me how to fix them."

Expected flow:

1. **Audit (live).** The assistant scans the runtime for public routes on a plain
   HTTP entrypoint with no TLS and no auth (`admin`, and the base `whoami`), then
   reads `traefik_tls_certs_not_after` from the metrics, converts the timestamp,
   and flags the certificate expiring soonest.
2. **Harden (reference).** `search_traefik_docs` for the fixes →
   `get_traefik_concept("http.middlewares.redirectscheme")` to force HTTPS,
   `get_traefik_concept("http.middlewares.basicauth")` to add authentication, and
   `get_traefik_concept("http.middlewares.headers")` for secure response headers.
3. **Author + validate.** The assistant writes a hardened dynamic config (TLS on
   the router + the auth/redirect/headers middlewares) and runs
   `validate_traefik_config` on it to prove it is well-formed against the official
   schema before you apply it — the server never writes config itself.

> Point the MCP server at the metrics endpoint (the demo's `--prometheus.url` /
> `get_metrics` already cover this) so the assistant can read the cert expiry.

## Reset

```bash
./demo.sh down 02-security-posture
```
