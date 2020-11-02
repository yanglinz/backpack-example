#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

RELEASE_BRANCH="main"
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

function generate_vm_artifact() {
  # Create env vars
  mkdir -p var/env
  ./backpack vars get --env=production
  ./backpack vars get --env=production > /dev/null 2>&1
  ./backpack vars get --env=production > var/env/production.json

  # Create tarball
  tar -czf /tmp/app-artifact.tar.gz .
  mv /tmp/app-artifact.tar.gz .
}

debug_info

if [[ "$GITHUB_REF" != "refs/heads/${RELEASE_BRANCH}" ]]; then
  echo "Not on ${RELEASE_BRANCH} branch. Nothing to publish."
elif [[ "$RUNTIME_PLATFORM" == "CLOUD_RUN" ]]; then
  echo "Building release artifact for CLOUD_RUN"
  build_release
  publish_gcp_registry
elif [[ "$RUNTIME_PLATFORM" == "HEROKU" ]]; then
  echo "Building release artifact for HEROKU"
  build_release
  publish_deploy_heroku
elif [[ "$RUNTIME_PLATFORM" == "VM" ]]; then
  echo "Building release artifact for VM"
  generate_vm_artifact
else
  echo "Not on GCP or Heroku. Nothing to publish."
fi
