# Auth Server

This is a toy auth server.
Currently, it only has a single endpoint `auth/<user>` that checks the user exists and returns success if they do.

It is designed to be run serverlessly on Google App Engine. Currently, it runs in the default service. 

## Commands

Run locally (exposing port 8080): `go run $REPO/service/auth/ --postgres_password=$PASSWORD`

Start cloud function: `gcloud app versions start $VERSION`, where available versions can be found from `gcloud app versions list`.

Stop cloud function: `gcloud app versions stop $VERSION`.

Deploy new version: `gcloud app deploy $REPO/services/auth/ --version v1`