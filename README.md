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

* House-server prints hello-world when accessed via load-balancer service with basic auth.

* Fix tracing

TODO (Short term):
* Fix build
* Move DB password out of plain text in repo
* DB allows capitals
* Move DB to terraform with setup script
* Add license

* Logs, metrics, tracing, service-mesh, local k8s, databases, proxies

TODO (Long term):
* Sort running on AWS
* Add a cloud function to act on database (once a day check 'robbie' exists)
* Add something using pubsub
* Add something in Java
* Add a front end thing (& look into CSS + HTML + React/Redux)
