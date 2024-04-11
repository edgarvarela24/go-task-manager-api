# Use the official Go image as the base image
FROM golang:1.22.1

# Set the working directory inside the container
WORKDIR /go/src/github.com/edgarvarela24/task-manager-api

# Copy the Go module files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code into the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd/server

# Expose the port on which your application will run
EXPOSE 8080

# Set the entry point for the container
CMD ["./main"]