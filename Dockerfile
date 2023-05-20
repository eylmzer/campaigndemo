# Use the official Golang image as the base image
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

COPY cmd/commands commands/

COPY . .

# Build the Go application
RUN go build -o campaingdemo ./cmd

# Set the command to run the binary executable
CMD ["./campaingdemo"]
