#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

if [[ -f "/app/scripts/docker/install-extra-deps.sh" ]]; then
  /app/scripts/docker/install-extra-deps.sh
else
  echo "No extra dependencies to install"
fi
