# Этап сборки
FROM golang:1.22.2-bookworm AS builder

# Система модулей
ENV GO111MODULE=on
WORKDIR /app

COPY .env .  
# Копируем зависимости
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o todo-app ./cmd/main.go

# Финальный образ
FROM debian:bookworm-slim

RUN apt-get update && \
    apt-get install -y postgresql-client && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Копируем артефакты
COPY --from=builder /app/.env /app/.env  
COPY --from=builder /app/todo-app /app/todo-app
COPY --from=builder /app/wait-for-postgres.sh /app/
COPY --from=builder /app/configs /app/configs

RUN chmod +x /app/wait-for-postgres.sh

WORKDIR /app

CMD ["./wait-for-postgres.sh", "db:5432", "--", "./todo-app"]