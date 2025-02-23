#!/bin/sh
host="$1"
port="$2"
shift 2

until nc -z "$host" "$port"; do
  echo "Waiting for PostgreSQL..."
  sleep 1
done

echo "PostgreSQL is ready!"
exec "$@"