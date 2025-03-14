# Начальная стадия: загрузка модулей
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Копируем файлы проекта
COPY . .

# Загружаем зависимости
RUN go mod download

# Компиляция Go-приложения
RUN go build -o main ./cmd

# Финальная стадия: запуск приложения
FROM alpine:latest AS runner

# Копируем скомпилированное приложение
COPY --from=builder /app/main /app/main

WORKDIR /app

# Устанавливаем переменные окружения
ENV CONFIG_PATH=./config/production.yaml

# Ожидаем пока БД будет готова и выполняем миграции, затем запускаем приложение
CMD ["/app/main"]

EXPOSE 8080