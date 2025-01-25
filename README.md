# todo_app

### Работа с докером:
```bash
docker run --name=todo-db -e POSTGRES_PASSWORD='03032006' -p 5436:5432 -d --rm postgres
docker ps
```

### Миграции - а-ля система контроля версий для баз данных
```bash
import "github.com/golang-migrate/migrate/v4" 
mkdir -p ./schema
migrate create -ext sql -dir ./schema -seq init
```

### Применение миграций:
```bash
migrate -path ./schema -database 'postgres://postgres:03032006@localhost:5436/postgres?sslmode=disable' up
```

### Запуск БД в контейнере(1я строка - берем айдишник postgres)
```bash
docker ps
docker exec -it 6188ec4b2fdd /bin/bash
psql -U postgres
\d
```
