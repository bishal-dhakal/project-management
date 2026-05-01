APP_NAME := projectmanagement
BIN_DIR := bin

BIN_PATH := $(BIN_DIR)/$(APP_NAME)

.PHONY: build run dev clean help

## build: Build the Go binary
build:
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_PATH) ./cmd

## run: Run the built binary
run: build
	@./$(BIN_PATH)

## dev: Run with live reload using air
dev:
	@air

## clean: Remove build artifacts
clean:
	@rm -rf $(BIN_DIR)

## help: Show available commands
help:
	@echo ""
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "  %-10s %s\n", $$1, $$2}'