#!/usr/bin/env bash
# Hit the broken route a few times. The requests fail (the router can't resolve
# its service/middleware), which is the point: the access log won't explain why,
# only the application log will.
set -euo pipefail

web="http://localhost:8081" # base stack publishes the web entrypoint here.

echo "sending 5 requests to reports.localhost (they will fail) ..."
for _ in $(seq 1 5); do
  curl -s -o /dev/null -H 'Host: reports.localhost' "$web/" || true
done
echo "done; check the application log (tail_traefik_logs) for the unresolved references"
