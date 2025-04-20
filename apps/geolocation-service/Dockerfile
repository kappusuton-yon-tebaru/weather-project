FROM golang:1.23.2-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o geolocation ./cmd/geolocation/

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/geolocation  .

ENV GIN_MODE=release

CMD ["./geolocation"]
