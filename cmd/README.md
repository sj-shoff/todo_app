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

### Проверка БД в контейнере(1я строка - берем айдишник postgres)
```bash
docker ps
docker exec -it 6188ec4b2fdd /bin/bash
psql -U postgres
\d
```
### Swagger
```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g cmd/main.go
```
### Чтобы посмотреть документацию Swagger: http://localhost:8000/swagger/index.html

### Скрипт wait-for-postgres.sh гарантирует, что приложение запустится только после готовности базы данных.

### Проверка занятости порта / освобождение порта
```bash
sudo lsof -i :5432
sudo kill 474
```

### Права для работы с папкой
```bash
sudo chmod -R 777 .database
```

### Очистка docker-compose (вместе с кэшем)
```bash
docker-compose down -v --rmi all
docker system prune -a --volumes
```

### Отключение Firewall
```bash
sudo ufw disable
```

### Проверка логов в докере
```bash
docker logs todo_app_todo-app_1
```

### Подключение к базе данных внутри контейнера Docker и изменение пароля
```bash
docker-compose exec db psql -U postgres -d postgres
ALTER USER postgres WITH PASSWORD '03032006';
\q
```