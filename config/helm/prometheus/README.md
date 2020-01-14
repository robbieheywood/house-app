# Helm

Helm is used to deploy Prometheus as I wanted to try it out.

To deploy, run `helm install stable/prometheus-operator --generate-name -f $REPO/config/helm/prometheus/values.yaml`