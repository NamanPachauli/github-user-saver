
# Use the official Golang 1.25.1 image
FROM golang:1.25.1

# Set the working directory inside the container
WORKDIR /app

# Copy all files from the current directory to the container
COPY . .

# Build the Go application
RUN go build -o server .

# Expose port 9090
EXPOSE 9090

# Command to run the application
CMD ["./server"]