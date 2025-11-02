MIGRATIONS_PATH=./migrations

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(GREENLIGHT_DB_DSN) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(GREENLIGHT_DB_DSN) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: seed
seed:
	@go run cmd/migrate/seed/main.go

.PHONY: gen-docs
gen-docs:
	@swag init -g main.go -d cmd/api,internal/store && swag fmt
