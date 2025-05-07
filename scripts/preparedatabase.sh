#!/bin/bash
set -e  

CONTAINER_NAME=718b91b57235
DB_NAME=retro_board
DB_ROOTUSER="usuario"
DB_USER="dev"
DB_PASSWORD="senhadodev"

# 1. Acessar container postgresql
docker exec -i "$CONTAINER_NAME" psql -U $DB_ROOTUSER <<EOF

-- 2. Criar banco de dados
CREATE DATABASE $DB_NAME;

-- 3. Criar usuário de desenvolvimento
CREATE ROLE $DB_USER WITH LOGIN PASSWORD '$DB_PASSWORD';

-- 4. Criar extensão usada no projeto
\c $DB_NAME;
CREATE EXTENSION IF NOT EXISTS citext;


-- 5. Garantir privilégios ao usuário de desenvolvimento
ALTER DATABASE $DB_NAME OWNER TO $DB_USER;

EOF

echo "Configuração concluída com sucesso!"
