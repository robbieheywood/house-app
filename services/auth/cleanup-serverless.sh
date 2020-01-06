#!/bin/bash

# Cleanup serverless app from app-engine

set -o nounset
set -o pipefail
set -o errexit
set -o xtrace

# Cleanup auth cloud-function server
gcloud app versions stop v1