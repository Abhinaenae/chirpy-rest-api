# Step 1: Use Go builder image to build the application
FROM golang:1.22 AS builder

# Set working directory inside the container
WORKDIR /app

# Copy Go files and dependencies
COPY . .

# Download the dependencies
RUN go mod download

# Build the Go application inside the container (this avoids platform incompatibility)
RUN go build -o chirpy .

# Step 2: Create the final container with a minimal Alpine image
FROM alpine:latest

# Install necessary libraries (glibc or others that Go may depend on)
RUN apk update && apk add --no-cache libc6-compat

# Set the working directory inside the container
WORKDIR /app

# Copy the .env file into the container
COPY .env /app/.env

# Set environment variables
ENV $(cat /app/.env | xargs)

# Copy the compiled Go executable from the builder stage
COPY --from=builder /app/chirpy .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./chirpy"]

