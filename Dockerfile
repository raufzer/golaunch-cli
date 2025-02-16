# ---- Build Stage ----
    FROM golang:1.23-alpine AS builder

    # 1. Install build dependencies and setup environment in one layer
    RUN apk add --no-cache git && \
        addgroup -S app && \
        adduser -S -G app app && \
        mkdir -p /app /go-mod-cache && \
        chown -R app:app /app /go-mod-cache
    
    # 2. Set Go module environment variables
    ENV GOMODCACHE=/go-mod-cache \
        GOPATH=/go \
        CGO_ENABLED=0
    
    # 3. Switch to app user
    USER app
    WORKDIR /app
    
    # 4. Copy dependency files
    COPY --chown=app:app go.mod go.sum ./
    
    # 5. Download dependencies
    RUN go mod download
    
    # 6. Copy source code
    COPY --chown=app:app . .
    
    # 7. Build optimized binary
    RUN go build \
        -ldflags="-s -w" \
        -trimpath \
        -o ./golaunch \
        ./cmd/golaunch
    
    # ---- Final Stage ----
    FROM alpine:3.21
    
    # 1. Install runtime dependencies in one layer
    RUN apk --no-cache add ca-certificates && \
        addgroup -S app && \
        adduser -S -G app app && \
        mkdir -p /app/assets && \
        chown -R app:app /app
    
    # 2. Switch to app user
    USER app
    WORKDIR /app
    
    # 3. Copy built artifacts
    COPY --from=builder --chown=app:app /app/golaunch .
    COPY --from=builder --chown=app:app /app/assets ./assets
    
    # 4. Add metadata
    LABEL maintainer="Your Name <your.email@example.com>"
    LABEL org.opencontainers.image.source="https://github.com/raufzer/golaunch-cli"
    
    # 5. Set entrypoint
    ENTRYPOINT ["./golaunch"]