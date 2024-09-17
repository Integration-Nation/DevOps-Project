# Use the official Golang image as the build stage
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Copy the .env file if it exists
COPY .env .env

# Build the Go app
RUN go build -o main ./cmd/main

# Start a new stage from scratch with Ubuntu
FROM ubuntu:latest

# Install necessary dependencies
RUN apt-get update && apt-get install -y ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy the .env file
COPY --from=builder /app/.env .

# Verify the binary file is copied correctly
RUN ls -l /root/

# Command to run the executable
CMD ["./main"]

# Expose the port the app runs on
EXPOSE 9090
