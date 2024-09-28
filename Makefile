.PHONY: build clean deploy

HOST ?= localhost:8000

## ===== Local =====

depends: ## Install & build dependencies
	go get ./...
	go build ./...
	go mod tidy

mod.clean:
	go clean -cache
	go clean -modcache

mod: ## Update dependencies
	go mod tidy && go mod vendor

provision: depends ## Provision dev environment
	sudo docker compose up -d
	scripts/waitdb.sh
	@$(MAKE) migrate

docker.run:
	sudo docker compose up -d

start: ## Bring up the server on dev environment
	go run cmd/api/main.go

migrate: ## Run database migrations
	go run cmd/migration/main.go

migrate.undo: ## Undo the last database migration
	go run cmd/migration/main.go --down

clean: 
	rm -rf ./server ./*.out
	rm -rf .serverless 0