ifneq (,$(wildcard ./.env))
    include .env
    export
endif

DOCKER_CONTAINER=dfcu_loan_manager_postgres_1
DB_URL=$(DB_DRIVER)://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable

createdb:
	docker exec -it $(DOCKER_CONTAINER) createdb --username=root --owner=root $(DB_NAME)
	
dropdb:
	docker exec -it $(DOCKER_CONTAINER) dropdb dfcu_bank

migrateup: ## Run all migrations up
	migrate --path db/migration --database "$(DB_URL)" --verbose up

migrateup1: ## Run migration up by single step
	migrate --path db/migration --database "$(DB_URL)" --verbose up 1

migratedown:  ## Run all migrations down
	migrate --path db/migration --database "$(DB_URL)" --verbose down

migratedown1: ## Run migration down by single step
	migrate --path db/migration --database "$(DB_URL)" --verbose down 1

db_docs: ## Generate database docs
	dbdocs build doc/db.dbml

db_schema: ## Generate database schema
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml 

sqlc: ## Generate sqlc files
	sqlc generate

test: ## Run tests
	go test -v -cover ./...
	
server: ## Run server
	go run main.go

mock: ## Generate test mocks
	mockgen --build_flags=--mod=mod --destination db/mock/store.go --package mockdb github.com/wizlif/dfcu_bank/db/sqlc Store 

simulation: ## Run simulation script
	go run simulator/simulator.go

key: ## Generate assymeric key
	openssl rand -hex 64 | head -c 32

build: ## Build server
	go build -o server main.go 

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: simulation createdb dropdb migratedown migratedown1 migrateup migrateup1 test server mock key
.DEFAULT_GOAL := help
