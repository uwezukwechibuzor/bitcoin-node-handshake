# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=btchandshake

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	go test ./...	

install:
	$(GOBUILD) -o $(BINARY_NAME) -v
	@echo "Installing $(BINARY_NAME) to GOPATH..."
	@cp $(BINARY_NAME) $(GOPATH)/bin/$(BINARY_NAME)

clean:
	$(GOCMD) clean
	rm -f $(BINARY_NAME)

.PHONY: all build install clean



