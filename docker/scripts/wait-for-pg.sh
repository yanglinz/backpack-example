#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

HOST="postgres"
PASSWORD="postgres"
USER="postgres"

until PGPASSWORD="$PASSWORD" psql -h "$HOST" -U "$USER" -c '\q'; do
  >&2 echo "postgres is unavailable - sleeping"
  sleep 1
done
