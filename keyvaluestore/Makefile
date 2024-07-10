# Variables
APP_NAME := api
SRC_DIRS := $(shell go list -f '{{.Dir}}' ./...)
CMD_DIR := ./cmd/$(APP_NAME)
BIN_DIR := bin

# Default target
.PHONY: all
all: build

# Build the project
.PHONY: build
build: $(BIN_DIR)
	@echo "Building the project ..."
	go build -o $(BIN_DIR)/$(APP_NAME)$(EXT) $(CMD_DIR)
	@if [ -f $(BIN_DIR)/$(APP_NAME)$(EXT) ]; then echo "Build successful"; chmod +x $(BIN_DIR)/$(APP_NAME)$(EXT); else echo "Build failed"; exit 1; fi

$(BIN_DIR):
	@mkdir -p $(BIN_DIR)

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
	rm -f $(BIN_DIR)/$(APP_NAME)$(EXT)

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	go mod tidy
	go install ./...

# Run the application
.PHONY: run
run: build
	@echo "Running the application..."
	./$(BIN_DIR)/$(APP_NAME)$(EXT)

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
