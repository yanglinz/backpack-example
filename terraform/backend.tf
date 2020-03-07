terraform {
  backend "remote" {
    organization = "yanglin"

    workspaces {
      name = "backpack-example"
    }
  }
}
