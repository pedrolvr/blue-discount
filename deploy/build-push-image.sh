#!/bin/bash

set -eo pipefail

COMMIT_ID="$(git rev-parse --short $CIRCLE_SHA1)"
IMAGE_NAME=phenrigomes/blue-discount:$CIRCLE_BRANCH-$COMMIT_ID
docker build -t $IMAGE_NAME .
docker login -u $DOCKERHUB_USERNAME -p $DOCKERHUB_PASSWORD
docker push $IMAGE_NAME
