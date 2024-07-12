# Variables
APP_NAME := keystore
SRC_DIRS := $(shell go list -f '{{.Dir}}' ./...)
CMD_DIR := ./cmd/$(APP_NAME)
BIN_DIR := bin

# Default target
.PHONY: all
all: build

$(BIN_DIR):
	@mkdir -p $(BIN_DIR)

# Build the project
.PHONY: build
build: $(BIN_DIR)
	@echo "Building the project ..."
	go build -o $(BIN_DIR)/$(APP_NAME)$(EXT) $(CMD_DIR)

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

.PHONY: check
check: fmt vet test
	@echo "All checks passed"

# Run the application
.PHONY: run
run: build
	@echo "Running the application..."
	./$(BIN_DIR)/$(APP_NAME)$(EXT)
