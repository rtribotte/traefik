#!/usr/bin/env bash
# Hit the broken route a few times. The router never registered (its v2-style
# rule fails to parse in v3), so every request 404s. The point is that the
# access log won't explain why — the router is simply absent — while the
# application log carries the rule parse error.
set -euo pipefail

web="http://localhost:8081" # base stack publishes the web entrypoint here.

echo "sending 5 requests to blog.localhost (they will 404) ..."
for _ in $(seq 1 5); do
  curl -s -o /dev/null -H 'Host: blog.localhost' "$web/api" || true
done
echo "done; the router never registered — check the application log (tail_traefik_logs) for the rule parse error"
