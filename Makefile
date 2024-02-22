run:
	@echo "running server"
	@go run ./cmd/app

migrate:
	@echo "migrating database"
	@go run ./cmd/migrate

.PHONY: run migrate
