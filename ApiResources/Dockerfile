# Use the official Golang image as the base image
FROM golang:1.23.4

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Build the Go application
RUN go build -o apiresources main.go

# Expose the port the application runs on
EXPOSE 8084

# Set environment variables (if needed)
ENV DB_HOST=3.211.243.19
ENV DB_PORT=5432
ENV DB_USER=esparrago
ENV DB_PASS=CHPaladin
ENV DB_SCHEMA=access_gate

# Command to run the executable
CMD ["./apiresources"]