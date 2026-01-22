# Multi-stage build for ego_site_htmx

# Build stage
FROM golang:1.25.6-alpine AS build

WORKDIR /app

# Install build tools (git needed if you fetch modules from VCS)
RUN apk add --no-cache ca-certificates git

# Copy go.mod first (go.sum may not exist and that's OK)
COPY go.mod ./

# Copy the rest of the source
COPY . .

# Download deps (safe even if there are none)
RUN go mod download

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o server ./cmd/web


# Runtime stage
FROM alpine:3.23

WORKDIR /app

RUN apk add --no-cache ca-certificates

# Create non-root user
RUN addgroup -S app && adduser -S app -G app

COPY --from=build --chown=app:app /app/server ./server
COPY --from=build --chown=app:app /app/web ./web

USER app

ENV PORT=8080
EXPOSE 8080

CMD ["./server"]
