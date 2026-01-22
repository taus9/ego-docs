# Multi-stage build for ego_site_htmx

# Build stage
FROM golang:1.25.6-alpine AS build

WORKDIR /app

# Install build tools (git needed if you fetch modules from VCS)
RUN apk add --no-cache ca-certificates git

# Go module files (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/web


# Runtime stage
FROM alpine:3.23

WORKDIR /app

# Certificates for HTTPS requests, if any
RUN apk add --no-cache ca-certificates

# Create non-root user
RUN addgroup -S app && adduser -S app -G app

# Copy binary and required assets
COPY --from=build --chown=app:app /app/server ./server
COPY --from=build --chown=app:app /app/web ./web

USER app

# Fly.io expects the app to listen on PORT (default 8080)
ENV PORT=8080
EXPOSE 8080

CMD ["./server"]
