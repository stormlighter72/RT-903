# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container's working directory
COPY . .

# Build the Go application
RUN go build -o motd-api .

# Expose the port on which the application will run
EXPOSE 5555

# Set environment variables
ENV APP_PORT=5555
ENV MESSAGE="Hello World!"

# Command to run the executable
CMD ["./motd-api"]
