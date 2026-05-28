#!/usr/bin/env bash
# Generate the 502s for the misconfigured-backend scenario. The billing service
# points at port 9999 while whoami listens on 80, so Traefik returns 502 Bad
# Gateway. These requests put those 502s in the access log.
set -euo pipefail

web="http://localhost:8081" # base stack publishes the web entrypoint here.

echo "sending 10 requests to billing.localhost (expect 502) ..."
for _ in $(seq 1 10); do
  curl -s -o /dev/null -H 'Host: billing.localhost' "$web/"
done
echo "done; check the access log for 502s on billing@docker"
