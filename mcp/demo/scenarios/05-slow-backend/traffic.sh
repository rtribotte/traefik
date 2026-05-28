#!/usr/bin/env bash
# Generate slow traffic for the slow-backend scenario. whoami has no "always
# slow" mode; it delays its reply when asked with ?wait, so each request takes
# ~2s and the access log records the high duration.
set -euo pipefail

web="http://localhost:8081" # base stack publishes the web entrypoint here.

echo "sending 20 slow requests to catalog.localhost (~2s each) ..."
for _ in $(seq 1 20); do
  curl -s -o /dev/null -H 'Host: catalog.localhost' "$web/?wait=2s"
done
echo "done; check the access log for ~2000ms durations on catalog-demo@docker"
