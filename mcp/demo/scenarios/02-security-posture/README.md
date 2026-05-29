# Scenario 02 — "is anything exposed, and is a cert about to expire?"

The full hardening loop: **audit → harden → validate**. Detection here is easy —
the findings are right there in the runtime — so the value is the second half:
grounding each fix in the embedded reference and proving the hardened config
well-formed with `validate_traefik_config` before you apply it. There is no
built-in Traefik audit today, so this composes what the server already has into
one answer.

## Run

```bash
# generate a cert that expires in 7 days, so list_certificates flags it
./scripts/gen-certs.sh

./demo.sh up 02-security-posture
```

## What's flagged

| Finding | Source the assistant reads |
|---------|----------------------------|
| `admin.localhost` exposed on plain HTTP — no TLS, no auth middleware | `/api/rawdata` / `list_routers` (public route on the `web` entrypoint, empty `tls`, no `middlewares`) |
| TLS certificate expiring in <30 days | `list_certificates` (`status: warning`, `daysUntilExpiry`, the SANs and issuer) |

`list_certificates` is the authoritative, live view: it reads Traefik's
certificates API and returns each cert's validity window, the `status` Traefik
itself computes (`warning` under 30 days, `expired` past `notAfter`) and a
`daysUntilExpiry` — richer than the `traefik_tls_certs_not_after` metric, which
carries only a CN and a Unix timestamp.

## Demo in Claude Desktop

- "Audit my Traefik setup for security issues and show me how to fix them."

Expected flow:

1. **Audit (live).** The assistant scans the runtime for public routes on a plain
   HTTP entrypoint with no TLS and no auth (`admin`, and the base `whoami`), then
   calls `list_certificates` and flags the one with `status: warning` (the
   short-lived demo cert), reporting its `daysUntilExpiry`.
2. **Harden (reference).** `search_traefik_docs` for each fix →
   `get_traefik_concept("http.middlewares.redirectscheme")` to force HTTPS,
   `get_traefik_concept("http.middlewares.basicauth")` to add authentication, and
   `get_traefik_concept("http.middlewares.headers")` for secure response headers.
   Grounding here is the point: the exact field names, types and defaults come
   from the reference, not from memory.
3. **Author + validate.** The assistant writes a hardened dynamic config (TLS on
   the router + the auth/redirect/headers middlewares) and runs
   `validate_traefik_config` on it to prove it is well-formed against the official
   schema before you apply it — the server never writes config itself.

> `list_certificates` needs only `--traefik.api-url` (already set); the cert API
> is part of the Traefik API the server reads.

## Reset

```bash
./demo.sh down 02-security-posture
```
