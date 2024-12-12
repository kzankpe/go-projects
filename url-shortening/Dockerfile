# Use the golang official image
FROM golang:1.20 AS builder

# Set Working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source code into the container
COPY ./config /config
COPY ./docs /docs
COPY ./handlers /handlers
COPY ./models /models
COPY ./routes /routes
COPY ./main.go /


# Build the Go app
RUN go build -o main .

# Start a new stage from scratch using a specific image version
FROM alpine:3.21 

# Create a non-root user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Change ownership of the application to the non-root user
RUN chown appuser:appgroup main

# Switch to the non-root user
USER appuser

# Expose port 8090 to the outside world
EXPOSE 8090

# Command to run the executable
CMD ["./main"]