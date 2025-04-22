FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/crm-api ./cmd/api

# Create a minimal image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/crm-api .
COPY --from=builder /app/migrations ./migrations

# Set environment variables
ENV GIN_MODE=release

# Expose the port
EXPOSE 8080

# Run the application
CMD ["/app/crm-api"]
