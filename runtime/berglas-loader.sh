#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

function load_vars() {
  berglas access "$BERGLAS_SECRET_PATH" > /tmp/berglas-app.json
  eval $(python3 "$(dirname "$0")/load_env.py" /tmp/berglas-app.json)
}

if [[ -n "${BERGLAS_SECRET_PATH+set}" ]]; then
  load_vars
fi
