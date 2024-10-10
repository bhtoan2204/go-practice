# Makefile

DB_HOST = localhost
DB_PORT = 3306
DB_USER = root
DB_PASS = root
DB_NAME = simple_bank

MIGRATION_DIR := ./migrations
DB_URL := mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?charset=utf8mb4&parseTime=True&multiStatements=true
MIGRATE := migrate
MIGRATE_CMD = $(MIGRATE) -path $(MIGRATION_DIR) -database "$(DB_URL)"

.PHONY: migrate-up
migrate-up:
	@echo "Running migrate up..."
	@$(MIGRATE_CMD) up

.PHONY: migrate-down
migrate-down:
	@echo "Running migrate down..."
	@$(MIGRATE_CMD) down

.PHONY: migrate-force
migrate-force:
	@echo "Forcing migration version..."
	@$(MIGRATE_CMD) force 4

lint:
	@echo "Linting Go files..."
	@golangci-lint run ./...
	@echo "Linting Go files completed"

run-dev:
	@echo "Running Go application in development mode..."
	@go run -race ./cmd/server/main.go

run:
	@echo "Running Go application in normal mode..."
	@go run ./cmd/server/main.go

build:
	@echo "Building Go application..."
	@go build -o bin/main ./cmd/main.go

run-bin:
	@echo "Running Go application from binary..."
	@./main