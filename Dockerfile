# Stage 1: Builder
FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

# Copy dependencies files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 go build -o main main.go

# Stage 2: Runtime
FROM debian:bullseye-slim

# Set the working directory
WORKDIR /app

# Copy only the built binary from the builder stage
COPY --from=builder /app/main .

# Expose the application's port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
