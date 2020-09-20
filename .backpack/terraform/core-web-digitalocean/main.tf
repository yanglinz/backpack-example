locals {
  # We can and should set this dynamically by fetching it from 
  # the digitalocean api, but this will do for now
  default_ssh_key_ids = ["27895764"]
}

resource "digitalocean_droplet" "web" {
  image    = "ubuntu-18-04-x64"
  name     = "backpack-${var.app_context.app_name}"
  region   = "nyc3"
  size     = "s-1vcpu-1gb"
  ssh_keys = local.default_ssh_key_ids
}

resource "digitalocean_floating_ip" "web" {
  droplet_id = digitalocean_droplet.web.id
  region     = digitalocean_droplet.web.region
}

output "ip_address" {
  value = digitalocean_floating_ip.web.ip_address
}
