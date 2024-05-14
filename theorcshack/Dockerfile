# Use the official Golang image as a builder
FROM golang:1.20-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o main .

# Use a smaller base image for the final build
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy the .env file
COPY --from=builder /app/.env .

# Set environment variables for the database connection
ENV POSTGRES_USER=dancingponysvc
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=dancingpony
ENV POSTGRES_HOST=postgres
ENV POSTGRES_PORT=5432

# Expose the application port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
