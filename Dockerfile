# Используем многоступенчатую сборку с Alpine
FROM golang:1.22.2-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags="-w -s" -o todo-app ./cmd/main.go

FROM alpine:3.19

RUN apk add --no-cache \
    postgresql15-client \
    tzdata

WORKDIR /app

COPY --from=builder /app/todo-app .
COPY --from=builder /app/wait-for-postgres.sh .
COPY --from=builder /app/configs ./configs
COPY --from=builder /app/.env . 
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

RUN chmod +x wait-for-postgres.sh

CMD ["./wait-for-postgres.sh", "db", "5432", "./todo-app"]