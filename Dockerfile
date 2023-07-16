# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY ./workspace .

# Build the Go application
RUN go build -o server .

# Expose the port on which the server will listen
EXPOSE 8080

# Set the command to run the server executable
CMD ["./server"]
