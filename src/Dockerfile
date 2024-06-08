# Start from the official Go image
FROM golang:1.22-alpine AS build
# Copy the source code directory
COPY . /src
# Set the current working directory inside the container
WORKDIR /src
# Download dependencies
RUN go mod download
# Build the Go application
RUN go build -o /app/smarty

# Start a new stage from scratch
FROM alpine:latest
# Set the current working directory inside the container
WORKDIR /app
# Copy the executable from the previous stage
COPY --from=build /app/smarty .
# Expose port 8080
EXPOSE 8080
# Command to run the executable
CMD ["./smarty"]