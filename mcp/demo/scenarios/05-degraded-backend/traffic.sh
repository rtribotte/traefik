#!/usr/bin/env bash
# Send a burst to api.localhost. The service load-balances across two servers,
# one healthy and one pointing at a dead port, with no health check to evict it,
# so roughly half the requests come back 502 and half 200 — intermittently, with
# no router or service in an error state. The point is that the failure rate and
# its shape are a metrics question (query_metrics over the request totals by
# code), not something a single access-log snapshot answers.
set -euo pipefail

web="http://localhost:8081" # base stack publishes the web entrypoint here.

echo "sending 40 requests to api.localhost (about half will 502) ..."
ok=0; bad=0
for _ in $(seq 1 40); do
  code=$(curl -s -o /dev/null -w '%{http_code}' -H 'Host: api.localhost' "$web/")
  if [[ "$code" == "200" ]]; then ok=$((ok + 1)); else bad=$((bad + 1)); fi
done
echo "done; $ok ok / $bad failed. The rate and trend are in the metrics (query_metrics), not a log tail."
