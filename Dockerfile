# Use the official Golang image to build btcd
FROM golang:latest AS builder

# Install necessary dependencies
RUN apt-get update && apt-get install -y git

# Clone btcd repository
RUN git clone https://github.com/btcsuite/btcd.git /go/src/github.com/btcsuite/btcd

# Build btcd
RUN cd /go/src/github.com/btcsuite/btcd && go install . ./cmd/...

# Use a lightweight base image
FROM alpine:latest

# Copy btcd binary from builder stage
COPY --from=builder /go/bin/btcd /usr/local/bin/btcd

# Expose necessary ports (8333 for mainnet, 18333 for testnet)
EXPOSE 8333 18333 9333

# Run btcd with default configuration
CMD ["btcd"]
