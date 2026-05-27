# Scenario 03 — security audit

Surface two recurring, high-stakes findings.

## Run

```bash
# generate a cert that expires in 7 days (so cert_expiry flags it)
./scripts/gen-certs.sh

./demo.sh up 03-security-audit
```

## What's flagged

| Finding | Source |
|---------|--------|
| `admin.localhost` exposed on plain HTTP, no TLS, no auth middleware | `admin` service labels |
| TLS certificate expiring in <30 days | short-lived cert mounted via `dynamic/tls.yml` |

## Demo in Claude Desktop

- "Audit my Traefik setup for security issues."
- "Which certificates expire in the next 30 days?"

Expected: the assistant scans the runtime config for public routes without TLS or
auth (`admin`, and the base `whoami`), lists certificates and ranks the expiring
one, and explains the fix (add TLS + an auth middleware, rotate the cert).

## Reset

```bash
./demo.sh down 03-security-audit
```
