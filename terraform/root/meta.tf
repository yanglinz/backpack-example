terraform {
  required_version = ">= 0.13"

  backend "remote" {
    organization = "{{BACKPACK_DEFAULT_ORG}}"

    workspaces {
      name = "{{BACKPACK_WORKSPACE}}"
    }
  }

  required_providers {
    digitalocean = {
      source = "digitalocean/digitalocean"
    }
  }
}

variable "digitalocean_token" {
  type = string
}

provider "digitalocean" {
  token = var.digitalocean_token
}
