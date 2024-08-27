
# Use the official Golang image as a base
FROM golang:1.22.6-alpine

# Set the current working directory inside the container
WORKDIR /go/src/GinTest

# Copy go.mod and go.sum to the container
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the entire project source code to the container
COPY . .

# Set the working directory to where your main.go file is
WORKDIR /go/src/GinTest/cmd/server

# Build the Go application
RUN go build -o /go/bin/main .

# Expose the port the application will run on
EXPOSE 8080

# Command to run the application
CMD ["/go/bin/main"]


