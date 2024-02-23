dev:
	@echo "MAKE: dev"
	@go run ./cmd/dev

migrate:
	@echo "MAKE: migrate"
	@go run ./cmd/migrate

.PHONY: dev migrate
