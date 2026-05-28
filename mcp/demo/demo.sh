#!/usr/bin/env bash
# Launch the demo stack, optionally with a scenario overlay.
#
#   ./demo.sh up                     # base stack only
#   ./demo.sh up 02-service-5xx      # base + scenario overlay, then its traffic
#   ./demo.sh traffic 02-service-5xx # re-run a scenario's traffic only
#   ./demo.sh down [scenario]        # tear everything down
#   ./demo.sh logs
#   ./demo.sh ps
#   ./demo.sh list                   # list scenarios
#
# Switching scenarios is always a clean slate: up/down pass --remove-orphans, so
# a service from a previous scenario (e.g. billing) is removed when you bring up
# a different one. A scenario may ship a scenarios/<name>/traffic.sh that demo.sh
# runs after the stack is up, so the steps to reproduce live with the scenario
# rather than only in its README.
set -euo pipefail

cd "$(dirname "$0")"

api_url="http://localhost:8088"

cmd="${1:-}"
scenario="${2:-}"

files=(-f docker-compose.yml)
if [[ -n "$scenario" ]]; then
  overlay="scenarios/$scenario/compose.yml"
  if [[ ! -f "$overlay" ]]; then
    echo "unknown scenario: $scenario" >&2
    echo "available:" >&2
    ls -1 scenarios >&2
    exit 1
  fi
  files+=(-f "$overlay")
fi

# wait_ready blocks until the Traefik API answers, so scenario traffic does not
# fire against a proxy that has not loaded its configuration yet.
wait_ready() {
  echo "waiting for Traefik API at $api_url ..."
  for _ in $(seq 1 30); do
    if curl -fsS -o /dev/null "$api_url/api/overview"; then
      return 0
    fi
    sleep 1
  done
  echo "Traefik API did not become ready at $api_url" >&2
  return 1
}

# run_traffic executes the scenario's traffic.sh, if it has one.
run_traffic() {
  if [[ -z "$scenario" ]]; then
    echo "no scenario given; nothing to run" >&2
    return 0
  fi
  local script="scenarios/$scenario/traffic.sh"
  if [[ -f "$script" ]]; then
    echo "running traffic for $scenario ..."
    API_URL="$api_url" bash "$script"
  fi
}

case "$cmd" in
  up)
    docker compose "${files[@]}" up -d --remove-orphans
    if [[ -n "$scenario" && -f "scenarios/$scenario/traffic.sh" ]]; then
      wait_ready
      run_traffic
    fi
    ;;
  traffic) wait_ready && run_traffic ;;
  down) docker compose "${files[@]}" down --remove-orphans ;;
  logs) docker compose "${files[@]}" logs -f ;;
  ps)   docker compose "${files[@]}" ps ;;
  list) ls -1 scenarios ;;
  *)
    echo "usage: ./demo.sh {up|down|traffic|logs|ps|list} [scenario]" >&2
    exit 1
    ;;
esac
