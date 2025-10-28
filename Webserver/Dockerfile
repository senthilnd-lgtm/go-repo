# Start from the official Go image
FROM golang:1.25

WORKDIR /app
COPY . .

# Build the Go binary
RUN go build -o webapi .


# Run the binary
EXPOSE 8080
CMD ["./webapi"]
