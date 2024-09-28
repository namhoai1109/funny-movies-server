#!/usr/bin/env bash

# Ensure the database container is online and usable
echo "Waiting for database..."
until docker exec -i funnymovies.db psql -h localhost -U funnymovies -d funnymovies -c "SELECT 1" &> /dev/null
do
  # printf "."
  sleep 1
done