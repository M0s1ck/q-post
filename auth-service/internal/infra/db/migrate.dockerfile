FROM alpine:latest

RUN apk add --no-cache curl bash

WORKDIR /app

# installing golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz \
    | tar xvz && mv migrate /usr/local/bin/migrate

COPY ./migrations ./migrations
COPY ./migrate-entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]