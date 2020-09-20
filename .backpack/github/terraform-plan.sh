#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

echo "$GCP_SERVICE_ACCOUNT_KEY" | base64 --decode > /tmp/service-account.json
export GCLOUD_KEYFILE_JSON="/tmp/service-account.json"

cd terraform
terraform init -backend-config="token=${TERRAFORM_CLOUD_TOKEN}"
terraform plan -var="digitalocean_token=${DIGITALOCEAN_TOKEN}"
cd -
