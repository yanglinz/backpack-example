#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

SOURCE_IMAGE="source-image"
DOCKER_REGISTRY="gcr.io/${GCP_PROJECT_ID}/${APP_NAME}"
RELEASE_TAG="$(. "$(dirname "$0")/hash-files.sh")"
HEROKU_APP_NAME="${APP_NAME}-backpack"

function debug_info() {
  sudo apt-get update && sudo apt-get install tree
  tree -d -I node_modules
}

function build_release() {
  docker build \
    -f .backpack/docker/python-prod.Dockerfile \
    --tag "$SOURCE_IMAGE" \
    .
}

function publish_gcp_registry() {
  gcloud auth configure-docker
  docker tag "$SOURCE_IMAGE" "$DOCKER_REGISTRY"
  docker tag "$SOURCE_IMAGE" "${DOCKER_REGISTRY}:${RELEASE_TAG}"
  docker push "$DOCKER_REGISTRY"
  docker push "${DOCKER_REGISTRY}:${RELEASE_TAG}"
}

function publish_deploy_heroku() {
  bash "$(dirname "$0")/install-heroku.sh"

  # Push to the container registry
  heroku container:login
  docker tag "$SOURCE_IMAGE" "registry.heroku.com/${HEROKU_APP_NAME}/web"
  docker tag "$SOURCE_IMAGE" "registry.heroku.com/${HEROKU_APP_NAME}/worker"
  docker push "registry.heroku.com/${HEROKU_APP_NAME}/web"
  docker push "registry.heroku.com/${HEROKU_APP_NAME}/worker"

 # Release the build
  heroku container:release web -a "$HEROKU_APP_NAME"
  heroku container:release worker -a "$HEROKU_APP_NAME"
}

function generate_do_artifact() {
  # Create env vars
  mkdir -p var/env
  ./backpack vars get --env=production
  ./backpack vars get --env=production > /dev/null 2>&1
  ./backpack vars get --env=production > var/env/production.json

  # Create tarball
  tar -czf /tmp/app-artifact.tar.gz .
  mv /tmp/app-artifact.tar.gz .
}

function publish_do_artifact() {
  # echo "$GCP_SERVICE_ACCOUNT_KEY" | base64 --decode > /tmp/service-account.json
  # export GCLOUD_KEYFILE_JSON="/tmp/service-account.json"

  cd terraform
  terraform init -backend-config="token=${TERRAFORM_CLOUD_TOKEN}"
  terraform plan -var="digitalocean_token=${DIGITALOCEAN_TOKEN}"
  # terraform apply -var="digitalocean_token=${DIGITALOCEAN_TOKEN}" -auto-approve
  terraform output -json > /tmp/terraform-output.json
  cd -

  local instance_ip=$(cat /tmp/terraform-output.json | jq -r '.ip_address.value')

  # Setup ansible hosts inventory
  mkdir -p ~/.ssh
  mkdir -p ./etc/ansible
  ssh-keyscan -H "$instance_ip" >> ~/.ssh/known_hosts
  echo "$instance_ip" > ./etc/ansible/hosts

  # Setup SSH key pair
  mkdir -p ~/.ssh
  echo "$DIGITALOCEAN_PRIVATE_KEY" > ~/.ssh/id_rsa
  sudo chmod 600 ~/.ssh/id_rsa
  ssh-keygen -y -f ~/.ssh/id_rsa > ~/.ssh/id_rsa.pub
  sudo chmod 600 ~/.ssh/id_rsa.pub
}

debug_info

if [[ "$RUNTIME_PLATFORM" == "CLOUD_RUN" ]]; then
  echo "Building release artifact for CLOUD_RUN"
  build_release
  publish_gcp_registry
elif [[ "$RUNTIME_PLATFORM" == "HEROKU" ]]; then
  echo "Building release artifact for HEROKU"
  build_release
  publish_deploy_heroku
elif [[ "$RUNTIME_PLATFORM" == "DIGITAL_OCEAN" ]]; then
  echo "Building release artifact for DIGITAL_OCEAN"
  generate_do_artifact
  publish_do_artifact
else
  echo "Nothing to publish for ${RUNTIME_PLATFORM}"
fi
