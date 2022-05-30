# ------ DEVELOPMENT BLOCK ------ #
FROM golang:1.18 as development

# Add a work directory
WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy app files
COPY . .

# Install Reflex for development -> Enables live reload
RUN go install github.com/cespare/reflex@latest

# Start app
CMD reflex -g '*.go' go run main.go --start-service

# ------ BUILDER BLOCK ------ #
FROM golang:1.18 as builder

# Define build env
ENV GOOS linux
ENV CGO_ENABLED 0

# Add a work directory
WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build app
RUN go build main.go

# ------ PRODUCTION BLOCK ------ #
FROM alpine:latest as production

# Add certificates
RUN apk add --no-cache ca-certificates

# Copy built binary from builder
COPY --from=builder /app/main .

# Exec built binary
CMD ./main
