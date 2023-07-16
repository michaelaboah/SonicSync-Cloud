# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

RUN mkdir -p /app/logs
# Copy the Go module files and download dependencies
COPY /workspace/go.mod /workspace/go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY ./workspace .

# Build the Go application
RUN go build -o server ./cmd/server.go

# Expose the port on which the server will listen
EXPOSE 8080

# Set the command to run the server executable
CMD ["./server"]
