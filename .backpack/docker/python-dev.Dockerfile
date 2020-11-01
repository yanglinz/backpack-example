FROM python:3.7

# Install watchman
# This step takes a LONG time, so cache this layer first
ARG DISABLE_WATCHMAN
COPY .backpack/docker/watchman/install-watchman.sh /tmp/
RUN /tmp/install-watchman.sh

ENV PYTHONUNBUFFERED 1

# Install gcloud
COPY .backpack/docker/scripts/install-gcloud.sh /tmp/
RUN /tmp/install-gcloud.sh

# Install dependencies
RUN apt-get update && apt-get install -y \
  gettext-base \
  supervisor \
  nginx \
  git \
  postgresql-client

# Stop nginx
RUN service nginx stop

# Install berglas
COPY --from=gcr.io/berglas/berglas:0.5.0 /bin/berglas /bin/berglas

# Setup the working directory
RUN mkdir /app && mkdir /home/app
WORKDIR /app
ENV HOME /home/app

# Install application dependencies
COPY .backpack/docker/scripts/get-poetry.py /tmp/
RUN python /tmp/get-poetry.py --version=1.0.5
ENV PATH="${HOME}/.poetry/bin:${PATH}"
COPY pyproject.toml /app/
COPY poetry.lock /app/
RUN poetry install

# Install custom dependencies
COPY scripts* /app/scripts
COPY .backpack/docker/scripts/install-extra-deps.sh /tmp/
RUN /tmp/install-extra-deps.sh

# Copy configurations
COPY .backpack/docker/watchman/watchman.json /etc/
COPY .backpack/docker/supervisord/supervisord-dev.conf /etc/supervisord.conf

# Copy application code
COPY . /app
