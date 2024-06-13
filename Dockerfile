# Build stage
FROM golang:1.22 as builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the worker binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /worker ./cmd/worker

# Build the server binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server ./cmd/server

# Final stage
FROM alpine:latest

# Set up the working directory
WORKDIR /root/

# Copy the worker binary from the build stage
COPY --from=builder /worker .

# Copy the server binary from the build stage
COPY --from=builder /server .

# Expose the server port
EXPOSE 8080

# Provide a start script to choose between server and worker
COPY start.sh .

# Make the script executable
RUN chmod +x start.sh

# Run the start script
ENTRYPOINT ["./start.sh"]
