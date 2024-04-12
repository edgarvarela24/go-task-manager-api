# Use the official Go image as the base image
FROM golang:1.22.1

# Install Delve debugger
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Set the working directory inside the container
WORKDIR /go/src/github.com/edgarvarela24/task-manager-api

# Copy the Go module files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code into the container
COPY . .

# Build the Go application with debugging information. Flags disable optimizations and inlining
RUN go build -gcflags "all=-N -l" -o main ./cmd/server

# Expose the application port and the debugging port
EXPOSE 8080 2345

# Set the entry point for the container
CMD ["dlv", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/go/src/github.com/edgarvarela24/task-manager-api/main"]