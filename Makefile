# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=btchandshake

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	go test ./...	

lint:
	golangci-lint run 
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "*_test.go" | xargs gofmt -d -s

format:
	@go install mvdan.cc/gofumpt@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name "*.pb.go" -not -name "*.pb.gw.go" -not -name "*.pulsar.go" -not -path "./crypto/keys/secp256k1/*" | xargs gofumpt -w -l
	golangci-lint run --fix
.PHONY: format	

install:
	$(GOBUILD) -o $(BINARY_NAME) -v
	@echo "Installing $(BINARY_NAME) to GOPATH..."
	@cp $(BINARY_NAME) $(GOPATH)/bin/$(BINARY_NAME)

clean:
	$(GOCMD) clean
	rm -f $(BINARY_NAME)

.PHONY: all build install clean



