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

docker.run:
	docker compose up -d

start: ## Bring up the server on dev environment
	go run cmd/api/main.go

migrate: ## Run database migrations
	go run cmd/migration/main.go

migrate.undo: ## Undo the last database migration
	go run cmd/migration/main.go --down

test:
	go test -v -timeout 30s -run TestLogin funnymovies/internal/api/authen/user
	go test -v -timeout 30s -run TestLoginFailed funnymovies/internal/api/authen/user
	go test -v -timeout 30s -run TestRegister funnymovies/internal/api/authen/user
	go test -v -timeout 30s -run TestList funnymovies/internal/api/public/link
	go test -v -timeout 30s -run TestTotal funnymovies/internal/api/public/link
	go test -v -timeout 30s -run TestView funnymovies/internal/api/user/me
	go test -v -timeout 30s -run TestNonAuthorizedView funnymovies/internal/api/user/me