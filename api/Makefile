help:
	@echo "Available commands:"
	@echo "dev: run development server"
	@echo "build: build executable"
	@echo "test-coverage: run tests"

dev:
	go run cmd/server

build:
	go build -o bin/api cmd/server

test-coverage:
	go text -v cover ./...

clean:
	rm -rf bin
	rm -f api
	go clean

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb jibe

dropdb:
	docker exec -it postgres12 dropdb jibe

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/jibe?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/jibe?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY:
	help dev build test-coverage clean postgres createdb dropdb migrateup migratedown sqlc