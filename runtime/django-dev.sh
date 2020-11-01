#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# Start the development server
poetry run python manage.py runserver 0.0.0.0:4567
