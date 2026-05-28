#!/usr/bin/env bash
# Launch the demo stack, optionally with a scenario overlay.
#
#   ./demo.sh up                     # base stack only
#   ./demo.sh up 02-service-5xx      # base + scenario, then its traffic
#   ./demo.sh traffic 02-service-5xx # re-run a scenario's traffic only
#   ./demo.sh down [scenario]        # tear everything down
#   ./demo.sh logs
#   ./demo.sh ps
#   ./demo.sh list                   # list scenarios
#
# Switching scenarios is always a clean slate:
#   - up/down pass --remove-orphans, so a service from a previous scenario
#     (e.g. billing) is removed when you bring up a different one;
#   - dynamic config files a scenario ships in scenarios/<name>/dynamic/ are
#     copied into the watched config/dynamic dir on up and cleared on the next
#     up/down (copying rather than bind-mounting, since the dynamic dir is a
#     read-only mount and nested mountpoints can't be created inside it).
#
# A scenario may also ship scenarios/<name>/{compose.yml,traffic.sh}: the compose
# overlay adds services, and traffic.sh is run after the stack is ready so the
# steps to reproduce live with the scenario rather than only in its README.
set -euo pipefail

cd "$(dirname "$0")"

api_url="http://localhost:8088"
dyn_dir="config/dynamic"

cmd="${1:-}"
scenario="${2:-}"

files=(-f docker-compose.yml)
if [[ -n "$scenario" ]]; then
  if [[ ! -d "scenarios/$scenario" ]]; then
    echo "unknown scenario: $scenario" >&2
    echo "available:" >&2
    ls -1 scenarios >&2
    exit 1
  fi
  if [[ -f "scenarios/$scenario/compose.yml" ]]; then
    files+=(-f "scenarios/$scenario/compose.yml")
  fi
fi

# clear_dynamic removes scenario-provided dynamic config, keeping .gitkeep, so
# the watched dir starts empty for every scenario.
clear_dynamic() {
  find "$dyn_dir" -type f ! -name '.gitkeep' -delete
}

# sync_dynamic copies the current scenario's dynamic config into the watched dir.
sync_dynamic() {
  if [[ -n "$scenario" && -d "scenarios/$scenario/dynamic" ]]; then
    find "scenarios/$scenario/dynamic" -maxdepth 1 -type f -exec cp {} "$dyn_dir/" \;
  fi
}

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
    clear_dynamic
    sync_dynamic
    docker compose "${files[@]}" up -d --remove-orphans
    if [[ -n "$scenario" && -f "scenarios/$scenario/traffic.sh" ]]; then
      wait_ready
      run_traffic
    fi
    ;;
  traffic) wait_ready && run_traffic ;;
  down)
    docker compose "${files[@]}" down --remove-orphans
    clear_dynamic
    ;;
  logs) docker compose "${files[@]}" logs -f ;;
  ps)   docker compose "${files[@]}" ps ;;
  list) ls -1 scenarios ;;
  *)
    echo "usage: ./demo.sh {up|down|traffic|logs|ps|list} [scenario]" >&2
    exit 1
    ;;
esac
