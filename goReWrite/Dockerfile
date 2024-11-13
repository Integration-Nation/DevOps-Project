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

# Set the environment variables from the build arguments
ARG DATABASE_PATH
ARG JWT_SECRET
ARG POSTGRES_DB_URL

# # Use the build arguments to set environment variables inside the container
ENV DATABASE_PATH=${DATABASE_PATH}
ENV JWT_SECRET=${JWT_SECRET}
ENV POSTGRES_DB_URL=${POSTGRES_DB_URL}

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

# Set environment variables in the final stage
ARG DATABASE_PATH
ARG JWT_SECRET
ARG POSTGRES_DB_URL
ENV DATABASE_PATH=${DATABASE_PATH}
ENV JWT_SECRET=${JWT_SECRET}
ENV POSTGRES_DB_URL=${POSTGRES_DB_URL}

# Command to run the executable
CMD ["./main"]

# Expose the port the app runs on
EXPOSE 9090
