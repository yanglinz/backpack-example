# 🎒Backpack Example

This is an example app that embeds
[backpack](https://github.com/yanglinz/backpack), which is a tool I wrote to
share a set of common infrastructure configuration for some personal projects.

## Features

It takes care of generating configurations for:

- Local development via [Docker](https://www.docker.com/).
- Continuous integration via
  [Github Actions](https://github.com/features/actions).
- Infrastructure as code via [Terraform](https://www.terraform.io/).
- Serverless runtime via [Google Cloud Run](https://cloud.google.com/run).
- Secret management via
  [Berglas](https://github.com/GoogleCloudPlatform/berglas).

## Application

This particular example is a hello world
[Django](https://www.djangoproject.com/) application deployed to
[Google Cloud Run](https://cloud.google.com/run), which is a serverless
container runtime.
