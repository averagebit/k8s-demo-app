# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

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
