# Dockerfile - chỉ dùng khi build image thật sự
FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download
RUN go build -o bds-square ./cmd/server

FROM alpine

WORKDIR /app
RUN apk add --no-cache bash curl

COPY --from=builder /build/bds-square .
COPY ./config ./config
COPY wait-for-it.sh .

RUN chmod +x wait-for-it.sh

ENTRYPOINT ["./wait-for-it.sh", "mysql:3306", "--", "./bds-square"]
