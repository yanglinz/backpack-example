#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

uwsgi \
  --wsgi-file $(pwd)/.backpack/docker/uwsgi/wsgi.py \
  --check-static $(pwd)/www \
  $(pwd)/.backpack/docker/uwsgi/uwsgi.ini
