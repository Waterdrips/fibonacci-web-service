FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY web-service/go.mod web-service/go.sum ./

RUN go mod download

COPY web-service/ .

# Build the Go application statically
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/main /main

# Set the entrypoint to the binary
ENTRYPOINT ["/main"]