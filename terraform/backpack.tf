provider "google" {
  project = "default-263000"
  zone    = "us-central1"
  region  = "us-central1-c"
}

variable "context_name" {
  type    = "string"
  default = "backpack-example"
}

variable "image_tag" {
  type    = "string"
  default = "latest"
}

module "app_web" {
  source                 = "../.backpack/terraform/web-django-module"
  context_name           = "${var.context_name}"
  project_name           = "core"
  django_settings_module = "projects.example.settings"
  image_tag              = "${var.image_tag}"
  gcp_project            = "default-263000"
}
