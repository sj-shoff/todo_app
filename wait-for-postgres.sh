#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
port="$2"
user="$3"
dbname="$4"
sslmode="$5"
shift 5
cmd="$@"

until PGPASSWORD=$DB_PASSWORD psql -h "$host" -p "$port" -U "$user" -d "$dbname" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd