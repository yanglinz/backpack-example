#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

ENV_SOURCE="/opt/backpack-app/var/env/production.json"

echo "Setting ${APP_NAME} application configs..."

for name in $(jq --raw-output 'keys | .[]' "$ENV_SOURCE"); do
  value=$(jq --raw-output ".${name}" "$ENV_SOURCE")
  dokku config:set --no-restart "$APP_NAME" "$name"="$value"
done

dokku config:set "$APP_NAME" TIMESTAMP=$(date +%s)
dokku domains:enable backpack-example
