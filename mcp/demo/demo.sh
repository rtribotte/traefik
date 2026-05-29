#!/usr/bin/env bash
# Launch the demo stack, optionally with a scenario overlay.
#
#   ./demo.sh up                     # base stack only
#   ./demo.sh up 01-broken-route     # base + scenario, then its traffic
#   ./demo.sh up all                 # every scenario at once, then all traffic
#   ./demo.sh traffic 01-broken-route # re-run a scenario's traffic only
#   ./demo.sh down [scenario|all]    # tear everything down
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

# scenario_dirs is the set of scenario directories this invocation acts on:
# none for the base stack, one for a named scenario, or every scenario for the
# special "all" target (all scenarios up at once — their hosts and resource
# names are distinct, so they coexist on the one stack).
scenario_dirs=()
if [[ "$scenario" == "all" ]]; then
  for d in scenarios/*/; do
    scenario_dirs+=("${d%/}")
  done
elif [[ -n "$scenario" ]]; then
  if [[ ! -d "scenarios/$scenario" ]]; then
    echo "unknown scenario: $scenario" >&2
    echo "available (or 'all'):" >&2
    ls -1 scenarios >&2
    exit 1
  fi
  scenario_dirs+=("scenarios/$scenario")
fi

files=(-f docker-compose.yml)
for sd in "${scenario_dirs[@]}"; do
  if [[ -f "$sd/compose.yml" ]]; then
    files+=(-f "$sd/compose.yml")
  fi
done

# clear_dynamic removes scenario-provided dynamic config, keeping .gitkeep, so
# the watched dir starts empty for every scenario.
clear_dynamic() {
  find "$dyn_dir" -type f ! -name '.gitkeep' -delete
}

# sync_dynamic copies every selected scenario's dynamic config into the watched
# dir (one scenario, or all of them for the "all" target).
sync_dynamic() {
  for sd in "${scenario_dirs[@]}"; do
    if [[ -d "$sd/dynamic" ]]; then
      find "$sd/dynamic" -maxdepth 1 -type f -exec cp {} "$dyn_dir/" \;
    fi
  done
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

# run_traffic executes each selected scenario's traffic.sh, if it has one.
run_traffic() {
  if [[ ${#scenario_dirs[@]} -eq 0 ]]; then
    echo "no scenario given; nothing to run" >&2
    return 0
  fi
  for sd in "${scenario_dirs[@]}"; do
    if [[ -f "$sd/traffic.sh" ]]; then
      echo "running traffic for ${sd#scenarios/} ..."
      API_URL="$api_url" bash "$sd/traffic.sh"
    fi
  done
}

case "$cmd" in
  up)
    clear_dynamic
    sync_dynamic
    docker compose "${files[@]}" up -d --remove-orphans
    # Traefik's static config (traefik.yml: tracing, metrics, entrypoints) is
    # only read at startup. docker compose up won't recreate an already-running
    # traefik on a mounted-file change, so force-recreate just that service to
    # pick up edits when switching scenarios.
    docker compose "${files[@]}" up -d --force-recreate --no-deps traefik
    if [[ ${#scenario_dirs[@]} -gt 0 ]]; then
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
    echo "usage: ./demo.sh {up|down|traffic|logs|ps|list} [scenario|all]" >&2
    exit 1
    ;;
esac
