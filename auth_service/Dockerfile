# Используем официальное изображение Golang
FROM golang:1.24-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем зависимости и загружаем модули
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/auth_service ./cmd

# Финальный образ
FROM alpine:3.17

# Устанавливаем сертификаты
RUN apk --no-cache add ca-certificates

# Копируем бинарный файл из builder образа
COPY --from=builder /go/bin/auth_service /usr/local/bin/auth_service

# Копируем .env внутрь контейнера
COPY .env /app/.env

# Указываем рабочую директорию
WORKDIR /app

# Запускаем приложение
ENTRYPOINT ["auth_service"]
