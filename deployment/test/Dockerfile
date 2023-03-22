FROM golang:1.20.2-alpine3.17 as builder

RUN apk update
RUN apk add build-base

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

RUN go install github.com/golang/mock/mockgen@v1.6.0

COPY . .
