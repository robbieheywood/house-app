# Kubernetes Config

This directory holds the kubernetes config for the apps in this repo.

Note: we use Helm for Prometheus as I wanted to try Helm out.
For Prometheus, see the Helm directory.

## Gotchas
* Note that the Jaeger CRDs are so big that you need to use `kubectl create/replace` rather than `kubectl apply`
as other wise the `last-applied-configuration` annotation is too large.

## Running in GKE

* Ensure the kubectl config is setup correctly by running `gcloud container clusters get-credentials $CLUSTER_NAME --zone $ZONE` if not.
* Run `kubectl create secret generic jaeger-secret --from-literal=ES_PASSWORD=${ELASTIC_PASSWORD} --from-literal=ES_USERNAME=elastic`
  * The value for `${ELASTIC_PASSWORD}` can be found by running `kubectl get secret elasticsearch-es-elastic-user -o=jsonpath='{.data.elastic}' | base64 --decode; echo`
  after deploying elastic using the step below.
* Run `kubectl create -k $REPO/config/$SERVICE/overlays/cloud` for each service you want, where `$SERVICE` is 'house', 'fluentd' etc..

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
* Run `kubectl create secret generic jaeger-secret --from-literal=ES_PASSWORD=${ELASTIC_PASSWORD} --from-literal=ES_USERNAME=elastic` to create the
secret Jaeger needs to speak to ElasticSearch.
* Run `kubectl create -k $REPO/config/$SERVICE/overlays/cloud` for each service you want, where `$SERVICE` is 'house', 'fluentd' etc..
* To cleanup, run `kind delete cluster`.

### MiniKube

Steps:
* Install [Minikube](https://minikube.sigs.k8s.io/)
* Ensure `~/gcr-key.json` has the service account key.
* Run `minikube start`
* Run `kubectl create secret docker-registry gcr-json-key --docker-server=eu.gcr.io --docker-username=_json_key --docker-password="$(cat $HOME/gcr-key.json)" --docker-email=any@valid.email`
to create the secret containing the GCR key used to pull images.
* Run `kubectl create secret generic jaeger-secret --from-literal=ES_PASSWORD=${ELASTIC_PASSWORD} --from-literal=ES_USERNAME=elastic` to create the
secret Jaeger needs to speak to ElasticSearch.
* Run `kubectl create -k $REPO/config/$SERVICE/overlays/cloud` for each service you want, where `$SERVICE` is 'house', 'fluentd' etc..
* To cleanup, run `minikube delete`.

#### Others

You can also use docker-for-mac, k3s/k3d or microk8s by following pretty much the same steps as above.