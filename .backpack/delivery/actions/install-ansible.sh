#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# Install ansible
sudo apt-add-repository ppa:ansible/ansible
sudo apt update && sudo apt install ansible
