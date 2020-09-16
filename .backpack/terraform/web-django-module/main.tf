locals {
  docker_image        = "gcr.io/${var.gcp_project}/${var.context_name}:${var.image_tag}"
  berglas_secret_path = "berglas-${var.context_name}/BERGLAS_APP_JSON"
  service_account     = "berglas-${var.context_name}@${var.gcp_project}.iam.gserviceaccount.com"
}
