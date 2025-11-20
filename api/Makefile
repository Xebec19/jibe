.PHONY: help run build test clean dev install fmt lint docker-build docker-run postgres createdb dropdb migrateup migratedown

# Default target
help:
	@echo "Available commands:"
	@echo "  make run          - Run the application"
	@echo "  make build        - Build the application"
	@echo "  make test         - Run tests"
	@echo "  make dev          - Run with hot reload (requires air)"
	@echo "  make install      - Install dependencies"
	@echo "  make fmt          - Format code"
	@echo "  make lint         - Run linter"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-run   - Run Docker container"

# Run the application
run:
	go run main.go

# Build the application
build:
	go build -o bin/api main.go

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -cover ./...

# Development mode with hot reload
dev:
	air

# Install dependencies
install:
	go mod download
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	go vet ./...

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f api
	go clean

# Docker build
docker-build:
	docker build -t jibe-api:latest .

# Docker run
docker-run:
	docker run -p 8080:8080 --env-file .env jibe-api:latest

# Install development tools
install-tools:
	go install github.com/air-verse/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb jibe

dropdb:
	docker exec -it postgres12 dropdb jibe

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/jibe?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/jibe?sslmode=disable" -verbose down
