#!/usr/bin/env bash
# Generate a mix of fast and slow traffic for the trace-latency scenario so the
# traces in Tempo contain both, and a {duration>1s} search isolates the slow
# ones. whoami delays its reply when asked with ?wait.
set -euo pipefail

web="http://localhost:8081" # base stack publishes the web entrypoint here.

echo "sending 15 fast + 5 slow requests to checkout.localhost ..."
for _ in $(seq 1 15); do
  curl -s -o /dev/null -H 'Host: checkout.localhost' "$web/"
done
for _ in $(seq 1 5); do
  curl -s -o /dev/null -H 'Host: checkout.localhost' "$web/?wait=2s"
done
echo "done; traces are in Tempo. Allow a few seconds for ingestion."
