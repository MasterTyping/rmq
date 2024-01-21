# Use the Ubuntu base image
FROM ubuntu:latest

# Install Go
RUN apt-get update && apt-get install -y golang

# Set the working directory
WORKDIR /src

# Copy the Go application to the container
COPY . .

# Build the Go application
RUN go build -o rmq-server

# Set the entrypoint command
CMD ["./rmq-server"]
