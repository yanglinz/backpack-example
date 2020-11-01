#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

mkdir -p "${GITHUB_WORKSPACE}/bin"
curl -o "${GITHUB_WORKSPACE}/bin/berglas" https://storage.googleapis.com/berglas/main/linux_amd64/berglas
chmod +x "${GITHUB_WORKSPACE}/bin/berglas"

echo "::add-path::${GITHUB_WORKSPACE}/bin"
