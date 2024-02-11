# STAGE 1
FROM golang:alpine as builder

LABEL maintainer="humanbelnik"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ggrep ./cmd/cli/

# STAGE 2
FROM alpine:latest

WORKDIR /root

COPY --from=builder ./app/ggrep .
