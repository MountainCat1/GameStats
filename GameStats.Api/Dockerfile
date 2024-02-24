# Start from the official Go image to build our application.
FROM golang:1.18 as builder

# Set the Current Working Directory inside the container.
WORKDIR /app

# Copy go mod and sum files.
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container.
COPY . .

# Build the Go app.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp ./cmd/api

# Start a new stage from scratch.
FROM alpine:latest

# Install ca-certificates in case your application makes external HTTPS calls.
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container.
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /app/myapp .

# Copy config to the working
COPY ./config ./config

# Expose port 8080 to the outside world.
EXPOSE 8080

ENV ENVIRONMENT=docker

# Command to run the executable.
CMD ["./myapp"]
