# Best-practice Go development Dockerfile for devcontainer
FROM golang:1.24

# Install common Go tools and dependencies
RUN apt-get update && apt-get install -y git bash build-essential curl \
    && rm -rf /var/lib/apt/lists/*

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