locals {
  # Ideally we should set the key id dynamically by fetching it from the DO 
  # API, but since it changes so infrequently, hardcoding it is fine for now.
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

resource "digitalocean_firewall" "web" {
  name = "backpack-web-firewall"

  droplet_ids = [digitalocean_droplet.web.id]

  inbound_rule {
    protocol         = "tcp"
    port_range       = "22"
    source_addresses = ["0.0.0.0/0", "::/0"]
  }

  inbound_rule {
    protocol         = "tcp"
    port_range       = "80"
    source_addresses = ["0.0.0.0/0", "::/0"]
  }

  inbound_rule {
    protocol         = "tcp"
    port_range       = "443"
    source_addresses = ["0.0.0.0/0", "::/0"]
  }

  inbound_rule {
    protocol         = "icmp"
    source_addresses = ["0.0.0.0/0", "::/0"]
  }

  outbound_rule {
    protocol         = "tcp"
    port_range       = "80"
    destination_addresses = ["0.0.0.0/0", "::/0"]
  }

  outbound_rule {
    protocol         = "tcp"
    port_range       = "443"
    destination_addresses = ["0.0.0.0/0", "::/0"]
  }

  outbound_rule {
    protocol              = "tcp"
    port_range            = "53"
    destination_addresses = ["0.0.0.0/0", "::/0"]
  }

  outbound_rule {
    protocol              = "udp"
    port_range            = "53"
    destination_addresses = ["0.0.0.0/0", "::/0"]
  }

  outbound_rule {
    protocol              = "icmp"
    destination_addresses = ["0.0.0.0/0", "::/0"]
  }
}

output "ip_address" {
  value = digitalocean_floating_ip.web.ip_address
}
