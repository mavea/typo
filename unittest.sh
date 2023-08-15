#!/bin/sh

echo "
version: '3.8'

services:
  project:
    image: golang:1.20-alpine
    container_name: typo
    volumes:
      - .:/opt/app/src
    working_dir: /opt/app/src
    tty: true
" > ./docker-compose.yml
docker compose up -d
docker exec -it typo sh -c 'go test -v ./...'
docker compose down
rm ./docker-compose.yml