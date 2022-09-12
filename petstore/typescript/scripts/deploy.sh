#!/bin/bash

NAME=$(node -p "require('./package.json').name");
VERSION=$(node -p "require('./package.json').version");

DOCKER_USERNAME="toxsickcoder"

docker build -t "$DOCKER_USERNAME/$NAME:$VERSION" -t "$DOCKER_USERNAME/$NAME:latest" .

docker login -u "$DOCKER_USERNAME"

docker push "$DOCKER_USERNAME/$NAME:$VERSION"
docker push "$DOCKER_USERNAME/$NAME:latest"
