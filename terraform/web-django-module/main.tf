locals {
  docker_image        = "gcr.io/${var.gcp_project}/${var.context_name}:${var.image_tag}"
  berglas_secret_path = "backpack-berglas-${var.context_name}/BERGLAS_APP_JSON"
  service_account     = "cloudrun-berglas-${var.context_name}@${var.gcp_project}.iam.gserviceaccount.com"
}

resource "google_cloud_run_service" "cloud_run" {
  name     = "${var.context_name}-${var.project_name}-service"
  location = "us-central1"

  template {
    spec {
      containers {
        image = local.docker_image
        env {
          name  = "DJANGO_SETTINGS_MODULE"
          value = var.django_settings_module
        }
        env {
          name  = "BERGLAS_SECRET_PATH"
          value = local.berglas_secret_path
        }
      }
      service_account_name = local.service_account
    }
  }
}

data "google_iam_policy" "cloud_run_policy" {
  binding {
    role    = "roles/run.invoker"
    members = ["allUsers"]
  }
}

resource "google_cloud_run_service_iam_policy" "cloud_run_policy" {
  location    = google_cloud_run_service.cloud_run.location
  project     = google_cloud_run_service.cloud_run.project
  service     = google_cloud_run_service.cloud_run.name
  policy_data = data.google_iam_policy.cloud_run_policy.policy_data
}
