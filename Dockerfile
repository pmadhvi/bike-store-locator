# Start from the latest golang alpine base image
FROM golang:1.14.3-alpine AS builder

#installing ca-certificates
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates openssl\
        && update-ca-certificates 2>/dev/null || true

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o bikelocator .

# Command to run the executable
CMD ["./bikelocator"]