# Best-practice Go development Dockerfile for devcontainer
FROM golang:1.23-alpine

# Install common Go tools and dependencies
RUN apk add --no-cache git bash build-base curl

WORKDIR /workspace

# Optionally, install goimports and delve for debugging and formatting
RUN go install golang.org/x/tools/cmd/goimports@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest

# Set PATH for Go binaries
ENV PATH="/go/bin:$PATH"

# Use a non-root user for security (optional, recommended)
# RUN adduser -D appuser && chown -R appuser /workspace
# USER appuser

# Entrypoint is handled by devcontainer.json