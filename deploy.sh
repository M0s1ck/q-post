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

sudo docker compose -f docker-compose.prod.yaml pull

sudo docker compose up psg -d

sleep 10

sudo docker compose run --rm auth-migrate
sudo docker compose run --rm user-migrate

sudo docker compose -f docker-compose.prod.yaml up -d

sleep 10

sudo docker compose logs --tail=200