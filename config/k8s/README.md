# Kubernetes Config

This directory holds the kubernetes config for the apps in this repo.
We use kustomize to manage the configuration.

## Running in GKE

* Ensure the kubectl config is setup correctly by running `gcloud container clusters get-credentials $CLUSTER_NAME --zone $ZONE` if not.
* Run `kubectl apply -k $REPO/config/$SERVICE/overlays/cloud` for each service you want, where `$SERVICE` is 'house', 'fluentd' etc..

## Running locally

Kubernetes can be run locally using several different methods, as detailed below.
In order to allow pulling the images from GCR, ensure that you have copied the GCR service account key to ~/gcr-key.json.

*Note: the local kubernetes methods described below were just me playing around with each of them.
There's no guarantee this app will continue to work with with any/all of them.*

### Kind (Kubernetes in Docker)

Steps:
* Install [Kind](https://github.com/kubernetes-sigs/kind).
* Ensure `~/gcr-key.json` has the service account key.
* Run `kind create cluster`
* Run `kubectl create secret docker-registry gcr-json-key --docker-server=eu.gcr.io --docker-username=_json_key --docker-password="$(cat $HOME/gcr-key.json)" --docker-email=any@valid.email`
to create the secret containing the GCR key used to pull images.
* Run `kubectl apply -k $REPO/config/$SERVICE/overlays/cloud` for each service you want, where `$SERVICE` is 'house', 'fluentd' etc..
* To cleanup, run `kind delete cluster`.

### MiniKube

Steps:
* Install [Minikube](https://minikube.sigs.k8s.io/)
* Ensure `~/gcr-key.json` has the service account key.
* Run `minikube start`
* Run `kubectl create secret docker-registry gcr-json-key --docker-server=eu.gcr.io --docker-username=_json_key --docker-password="$(cat $HOME/gcr-key.json)" --docker-email=any@valid.email`
to create the secret containing the GCR key used to pull images.
* Run `kubectl apply -k $REPO/config/$SERVICE/overlays/cloud` for each service you want, where `$SERVICE` is 'house', 'fluentd' etc..
* To cleanup, run `minikube delete`.

#### Others

You can also use docker-for-mac, k3s/k3d or microk8s by following pretty much the same steps as above.