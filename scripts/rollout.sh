#!/bin/bash

set -o nounset
set -o errexit

# Rollout auth server
cd go/serverless/auth
gcloud app versions start v1
gcloud app deploy --version v1