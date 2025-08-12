FROM golang:1.24.2-alpine AS builder

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .
ENV GIN_MODE=release
# Build (static compilation recommended for containers)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main cmd/bot/main.go

# Runtime image
FROM alpine
WORKDIR /app

COPY --from=builder /app/main .
COPY ./config/ /app/config/

ENV GIN_MODE=release
CMD ["./main -config='config/local.yaml' -env='config/.env'"]