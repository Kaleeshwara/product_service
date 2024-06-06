# Use the official Golang image as the base image for building
FROM golang:1.22.3-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and build files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/main .

# Copy the .env file from the builder stage
COPY --from=builder /app/.env .

# Expose the port on which your application will run
EXPOSE 8080

# Command to run your application
CMD ./main
