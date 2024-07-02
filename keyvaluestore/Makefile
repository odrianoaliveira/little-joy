# Variables
APP_NAME := keybaluestore
SRC_DIRS := $(shell go list -f '{{.Dir}}' ./...)

# Default target
.PHONY: all
all: build

# Build the project
.PHONY: build
build:
	@echo "Building the project..."
	go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)

# Run tests with verbose output
.PHONY: test
test:
	@echo "Running tests..."
	go test -v ./...

# Format the code
.PHONY: fmt
fmt:
	@echo "Formatting the code..."
	go fmt $(SRC_DIRS)

# Run go vet
.PHONY: vet
vet:
	@echo "Running go vet..."
	go vet $(SRC_DIRS)

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -f bin/$(APP_NAME)

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	go mod tidy

# Run the application
.PHONY: run
run: build
	@echo "Running the application..."
	./bin/$(APP_NAME)

# Help
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  all     - Default target: build the project"
	@echo "  build   - Build the project"
	@echo "  test    - Run tests"
	@echo "  fmt     - Format the code"
	@echo "  vet     - Run go vet"
	@echo "  clean   - Clean build artifacts"
	@echo "  deps    - Install dependencies"
	@echo "  run     - Run the application"
	@echo "  help    - Show this help message"
