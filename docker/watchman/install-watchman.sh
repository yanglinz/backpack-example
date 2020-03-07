#!/usr/bin/env bash
set -eo pipefail
IFS=$'\n\t'

if [ -z "$DISABLE_WATCHMAN" ]; then
  # Compiling watchman from source is quite slow, so in CI, we pass the
  # environmental variable DISABLE_WATCHMAN to avoid this step.
  git clone https://github.com/facebook/watchman.git
  cd watchman && git checkout v4.9.0 && ./autogen.sh && ./configure && make && make install
  rm -rf watchman
else
  echo "Compiling watchman is disabled"
fi
