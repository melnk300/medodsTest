FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o ./main ./cmd/app/main.go


FROM alpine:latest
LABEL authors="ruslanmelnik"

WORKDIR /root/

COPY --from=builder /app/main .
COPY ./.env .

RUN chmod +x /root/main

EXPOSE 3000

CMD ["./main"]