# Dockerfile.dev
FROM golang:1.24.3-alpine

WORKDIR /app

RUN apk add --no-cache bash curl
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8002
 