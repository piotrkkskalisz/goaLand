# Simple Makefile for a Go project

# Create .env file from .env.example file
env:
	@cp -n .env.example .env

# Build the application
build:
	@echo "Building..."
	@mkdir -p bin
	@go build -o bin/main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Create DB container
docker-up:
	@docker compose up --build
		

# Shutdown DB container
docker-down:
	@docker compose down

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -rf bin

# Live Reload
watch:
	@air

.PHONY: build run test clean watch docker-up docker-down itest