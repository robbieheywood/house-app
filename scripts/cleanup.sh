#!/bin/bash

set -o nounset
set -o errexit

# Cleanup auth server
gcloud app versions stop v1