#!/bin/bash
set -e

echo "Esperando a que CockroachDB esté disponible..."
sleep 8

echo "Creando base de datos 'stocksdb'..."
cockroach sql --insecure --host=db -e "CREATE DATABASE IF NOT EXISTS stocksdb;"

echo "Creando tabla 'stocks'..."
cockroach sql --insecure --host=db -d stocksdb -e "
CREATE TABLE IF NOT EXISTS stocks (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  ticker STRING,
  company STRING,
  brokerage STRING,
  action STRING,
  rating_from FLOAT,
  rating_to FLOAT,
  target_from FLOAT,
  target_to FLOAT,
  created_at TIMESTAMP DEFAULT now()
);
"

echo "Base de datos y tabla 'stocks' creadas correctamente."
