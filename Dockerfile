FROM golang:1.26.3-bookworm

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o ./library-bin /app/cmd/app/main.go

# Запускаем готовый бинарник напрямую
CMD ["./library-bin"]