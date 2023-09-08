#FROM golang:1.19 AS builder
#
#WORKDIR /app
#COPY . .
#
#RUN go mod download
#RUN go build -o main cmd/main.go
#
#FROM alpine:latest
#WORKDIR /root
#COPY --from=builder /app .
#
#CMD ["./cmd/main.go"]

FROM golang:1.18-alpine AS builder
# Set the Current Working Directory inside the container
WORKDIR /app

RUN export GO111MODULE=on

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY . .

# Build the application
RUN go build -o main ./cmd/main.go

# Expose port 9000 to the outside world
EXPOSE 8888

# Command to run the executable
CMD ["./main"]
