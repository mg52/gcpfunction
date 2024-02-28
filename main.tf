terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 4.34.0"
    }
  }
}

locals {
  projectId = "project-id" # Google Cloud Platform Project ID
}

provider "google" {
  project = local.projectId
  region  = "us-west1"
}

resource "random_id" "default" {
  byte_length = 8
}

resource "google_storage_bucket" "default" {
  name                        = "${random_id.default.hex}-gcf-source" # Every bucket name must be globally unique
  location                    = "us-west1"
  uniform_bucket_level_access = true
}

data "archive_file" "default" {
  type        = "zip"
  output_path = "/tmp/function-source.zip"
  source_dir  = "function/"
}
resource "google_storage_bucket_object" "object" {
  name   = "function-source.zip"
  bucket = google_storage_bucket.default.name
  source = data.archive_file.default.output_path # Add path to the zipped function source code
}

resource "google_cloudfunctions2_function" "default" {
  name        = "go-http-function-2"
  location    = "us-west1"
  description = "a new function"

  build_config {
    runtime     = "go121"
    entry_point = "router" # Set the entry point
    source {
      storage_source {
        bucket = google_storage_bucket.default.name
        object = google_storage_bucket_object.object.name
        generation = tonumber(regex("generation=(\\d+)", google_storage_bucket_object.object.media_link)[0])
      }
    }
  }

  service_config {
    max_instance_count = 100
    available_memory   = "256M"
    timeout_seconds    = 60
    ingress_settings = "ALLOW_ALL"
    environment_variables = {
        SERVICE_CONFIG_TEST = "config_test"
    }
    secret_environment_variables {
      key        = "DB_CONNECTION_STRING"
      project_id = local.projectId
      secret     = "db-connection-string"
      version    = "latest"
    }
  }

  depends_on = [ google_storage_bucket_object.object ]
}

resource "google_cloud_run_service_iam_member" "member" {
  location = google_cloudfunctions2_function.default.location
  service  = google_cloudfunctions2_function.default.name
  role     = "roles/run.invoker"
  member   = "allUsers"
}

output "function_uri" {
  value = google_cloudfunctions2_function.default.service_config[0].uri
}