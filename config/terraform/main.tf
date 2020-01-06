terraform {
  required_version = "~> 0.12.0"

  backend "remote" {
    organization = "robbieheywood"

    workspaces {
      name = "primary"
    }
  }
}

provider "google" {
  credentials = file(var.credentials_file)

  project = var.project
  region  = var.region
}

resource "google_container_cluster" "primary-cluster" {
  name               = "gke-${var.region}-1"
  location           = var.region

  # Use node pool specified below instead of the default node pool
  remove_default_node_pool = true
  initial_node_count       = 1

  master_auth {
    username = ""
    password = ""

    client_certificate_config {
      issue_client_certificate = false
    }
  }
}

resource "google_container_node_pool" "primary-nodes" {
  name       = "gke-${var.region}-nodes-1"
  location   = var.region
  cluster    = google_container_cluster.primary-cluster.name
  node_count = 2

  node_config {
    # Use preemptible nodes as they are cheaper
    preemptible  = true
    machine_type = "n1-standard-1"

    oauth_scopes = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]
  }
}