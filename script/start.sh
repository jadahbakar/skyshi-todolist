#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "mysql://todo:secret@tcp(mysql:3306)/todolist" -verbose up

echo "start app"
exec "$@"