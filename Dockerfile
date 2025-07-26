# Stage 1: Build the Go application
FROM golang:1.24 AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker's build cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go binary
# CGO_ENABLED=0 creates a statically linked binary, improving portability
# GOOS=linux ensures the binary is built for Linux, which is common for containers
RUN CGO_ENABLED=0 GOOS=linux go build -o auth-logs .

# Stage 2: Create a minimal production image
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/auth-logs .

# Expose the port your application listens on (if applicable)
EXPOSE 8080

# Command to run the application when the container starts
CMD ["./auth-logs"]