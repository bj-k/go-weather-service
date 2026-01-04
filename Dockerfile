# Build stage
FROM docker.io/golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o weather-service ./cmd/weather-service

# Final stage
FROM docker.io/alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/weather-service .

# Copy the OpenAPI spec file (required by the application)
# We need to maintain the directory structure relative to the binary or update the code.
# The code expects "api/openapi.yaml" relative to the working directory.
COPY --from=builder /app/api/openapi.yaml ./api/openapi.yaml

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./weather-service"]
