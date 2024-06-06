# Use the official Golang image as the base image for building
FROM golang:1.22.3-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the rest of the application source code
COPY . .

# Download dependencies
RUN go mod download


# Expose the port on which your application will run
EXPOSE 8080

# Command to run your application
CMD ["go","run" ,"main.go"]
