#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# Start the development server
pipenv run python manage.py runserver 0.0.0.0:4567
