#!/bin/bash

# Rollout serverless app to app-engine

set -o nounset
set -o pipefail
set -o errexit

# Rollout auth server
cd go/serverless
gcloud app versions start v1 @@@ ADD ENV VARIABLE
gcloud app deploy --version v1