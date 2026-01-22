# Multi-stage build for ego_site_htmx
# Build stage
FROM golang:1.22-alpine AS build

WORKDIR /app

# Install build tools
RUN apk add --no-cache ca-certificates git

# Go module files
COPY go.mod ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/web

# Runtime stage
FROM alpine:3.19

WORKDIR /app

# Certificates for HTTPS requests, if any
RUN apk add --no-cache ca-certificates

# Copy binary and required assets
COPY --from=build /app/server ./server
COPY --from=build /app/web ./web

# Fly.io expects the app to listen on PORT (default 8080)
ENV PORT=8080
EXPOSE 8080

# Run the server
CMD ["./server"]
