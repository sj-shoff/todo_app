# todo_app
docker run --name=todo-db -e POSTGRES_PASSWORD='03032006' -p 5436:5432 -d --rm postgres
docker ps
миграции - а-ля система контроля версий для баз данных
import "github.com/golang-migrate/migrate/v4" (библиотека для миграций)
mkdir -p ./schema
migrate create -ext sql -dir ./schema -seq init