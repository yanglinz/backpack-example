#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# Manage static files
python manage.py collectstatic --noinput
[ -d "./build" ] && cp -r ./build/. ./www
rm -f ./www/index.html  # Disallow index.html in static files

# Run migration
python manage.py migrate

# Start production server
uwsgi --ini /app/.backpack/docker/uwsgi/uwsgi.ini
