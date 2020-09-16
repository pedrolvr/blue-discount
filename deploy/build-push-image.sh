#!/bin/bash

set -eo pipefail

COMMIT_ID="$(git rev-parse --short $CIRCLE_SHA1)"
IMAGE_NAME=phenrigomes/blue-discount:$CIRCLE_BRANCH-$COMMIT_ID
docker build -t $IMAGE_NAME .
echo "$DOCKERHUB_PASSWORD" | docker login -u "$DOCKERHUB_USERNAME" --password-stdin
docker push $IMAGE_NAME
