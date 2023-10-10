# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go app
RUN go build -o main ./main/main.go

# Change file permissions
RUN chmod +x /app/main

# Run the Go app
CMD ["/app/main"]
