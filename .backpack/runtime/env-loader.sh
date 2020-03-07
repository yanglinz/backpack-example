#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

function load_vars() {
  eval $(python3 "$(dirname "$0")/load_env.py" /app/etc/development.json)
}

if [[ -n "${BACKPACK_DOCKER_COMPOSE+set}" ]]; then
  load_vars
fi

