resource "digitalocean_droplet" "web" {
  image  = "ubuntu-18-04-x64"
  name   = "backpack-${var.app_context.app_name}"
  region = "nyc3"
  size   = "s-1vcpu-1gb"
}

resource "digitalocean_floating_ip" "web" {
  droplet_id = digitalocean_droplet.web.id
  region     = digitalocean_droplet.web.region
}
