#!/bin/bash

while !</dev/tcp/${DB_HOST}/${DB_PORT}; do sleep 1; done;

DBSTRING="postgresql://${DB_ROOT}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

goose postgres "$DBSTRING" up