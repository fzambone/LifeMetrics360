# Use the official Go image as a base image
FROM golang:1.21.1-alpine

# Set the current working directory in the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the application
CMD ["go", "run", "./cmd/main.go"]