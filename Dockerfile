# syntax=docker/dockerfile:1

FROM golang:1.24.1 AS builder

WORKDIR /app

# Copy Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the app
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Final image
FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=builder /app/server .
ENTRYPOINT ["/app/server"]