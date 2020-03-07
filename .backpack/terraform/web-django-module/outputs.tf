output "cloudrun_context" {
  value = google_cloud_run_service.cloud_run.status
}

output "docker_image" {
  value = local.docker_image
}
