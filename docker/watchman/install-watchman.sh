#!/usr/bin/env bash
set -eo pipefail
IFS=$'\n\t'

if [ -z "$DISABLE_WATCHMAN" ]; then
  # Compiling watchman from source is quite slow, so in CI, we pass the
  # environmental variable DISABLE_WATCHMAN to avoid this step.
  WATCHMAN_VERSION="v2020.09.21.00"
  curl -L "https://github.com/facebook/watchman/releases/download/${WATCHMAN_VERSION}/watchman-${WATCHMAN_VERSION}-linux.zip" > /tmp/watchman.zip
  unzip -o /tmp/watchman.zip -d /tmp/
  mkdir -p /usr/local/{bin,lib} /usr/local/var/run/watchman
  cp "/tmp/watchman-${WATCHMAN_VERSION}-linux"/bin/* /usr/local/bin
  cp "/tmp/watchman-${WATCHMAN_VERSION}-linux"/lib/* /usr/local/lib
  chmod 755 /usr/local/bin/watchman
  chmod 2777 /usr/local/var/run/watchman
else
  echo "Compiling watchman is disabled"
fi
