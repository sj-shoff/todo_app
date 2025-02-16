#!/bin/sh

host="$1"
port="$2"
cmd="$3"

until pg_isready -h "$host" -p "$port"; do
  echo "Postgres is unavailable - sleeping"
  sleep 2
done

echo "Postgres is up - executing command"
exec "$cmd"