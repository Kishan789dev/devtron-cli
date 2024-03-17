# Use alpine as the base image
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Install Go and other necessary tools
RUN apk --no-cache add go git

# Copy the source code into the container
COPY . .

# Download dependencies
RUN go mod download

# Build the Go app
RUN go build -o mycli

# Set permissions for the executable
RUN chmod +x mycli

# Expose any necessary ports
EXPOSE 8080

# Command to run the executable
CMD ["./mycli"]
