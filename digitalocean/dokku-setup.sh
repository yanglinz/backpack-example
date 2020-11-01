#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

dokku ssh-keys:add backpack ~/.ssh/authorized_keys || true
dokku ssh-keys:add "$APP_NAME" ~/.ssh/authorized_keys || true
dokku git:initialize "$APP_NAME"
