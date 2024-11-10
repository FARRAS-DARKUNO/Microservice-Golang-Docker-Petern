#!/bin/bash

# Wait for PostgreSQL to be ready
host="$1"
port="$2"
shift 2
cmd="$@"

# Wait for the PostgreSQL service to be available
until pg_isready -h "$host" -p "$port" -U postgres; do
  echo "Waiting for PostgreSQL at $host:$port..."
  sleep 2
done

echo "PostgreSQL at $host:$port is ready!"

# Execute the command passed to the script (your service)
exec $cmd