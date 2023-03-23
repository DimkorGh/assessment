FROM golang:1.20.2-alpine3.17 as builder

RUN apk update
RUN apk add build-base

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main

FROM alpine:3.17.1

WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/config/config.yaml /app/config/

EXPOSE 8080

CMD ["./main", "-p", "8080", "-addr", "0.0.0.0"]