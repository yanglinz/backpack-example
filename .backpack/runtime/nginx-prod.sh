#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# Generate nginx conf
# https://unix.stackexchange.com/questions/294378/replacing-only-specific-variables-with-envsubst/294400
envsubst '${PORT}' < /app/.backpack/docker/nginx/nginx-prod.tmpl.conf > /etc/nginx/nginx.conf

"$(dirname "$0")/../docker/scripts/wait-for-it.sh" 0.0.0.0:4567 -t 5

nginx -c /etc/nginx/nginx.conf
