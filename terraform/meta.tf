terraform {
  required_version = ">= 0.13"

  backend "remote" {
    organization = "yanglin"

    workspaces {
      name = "backpack-example"
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
