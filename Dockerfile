# ---- Build Stage ----
    FROM golang:1.23-alpine AS builder

    # 1. Install build dependencies as root
    RUN apk add --no-cache git
    
    # 2. Create app user and set up directories
    RUN addgroup -S app && adduser -S -G app app \
        && mkdir -p /app \
        && chown -R app:app /app \
        && mkdir -p /go-mod-cache \
        && chown -R app:app /go-mod-cache
    
    # 3. Set Go module environment variables
    ENV GOMODCACHE=/go-mod-cache \
        GOPATH=/go
    
    # 4. Switch to app user
    USER app
    WORKDIR /app
    
    # 5. Copy dependency files
    COPY --chown=app:app go.mod go.sum ./
    
    # 6. Download dependencies
    RUN go mod download
    
    # 7. Copy source code
    COPY --chown=app:app . .
    
    # 8. Build binary
    RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-s -w" -trimpath -o ./golaunch ./cmd/golaunch/main.go

    
    # ---- Final Stage ----
    FROM alpine:3.21
    
    # 1. Install runtime dependencies (if any)
    RUN apk --no-cache add ca-certificates
    
    # 2. Create app user and set up directories
    RUN addgroup -S app && adduser -S -G app app \
        && mkdir -p /app/assets \
        && chown -R app:app /app
    
    # 3. Switch to app user
    USER app
    WORKDIR /app
    
    # 4. Copy the built binary and assets
    COPY --from=builder --chown=app:app /app/golaunch .
    COPY --from=builder --chown=app:app /app/assets ./assets
    
    # 5. Set the entry point
    ENTRYPOINT ["./golaunch"]