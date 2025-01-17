# Start from the official Go image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o pplx-bot .

# Use a .dockerignore file to exclude unnecessary files
COPY .dockerignore .

# Run the application
CMD ["./pplx-bot"]