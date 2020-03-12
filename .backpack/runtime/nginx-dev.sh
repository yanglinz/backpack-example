#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

"$(dirname "$0")/../docker/scripts/wait-for-it.sh" 0.0.0.0:4567 -t 10

nginx -c /etc/nginx/nginx.conf
