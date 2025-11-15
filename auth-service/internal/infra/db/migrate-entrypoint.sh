#!/bin/bash
set -e

echo "running create auth schema migration in public"
migrate -path=/migrations \
    -database="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable&search_path=public" \
    up 1

echo "running all other migrations in auth schema"
migrate -path=/migrations \
    -database="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable&search_path=${POSTGRES_AUTH_SCHEME}" \
    up
