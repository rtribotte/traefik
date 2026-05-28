#!/usr/bin/env bash
# Generate traffic for the trace-latency scenario. Every request is delayed a
# random amount by the randomplugin middleware, so the traces in Tempo show a
# spread of durations and a {duration>1s} search isolates the slow ones.
set -euo pipefail

web="http://localhost:8081" # base stack publishes the web entrypoint here.

echo "sending 20 requests to checkout.localhost (each delayed a random amount) ..."
for _ in $(seq 1 20); do
  curl -s -o /dev/null -H 'Host: checkout.localhost' "$web/"
done
echo "done; traces are in Tempo. Allow a few seconds for ingestion."
