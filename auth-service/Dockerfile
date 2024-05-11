# Use a lightweight Alpine image
FROM golang:latest as build_container

# Set the working directory
WORKDIR /app

# Copy Go modules files
COPY ./auth-service/go.mod ./auth-service/go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the application source code
COPY ./auth-service/ .

# Build the application
RUN go build -o server ./main.go

# Use a minimal Alpine base image
FROM alpine:latest

# Set the working directory
WORKDIR /usr/bin

# Copy the built binary from the previous stage
COPY --from=build_container /app/server .

# Expose the port on which the application will run
EXPOSE 8000

# Set the entry point for the container
ENTRYPOINT ["./server"]
