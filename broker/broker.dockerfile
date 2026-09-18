# --- Build Stage ---
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o broker ./cmd/api

# --- Run Stage ---
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/broker .

EXPOSE 8080

CMD ["./broker"]