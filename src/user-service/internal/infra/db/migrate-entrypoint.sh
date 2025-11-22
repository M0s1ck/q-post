#!/bin/bash
set -e

echo "Creating schema '${POSTGRES_COMMUNITY_SCHEME}' if it does not exist..."
PGPASSWORD="${POSTGRES_PASSWORD}" psql -h "${POSTGRES_HOST}" -U "${POSTGRES_USER}" -d "${POSTGRES_DB}" \
    -c "CREATE SCHEMA IF NOT EXISTS ${POSTGRES_COMMUNITY_SCHEME};"

echo "running migrations in community schema"
migrate -path=/app/migrations \
    -database="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable&search_path=${POSTGRES_COMMUNITY_SCHEME}" \
    up
