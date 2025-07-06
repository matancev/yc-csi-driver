# Copyright 2024 YANDEX LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Variables
BUILD_DIR := bin
DOCKER_IMAGE := yc-csi
VERSION ?= $(shell git describe --tags --always --dirty)
LDFLAGS := -ldflags "-X github.com/yandex-cloud/yc-csi-driver/pkg/version.version=$(VERSION)"

# Go parameters
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
GOMOD := $(GOCMD) mod
GOGENERATE := $(GOCMD) generate
GOINSTALL := $(GOCMD) install

.PHONY: all build test clean lint generate vendor deps docker help

# Default target
all: clean generate build test

# Build the binary
build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/ ./cmd/...

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -race -shuffle=on -v -coverprofile=coverage.out ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run --timeout 3m ./...

# Generate code
generate:
	@echo "Generating code..."
	$(GOGENERATE) ./...

# Vendor dependencies
vendor:
	@echo "Vendoring dependencies..."
	$(GOMOD) tidy
	$(GOMOD) vendor

# Build Docker images
docker:
	@echo "Building Docker images..."
	docker build -t $(DOCKER_IMAGE)-node:$(VERSION) -f build/package/Dockerfile.node .
	docker build -t $(DOCKER_IMAGE)-controller:$(VERSION) -f build/package/Dockerfile.controller .

# Help target
help:
	@echo "Available targets:"
	@echo "  all          - Clean, generate, build and test"
	@echo "  build        - Build the binary"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build artifacts"
	@echo "  lint         - Run linter"
	@echo "  generate     - Generate code"
	@echo "  vendor       - Vendor dependencies"
	@echo "  docker       - Build Docker images"
	@echo "  help         - Show this help message"