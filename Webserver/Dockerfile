# Start from the official Go image
FROM golang:1.25 AS builder

#RUN adduser -D builduser

WORKDIR /app
COPY . .

# Build the Go binary
RUN go build -o webapi .


 FROM alpine:3.19

 WORKDIR /app

 COPY --from=builder /app/webapi .

# USER 1000

EXPOSE 8080

ENTRYPOINT ["/app/webapi"]