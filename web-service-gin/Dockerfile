# Use an official Golang runtime as a base image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Download and install any required third-party dependencies
RUN go mod download

# Build the application
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

