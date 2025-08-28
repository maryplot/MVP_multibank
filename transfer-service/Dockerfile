FROM golang:1.25-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /transfer-service ./cmd

EXPOSE 8082

CMD ["/transfer-service"]