# Stage 1: Build the application
FROM golang:1.23 AS builder

WORKDIR /app

COPY main.go .

RUN CGO_ENABLED=0 go build -o hello main.go

# Stage 2: Create the minimal production image
FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/hello .

EXPOSE 8080

CMD ["./hello"]
