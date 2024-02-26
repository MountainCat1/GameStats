# Start from the official Go image to build our application.
FROM golang:1.21 as builder

# Set the Current Working Directory inside the container.
WORKDIR /app

# Assuming you have multiple modules, you need to copy each go.mod and go.sum file.
# For example, if you have a module in the 'GameStats.Api' directory, you would do something like this:
COPY GameStats.Api/go.mod               GameStats.Api/go.sum            ./GameStats.Api/
COPY GameStats.Domain/go.mod                                            ./GameStats.Domain/
COPY GameStats.Infrastructure/go.mod    GameStats.Infrastructure/go.sum ./GameStats.Infrastructure/

# Navigate to each module's directory and download dependencies. Repeat this for each module.
WORKDIR /app/GameStats.Api
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container.
WORKDIR /app/
COPY . .

WORKDIR /app/GameStats.Api/cmd/api

RUN go mod tidy

# Build the Go app.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp /app/GameStats.Api/cmd/api

# Start a new stage from scratch.
FROM alpine:latest

# Install ca-certificates in case your application makes external HTTPS calls.
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container.
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /app/GameStats.Api/cmd/api/myapp .

# Copy config to the working
COPY ./GameStats.Api/config ./config

# Expose port 8080 to the outside world.
EXPOSE 8080

ENV ENVIRONMENT=docker

# Command to run the executable.
CMD ["./myapp"]
