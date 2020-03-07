#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

git ls-files -s | shasum | awk '{print $1}' | head -c 12
