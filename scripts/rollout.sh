#!/bin/bash

set -o nounset
set -o errexit

# Ensure that the environment variable $HOUSE-APP-CONNECTION has the postgres db connection details &
GOOGLE_APPLICATION_CREDENTIALS has service account details
HOUSE-APP-CONNECTION = user=postgres password=157AbbeyRoad dbname=house-users host=tensile-imprint-156310:europe-west1:house-users
GOOGLE_APPLICATION_CREDENTIALS =

# Rollout auth server
cd go/serverless/auth
gcloud app versions start v1 @@@ ADD ENV VARIABLE
gcloud app deploy --version v1