#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

IMAGE_TAG="$(. "$(dirname "$0")/hash-files.sh")"

echo "$GCP_SERVICE_ACCOUNT_KEY" | base64 --decode > /tmp/service-account.json
export GCLOUD_KEYFILE_JSON="/tmp/service-account.json"

cd terraform
terraform init -backend-config="token=${TERRAFORM_CLOUD_TOKEN}"
TF_VAR_image_tag="$IMAGE_TAG" terraform plan
cd -
