FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go build -o app cmd/api/main.go

EXPOSE 8080

CMD ["./app"]