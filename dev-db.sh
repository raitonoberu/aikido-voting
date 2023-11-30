#!/bin/bash

source .env
podman run --name postgres-dev \
  --volume postgres:/var/lib/postgresql/data \
  -e POSTGRES_USER=${POSTGRES_USER} \
  -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
  -e POSTGRES_DB=aikido-db-dev \
  -p 5432:5432 \
  --cap-add NET_RAW \
  docker.io/postgres:alpine
