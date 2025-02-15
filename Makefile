# Загрузка переменных окружения из .env
include .env
export

# Сборка Docker-контейнера
build:
	docker-compose build

# Запуск контейнера
run:
	docker-compose up

# Миграции базы данных
migrate:
	docker-compose run app migrate -path ./schema -database 'postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}' up

# Генерация Swagger документации
swag:
	docker-compose run app swag init -g cmd/main.go