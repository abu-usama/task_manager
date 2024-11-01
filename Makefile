# run: main.go
# 	 nodemon --watch './**/*.go' --ignore './docs/*' --signal SIGTERM --exec 'go' run main.go
# Load environment variables from .env
include config/.env.development

.PHONY: help run run-hot swag-init migrate-create migrate-up migrate-down

help: ## Show available targets
	@echo "Available targets:"
	@echo "  run                 Run the application without hot reload"
	@echo "  run-hot             Run the application with nodemon for hot reload"
	@echo "  swag-init           Initialize Swagger documentation"
	@echo "  migrate-create      Create a new migration"
	@echo "  migrate-up          Run migrations up sequentially"
	@echo "  migrate-down        Rollback migrations sequentially"
	@echo "  deploy              Deploy the application using deploy.sh"


run: main.go ## Run the application without hot reload
	go run main.go

# run-hot: main.go ## Run the application with nodemon for hot reload
# 	nodemon --watch './**/*.go' --ignore './docs/*' --signal SIGTERM --exec go run main.go

swag-init: ## Initialize Swagger documentation
	swag init

migrate-create: ## Create a new migration
	migrate create -ext sql -dir database/migrations -seq 'my_new_migration'

migrate-up: ## Run migrations up sequentially
	migrate -path database/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" up

migrate-down: ## Rollback migrations sequentially
	migrate -path database/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" down

# deploy: ## Deploy the application using deploy.sh
# 	./deploy.sh