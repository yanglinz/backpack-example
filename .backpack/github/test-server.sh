#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

function check_requirements() {
  if cmp -s ./requirements.txt <(poetry export --format requirements.txt); then
    echo "requirements.txt and pyproject.toml are in sync"
  else
    echo "requirements.txt and pyproject.toml are out of sync"
    exit 1
  fi
}

check_requirements
