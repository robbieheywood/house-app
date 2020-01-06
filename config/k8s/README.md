# Kubernetes Config

This directory holds the kubernetes config for the apps in this repo.

## How to use
* Ensure the kubectl config is setup correctly - run `gcloud container clusters get-credentials $CLUSTER_NAME --zone 
$ZONE` if not.
* Run `kubectl apply -f $REPO/config/k8s/ --recursive` to apply the config. 