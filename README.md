# REST API Для Создания TODO Списков на Go

## Стек: Go, <a href="https://github.com/gin-gonic/gin">Gin</a>, PostgreSQL, Docker, <a href="https://github.com/jmoiron/sqlx">sqlx</a>, <a href="https://github.com/golang-migrate/migrate">golang-migrate</a>, <a href="https://github.com/spf13/viper">Viper</a>, JWT, <a href="https://github.com/swaggo/swag">Swagger</a>, Slog, <a href="https://github.com/stretchr/testify">testify</a>, Makefile

## Ключевые достижения и реализованные функции:

### REST API на Go:
- Разработал сервер на фреймворке <a href="https://github.com/gin-gonic/gin">Gin</a> с эндпоинтами для управления задачами (создание, обновление, удаление, просмотр).
- Реализовал **чистую архитектуру** с разделением на слои: HTTP-обработчики, бизнес-логика, репозиторий для работы с БД.
- Настроил **graceful shutdown** для безопасного завершения работы сервера и соединений с БД.
- Добавил Swagger-документацию.

### Работа с данными:
- Использовал **PostgreSQL** для хранения задач и метаданных пользователей.
- Реализовал миграции БД с помощью библиотеки <a href="https://github.com/golang-migrate/migrate">golang-migrate</a> для управления схемой данных.
- Организовал CRUD-операции через библиотеку <a href="https://github.com/jmoiron/sqlx">sqlx</a> с защитой от SQL-инъекций.

### Аутентификация и безопасность:
- Внедрил аутентификацию через **JWT** (регистрация, вход, проверка токена).
- Добавил middleware для авторизации защищенных эндпоинтов (например, изменение/удаление задач).
- Настроил валидацию входных данных для предотвращения некорректных запросов.

### Конфигурация:
- Реализовал гибкую настройку приложения через **environment variables** и файлы конфигурации с использованием <a href="https://github.com/spf13/viper">Viper</a>.


### Инфраструктура:
- Завернул приложение и БД в **Docker-контейнеры** с помощью Docker Compose.
- Автоматизировал сборку и запуск через **Makefile** (цели: `build`, `run`, `migrate`).
- Настроил переменные окружения для разных сред (dev, prod).

## Для запуска приложения:

```
make build && make run
```

### Если приложение запускается впервые, необходимо применить миграции к базе данных:

```
make migrate
```