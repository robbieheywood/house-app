#!/bin/bash

# Cleanup serverless app from app-engine

set -o nounset -o pipefail -o errexit -o xtrace

# Cleanup auth server
gcloud app versions stop v1