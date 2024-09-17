# # Stage 1: Build stage
# FROM golang:1.23-alpine AS build

# # Opretter en mappe i containeren
# WORKDIR /app

# # Copy go.mod and go.sum to download dependencies
# COPY go.mod go.sum ./

# # Download afhængigheder
# RUN go mod download

# # Kopier hele projektet til containeren
# COPY . .

# # Byg applikationen
# RUN go build -o main ./cmd/main


# # Stage 2: Runtime stage (mindre image)
# FROM alpine:latest

# # Opret et nyt arbejdsdirectory
# WORKDIR /app

# # Kopier den binære fil fra build-stagen
# COPY --from=build /app/main .

# # Expose porten applikationen vil køre på
# EXPOSE 9090

# # Definér entrypoint kommandoen for at køre applikationen
# CMD ["./main"]


FROM alpine:latest
CMD ["echo", "Hello, World!"]
