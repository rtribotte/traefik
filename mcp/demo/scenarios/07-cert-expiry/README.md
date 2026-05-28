# Scenario 07 — "are any of my TLS certificates about to expire?"

A short-lived certificate is loaded into Traefik. Its expiry isn't in the access
log or the application log — it's exposed as a Prometheus metric,
`traefik_tls_certs_not_after` (a Unix timestamp per certificate). This is the
scenario solved with the metrics tool.

## Run

```bash
# generate a cert that expires in 2 days (writes certs/demo.crt + demo.key)
./scripts/gen-certs.sh 2 api.localhost

./demo.sh up 07-cert-expiry
```

## What to look at

Traefik exposes one `traefik_tls_certs_not_after` series per loaded certificate,
labelled with its `cn` and `sans`, whose value is the expiry as a Unix
timestamp:

```
traefik_tls_certs_not_after{cn="api.localhost",sans="api.localhost",serial="..."} 1.780153638e+09
```

## Demo in Claude Desktop

- "Are any of my TLS certificates expiring soon?"

Expected: the assistant calls `get_metrics`, finds the
`traefik_tls_certs_not_after` series, converts the Unix timestamp(s) to a date,
compares against now, and reports which certificates expire and when — flagging
the one only ~2 days out. The access and application logs don't carry this; the
answer comes from the metrics.

## Reset

```bash
./demo.sh down 07-cert-expiry
```
