#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# Install ansible
sudo apt-add-repository ppa:ansible/ansible
sudo apt update && sudo apt install ansible

instance_ip=$(cat /tmp/terraform-output.json | jq '.ip_address.value')

# Setup ansible hosts inventory
mkdir -p ~/.ssh
mkdir -p ./etc/ansible
ssh-keyscan -H "$instance_ip" >> ~/.ssh/known_hosts
echo "$instance_ip" > ./etc/ansible/hosts

# Setup SSH key pair
mkdir -p ./etc/ssh
echo "$DIGITALOCEAN_PRIVATE_KEY" > ./etc/ssh/id_rsa
sudo chmod 600 ./etc/ssh/id_rsa
ssh-keygen -y -f ./etc/ssh/id_rsa > ./etc/ssh/id_rsa.pub
sudo chmod 600 ./etc/ssh/id_rsa.pub
