# Makefile

DB_HOST=localhost
DB_PORT=5432
DB_USER=youruser
DB_PASS=yourpassword
DB_NAME=yourdatabase
DB_SSLMODE=disable

MIGRATE_CMD=migrate -database "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" -path migrations

.PHONY: migrate-up migrate-down migrate-force

migrate-up:
	$(MIGRATE_CMD) up

migrate-down:
	$(MIGRATE_CMD) down

migrate-force:
	$(MIGRATE_CMD) force

debug:
	ls -la $(PWD)/migrations

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