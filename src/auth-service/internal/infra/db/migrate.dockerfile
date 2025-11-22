FROM alpine:latest

# installing golang-migrate and psql
RUN apk add --no-cache curl bash postgresql-client
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz \
    | tar xvz && mv migrate /usr/local/bin/migrate

WORKDIR /app

COPY migrations ./migrations
COPY migrate-entrypoint.sh ./entrypoint.sh

RUN chmod +x ./entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]