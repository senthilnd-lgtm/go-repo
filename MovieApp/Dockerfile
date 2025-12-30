# Stage 1: Build the Go binary
FROM golang:1.25-alpine AS builder
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy go.mod and go.sum first (for dependency caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .


# Build the binary statically
RUN CGO_ENABLED=0 GOOS=linux go build -o movieapp .

# Stage 2: Create a minimal image
FROM scratch

# copy the compiled binary 
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/movieapp /movieapp

# expose port

EXPOSE 8080

#Run the binary
ENTRYPOINT ["/movieapp"]
