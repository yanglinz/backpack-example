#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# Load environment variables
. "$(dirname "$0")/env-loader.sh"
. "$(dirname "$0")/berglas-loader.sh"

# Generate nginx conf
# https://unix.stackexchange.com/questions/294378/replacing-only-specific-variables-with-envsubst/294400
envsubst '${PORT}' < /app/.backpack/docker/nginx/nginx-prod.tmpl.conf > /etc/nginx/nginx.conf

# Start production server
mkdir -p /app/var
/usr/bin/supervisord
