FROM golang:1.18 AS builder
WORKDIR /app

# Download dependencies (unless cached).
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o phonebook

# Use the official Debian-based image as a base image.
FROM debian:bullseye-slim

# Copy the binary from the builder stage.
COPY --from=builder /app/phonebook /phonebook

EXPOSE 8080
CMD ["/phonebook"]
