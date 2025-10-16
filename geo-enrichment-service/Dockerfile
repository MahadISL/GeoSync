# --- Stage 1: The Builder ---
# Use the official Go image as a parent image.
# The 'alpine' tag means it's a lightweight version of the image.
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container.
WORKDIR /app

# Copy the Go module files and download dependencies.
# This is done in a separate step to leverage Docker's layer caching.
# If go.mod and go.sum don't change, Docker will reuse this cached layer.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code into the container.
COPY . .

# Build the Go application.
# -o /app/main specifies the output file path.
# -ldflags="-s -w" strips debugging information, making the binary smaller.
# CGO_ENABLED=0 disables CGO, which is needed for a static binary.
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main -ldflags="-s -w" .

# --- Stage 2: The Final Image (Debug-Friendly) ---
FROM alpine:latest

# Add curl and CA certificates to handle HTTPS requests
RUN apk add --no-cache ca-certificates curl

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose port
EXPOSE 8081

# Run the binary
ENTRYPOINT ["/app/main"]