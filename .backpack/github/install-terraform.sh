#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

TF_VERSION="0.13.3"

git clone https://github.com/tfutils/tfenv.git ~/.tfenv
~/.tfenv/bin/tfenv install "$TF_VERSION"
~/.tfenv/bin/tfenv use "$TF_VERSION"
echo "::add-path::${HOME}/.tfenv/bin"
