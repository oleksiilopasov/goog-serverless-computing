# Start from the official Go image
FROM golang:1.22-alpine AS build

# Set the current working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the source code directory
COPY src/ ./src

# Build the Go application
RUN go build -o app ./src

# Start a new stage from scratch
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /root/

# Copy the executable from the previous stage
COPY --from=build /app/app .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./app"]