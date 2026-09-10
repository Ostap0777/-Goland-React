FROM golang:1.24-alpine

WORKDIR /app

# Обов'язково додаємо GOTOOLCHAIN=auto, щоб Go сам підтягував потрібний інструментарій при встановленні пакета
ENV GOTOOLCHAIN=auto

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air"]