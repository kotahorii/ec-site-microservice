FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o gateway .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/gateway .

CMD ["./gateway"]
