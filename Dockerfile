# Start from the official Go image to build our application
FROM golang:1.21 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the application for a smaller and secure container
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o nbpCaller .

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the executable from the builder container
COPY --from=builder /app/nbpCaller .

RUN mkdir data



# Command to run the executable
CMD ["./nbpCaller"]
ENTRYPOINT ["./nbpCaller"]
