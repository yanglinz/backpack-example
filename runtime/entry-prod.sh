#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# Load environment variables
. "$(dirname "$0")/env-loader.sh"
. "$(dirname "$0")/berglas-loader.sh"

# Start production server
mkdir -p /app/var
/usr/bin/supervisord
