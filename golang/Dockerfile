# Use official golang image
FROM golang:1.19-alpine

# Install git for go get
RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port
EXPOSE 6942

# Command to run the executable
CMD ["./main"]
