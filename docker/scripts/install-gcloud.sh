#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# Add the Cloud SDK distribution URI as a package source
echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] http://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list

# Import the Google Cloud Platform public key
curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -

# Update the package list and install the Cloud SDK
apt-get update && apt-get install -y google-cloud-sdk

# Install cloud sql proxy
curl https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64 > /bin/cloud_sql_proxy
