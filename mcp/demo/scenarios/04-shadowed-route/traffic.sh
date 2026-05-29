#!/usr/bin/env bash
# Hit shop.localhost a few times. The shop router is enabled and its backend is
# healthy, yet every response comes back from the maintenance backend, because
# the catch-all router (priority 1000) outranks the shop router. The body shows
# "Name: maintenance" instead of the shop backend, and the access log attributes
# the requests to catchall@file / maintenance@file — but nothing is in an error
# state, so the cause is the priority, not a broken router.
set -euo pipefail

web="http://localhost:8081" # base stack publishes the web entrypoint here.

echo "sending 5 requests to shop.localhost (served by the wrong backend) ..."
for _ in $(seq 1 5); do
  curl -s -H 'Host: shop.localhost' "$web/" | grep -i '^Name:' || true
done
echo "done; shop.localhost is answered by the maintenance backend — compare the two routers' priorities"
