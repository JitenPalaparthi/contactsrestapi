# Use an official Golang runtime as a parent image
FROM golang:1.21-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o service .

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/service .

# Expose port 8080 to the outside world
RUN mkdir uploads

# ARG PORT

# EXPOSE ${PORT}

# Command to run the executable
ENTRYPOINT ["./service","--mode","postgres"]
