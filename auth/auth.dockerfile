# Stage 1: Build the binary
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy dependency files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Build the executable
RUN CGO_ENABLED=0 GOOS=linux go build -o auth main.go

# Stage 2: Run the binary in a lightweight image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/auth .

# Expose the TCP RPC port used by the broker service
EXPOSE 5001

CMD ["./auth"]