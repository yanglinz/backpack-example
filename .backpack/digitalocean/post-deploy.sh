#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

ENV_SOURCE="./var/env/production.json"

echo "Setting ${APP_NAME} application configs..."

for name in $(jq --raw-output 'keys | .[]' "$ENV_SOURCE"); do
  value=$(jq --raw-output ".${name}" "$ENV_SOURCE")
  echo "$name - $value"
done

# Set a dummy variable to reload
