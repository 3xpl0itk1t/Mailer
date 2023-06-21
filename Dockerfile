# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code from the host to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose a port (if your application listens on a specific port)
EXPOSE 8000

# Define the command to run your application
CMD ["./main"]
