FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

RUN go mod download

COPY ./ ./

CMD ["go", "run", "."]
