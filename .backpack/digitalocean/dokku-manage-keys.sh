#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

dokku ssh-keys:add backpack ~/.ssh/authorized_keys || true
