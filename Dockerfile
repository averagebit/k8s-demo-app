# ------ DEVELOPMENT BLOCK ------ #
FROM golang:1.18-alpine as development

# Switch working directory
WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum /app/
RUN go mod download

# Copy app files
COPY . /app

# Install Reflex for development -> Enables live reload
RUN go install github.com/cespare/reflex@latest

# Create non-root user and group
RUN addgroup -S -g 1050 app && adduser -S -u 1050 -G app app

# Change ownership of /app to user app
RUN chown -R app:app /app

# Switch to non-root user
USER app

# Start app
CMD ["reflex", "-g", "'*.go'", "go run", "main.go", "--start-service"]

# ------ BUILDER BLOCK ------ #
FROM golang:1.18-alpine as builder

# Define build env
ENV GOOS linux
ENV CGO_ENABLED 0

# Switch working directory
WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum /app/
RUN go mod download

COPY . /app

# Build app
RUN ["go", "build", "main.go"]

# ------ PRODUCTION BLOCK ------ #
FROM alpine:3.16.0 as production

# Install dependencies
RUN apk add --no-cache ca-certificates

# Switch working directory
WORKDIR /app

# Create non-root user and group
RUN addgroup -S -g 1050 app && adduser -S -u 1050 -G app app

# Copy built binary from builder
COPY --from=builder /app/main /app/

# Change ownership of /app to user app
RUN chown -R app:app /app

# Switch to non-root user
USER app

# Exec built binary
CMD ["./main"]
