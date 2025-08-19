.PHONY: help build run test clean docker-build docker-run deploy

# Default target
help:
	@echo "Available commands:"
	@echo "  build        - Build the Go binary"
	@echo "  run          - Run the application locally"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build artifacts"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"
	@echo "  deploy       - Deploy to Railway"

# Build the Go binary
build:
	@echo "Building Go binary..."
	go build -o go-backend-api main.go
	@echo "Build complete! Binary: go-backend-api"

# Run the application locally
run:
	@echo "Starting Go Backend API..."
	go run main.go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f go-backend-api
	@echo "Clean complete!"

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t go-backend-api .
	@echo "Docker build complete!"

# Run Docker container
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 go-backend-api

# Deploy to Railway
deploy:
	@echo "Deploying to Railway..."
	./deploy.sh

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy
	@echo "Dependencies installed!"

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...
	@echo "Code formatted!"

# Lint code
lint:
	@echo "Linting code..."
	golangci-lint run
	@echo "Linting complete!"
