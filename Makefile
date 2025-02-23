include .env
build:
	docker-compose build todo-app

run:
	docker-compose up todo-app

migrate:
	migrate -path ./schema -database 'postgres://postgres:${DB_PASSWORD}@0.0.0.0:5432/postgres?sslmode=disable' up

swag:
	swag init -g cmd/main.go