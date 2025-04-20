FROM golang:1.23.2-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o apigateway ./cmd/

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/apigateway  .

ENV GIN_MODE=release

CMD ["./apigateway"]
