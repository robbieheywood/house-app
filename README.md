[![house-server GoDoc](https://godoc.org/github.com/robbieheywood/house-app/go/house-server?status.svg)](https://godoc.org/github.com/robbieheywood/house-app/go/house-server)
[![house-server CircleCI status](https://circleci.com/gh/robbieheywood/house-app.png?circle-token=:circle-token "CircleCI status")](https://circleci.com/gh/robbieheywood/house-app)

# HouseApp

A simple app for displaying useful household information.

Note: this app is a vehicle for me to practice with new technologies - 
it is definitely NOT the best implementation for this and is deliberately crazy/over-engineered.

## Setup

Ensure the following are installed:
* Go
* Docker
* Gcloud
* Kubectl
* Terraform

Then follow these steps:
* Follow the steps at `config/terraform/README.md` to create the GKE cluster.
* Follow the steps at `config/k8s/README.md` to create the required services in GKE.

## Status

* House-server prints hello-world from 

TODO:
* Fix build
* Sort gcloud auth (gcloud auth application-default login)
* Sort rollout with env variables
* DB allows capitals

* Add license
* Setup logs and metrics and dashboard
* Cloud function acts on database
* k8s-ize Go backend + rollout to GKE
* Something in Java (& recap)
* Front end (& look into CSS + HTML + React/Redis)
* Add pub/sub
* Add redis (+ look into it and redisearch)
* Add nginx (+ look into it)
* Allow moving to AWS
