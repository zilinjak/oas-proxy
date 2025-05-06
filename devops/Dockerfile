# Build stage
FROM golang:1.24.1-alpine AS builder
LABEL authors="jakubzilinek"

# Install dependencies (ca-certificates for SSL, git for go mod)
RUN apk add --no-cache ca-certificates git

# Set working directory
WORKDIR /app

# Copy go mod files first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
# -ldflags="-s -w" reduces binary size by stripping debug info
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app cmd/app/main.go

# Final stage
FROM alpine:latest

# Install timezone data if your app needs it
# RUN apk add --no-cache tzdata

# Copy SSL certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy binary from builder
COPY --from=builder /app/app /app/

# Set working directory
WORKDIR /app

# Expose the port your app runs on (change if needed)
EXPOSE 8080
ENV GIN_MODE=release
# Command to run the application
CMD ["./app"]