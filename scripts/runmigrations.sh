#!/bin/bash

set -e

DB_NAME="retro_board"
DB_USER="dev"
DB_PASSWORD="senhadodev"
MIGRATIONS_PATH="./migrations"

echo "Aplicando migrations..."
migrate -path="$MIGRATIONS_PATH" -database="postgres://$DB_USER:$DB_PASSWORD@localhost:5432/$DB_NAME?sslmode=disable" up
