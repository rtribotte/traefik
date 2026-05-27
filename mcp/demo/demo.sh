#!/usr/bin/env bash
# Launch the demo stack, optionally with a scenario overlay.
#
#   ./demo.sh up                     # base stack only
#   ./demo.sh up 01-router-missing   # base + scenario overlay
#   ./demo.sh down 01-router-missing
#   ./demo.sh logs
#   ./demo.sh list                   # list scenarios
set -euo pipefail

cd "$(dirname "$0")"

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

case "$cmd" in
  up)   docker compose "${files[@]}" up -d ;;
  down) docker compose "${files[@]}" down ;;
  logs) docker compose "${files[@]}" logs -f ;;
  ps)   docker compose "${files[@]}" ps ;;
  list) ls -1 scenarios ;;
  *)
    echo "usage: ./demo.sh {up|down|logs|ps|list} [scenario]" >&2
    exit 1
    ;;
esac
