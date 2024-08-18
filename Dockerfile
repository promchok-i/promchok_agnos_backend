# Step 1: Build the Go API
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o gin-api

# Step 2: Run the tests in the container
FROM builder AS run-test-stage
RUN go test -v ./...

# Step 3: Create a small image for the Gin API
FROM alpine:latest

WORKDIR /root/
COPY --from=run-test-stage /app/gin-api .

EXPOSE 8080
CMD ["./gin-api"]
