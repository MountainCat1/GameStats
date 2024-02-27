# Start from the official Go image to build our application.
FROM golang:1.21 as builder

# Set the Current Working Directory inside the container.
WORKDIR /app

# Assuming you have multiple modules, you need to copy each go.mod and go.sum file.
# For example, if you have a module in the 'GameStats.Api' directory, you would do something like this:
COPY gamestats-application/go.mod       gamestats-application/go.sum    ./gamestats-application/
COPY gamestats-domain/go.mod                                            ./gamestats-domain/
COPY gamestats-infrastructure/go.mod    gamestats-infrastructure/go.sum ./gamestats-infrastructure/

# Navigate to each module's directory and download dependencies. Repeat this for each module.
WORKDIR /app/gamestats-application
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container.
WORKDIR /app/
COPY . .

WORKDIR /app/gamestats-application/cmd/api

RUN go mod tidy

# Build the Go app.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp /app/gamestats-application/cmd/api

# Start a new stage from scratch.
FROM alpine:latest

# Install ca-certificates in case your application makes external HTTPS calls.
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container.
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /app/gamestats-application/cmd/api/myapp .

# Copy config to the working
COPY ./gamestats-application/config ./config

# Expose port 8080 to the outside world.
EXPOSE 8080

ENV ENVIRONMENT=docker

# Command to run the executable.
CMD ["./myapp"]
