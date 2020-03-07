#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# Load environment variables
. "$(dirname "$0")/env-loader.sh"

# Start a pipenv shell
pipenv shell
