# Use the official Golang image as the base image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY motd-api.go .

# Build the Go application
RUN go build -o motd-api .

# Expose the port that the application will run on
EXPOSE 8080

# Set environment variables for the application
ENV MESSAGE="Default Message"
ENV APP_PORT=8080

# Command to run the application
CMD ["./motd-api"]
