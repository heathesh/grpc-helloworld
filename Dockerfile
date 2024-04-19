# Use the official Go image as the base image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o app

# Set the command to run the executable
CMD ["./app"]