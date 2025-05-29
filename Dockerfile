FROM golang:1.24.3-alpine

WORKDIR /app

RUN apk add --no-cache bash curl
RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN cp /go/bin/dlv /usr/local/bin/

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -gcflags="all=-N -l" -o app ./cmd/server/main.go
EXPOSE 8002 2345
