# Use the official Golang image as a base image
FROM golang:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the source in the container
COPY . .

# Compile the Go project
RUN go build -o main

# Expose the port on which the project will listen
EXPOSE 8080
