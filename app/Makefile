.PHONY: build run test clean docker-build docker-up docker-down

# Build the application
build:
	go build -o bin/todo-service .

# Run the application
run:
	go run .

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod tidy
	go mod download

# Build Docker image
docker-build:
	docker build -t todo-service .

# Start services with Docker Compose
docker-up:
	docker-compose up --build -d

# Stop services
docker-down:
	docker-compose down

# View logs
logs:
	docker-compose logs -f

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Development setup
dev-setup: deps
	@echo "Development environment ready!"

# Production build
prod-build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/todo-service .

help:
	@echo "Available commands:"
	@echo "  build       - Build the application"
	@echo "  run         - Run the application"
	@echo "  test        - Run tests"
	@echo "  clean       - Clean build artifacts"
	@echo "  deps        - Install dependencies"
	@echo "  docker-build- Build Docker image"
	@echo "  docker-up   - Start services with Docker Compose"
	@echo "  docker-down - Stop services"
	@echo "  logs        - View Docker Compose logs"
	@echo "  fmt         - Format code"
	@echo "  lint        - Lint code"
	@echo "  dev-setup   - Setup development environment"
	@echo "  prod-build  - Production build"
