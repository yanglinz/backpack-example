#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

function check_setup() {
  ./backpack setup

  git status
  if [ -n "$(git status --porcelain)" ]; then
    echo "Please run ./backpack setup"
    echo "Commit the auto-generated files"
    exit 1;
  fi
}

function run_tests() {
  cd .backpack

  # Run tests
  make test
  
  # Run formatting
  make format
  git status
  if [ -n "$(git status --porcelain)" ]; then
    echo "Please run go fmt on the source code"
    exit 1;
  fi

  cd -
}

check_setup
run_tests
