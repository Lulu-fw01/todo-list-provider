.PHONY: help build run test clean docker-build docker-run docker-stop deps

# Default target
help:
	@echo "Available commands:"
	@echo "  build        - Build the Go application"
	@echo "  run          - Run the application locally"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build artifacts"
	@echo "  deps         - Download Go dependencies"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run with Docker Compose"
	@echo "  docker-stop  - Stop Docker Compose services"

# Download dependencies
deps:
	go mod download
	go mod tidy

# Build the application
build: deps
	go build -o bin/server ./cmd/server

# Run the application locally
run: build
	./bin/server

# Run tests
test: deps
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/
	go clean

# Build Docker image
docker-build:
	docker build -t todo-list-provider .

# Run with Docker Compose
docker-run:
	docker-compose up -d

# Stop Docker Compose services
docker-stop:
	docker-compose down

# Show logs
docker-logs:
	docker-compose logs -f

# Reset database
docker-reset:
	docker-compose down -v
	docker-compose up -d postgres
	sleep 5
	docker-compose up -d app
