FROM golang:1.25-alpine

WORKDIR /app

# Копируем зависимости сначала для кэширования
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Сборка
RUN go build -o /auth-service ./cmd

EXPOSE 8080

# Запускаем сервис ( .env загружается автоматически через os.Getenv в коде)
CMD ["/auth-service"]