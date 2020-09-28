#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

instance_ip=$(cat /tmp/terraform-output.json | jq -r '.ip_address.value')
local_branch=$(git rev-parse --abbrev-ref HEAD)

# Setup ansible hosts inventory
mkdir -p ~/.ssh
ssh-keyscan -H "$instance_ip" >> ~/.ssh/known_hosts

# Setup SSH key pair
mkdir -p ~/.ssh
echo "$DIGITALOCEAN_PRIVATE_KEY" > ~/.ssh/id_rsa
sudo chmod 600 ~/.ssh/id_rsa
ssh-keygen -y -f ~/.ssh/id_rsa > ~/.ssh/id_rsa.pub
sudo chmod 600 ~/.ssh/id_rsa.pub

# Setup git remote and push remote
git remote add dokku dokku@"$instance_ip":"$APP_NAME"
GIT_SSH_COMMAND="ssh -l dokku" git push dokku "$local_branch":master -f
