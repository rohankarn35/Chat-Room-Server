# Environment setup
export GOPATH=$(HOME)/go
export PATH := $(GOPATH)/bin:$(PATH)

# Variables
BINARY_NAME := websocket-golang

# Default target
all: serve

# Build target
build:
	go build -o $(BINARY_NAME) ./cmd/server

# Run the application
serve:
	gow run cmd/main.go

# Clean target
clean:
	rm -f $(BINARY_NAME)

# Manage dependencies
deps:
	go mod tidy

.PHONY: all build serve clean deps
