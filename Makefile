.PHONY: help run app

deps:
	go mod download
	go mod tidy

run:
	go run ./cmd/server/main.go

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
