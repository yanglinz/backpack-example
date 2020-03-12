# Backpack 🎒

Backpack is a tool I made for personal project that helps me sharing development
and runtime configuration for docker based applications.

Mechanically, it's a meant to be hosted in application repositories as a
[git subtree](https://www.atlassian.com/git/tutorials/git-subtree), and
functions as CLI tool that ties together bash scripts and configuration files.

> NOTE: Backpack is made for my personal use, and is very specific to my own
> setup. People may not be able to use it directly, but I did want to open
> source it because the pattern of embedding a set of configuration as a
> packaged tool has worked for me very well over time.

## Features

Backpack takes care of generating configurations for:

- Local development via [Docker](https://www.docker.com/).
- Continuous integration via
  [Github Actions](https://github.com/features/actions).
- Infrastructure as code via [Terraform](https://www.terraform.io/).
- Serverless runtime via [Google Cloud Run](https://cloud.google.com/run).
- Secret management via
  [Berglas](https://github.com/GoogleCloudPlatform/berglas).
