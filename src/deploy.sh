#!/bin/bash
set -e

# always run from script's directory
cd "$(dirname "$0")"
pwd
ls

# get env variables from file
source ./env

# login to docker hub
echo "${DOCKERHUB_TOKEN}" | docker login -u "${DOCKERHUB_USERNAME}" --password-stdin

# pull images
docker compose -f docker-compose.prod.yaml pull

# bring up db
docker compose -f docker-compose.prod.yaml up psg -d
sleep 10

# apply migrations
docker compose -f docker-compose.prod.yaml run --rm auth-migrate
docker compose -f docker-compose.prod.yaml run --rm user-migrate

# bring up services
docker compose -f docker-compose.prod.yaml up -d
sleep 10

docker compose -f docker-compose.prod.yaml logs --tail=200