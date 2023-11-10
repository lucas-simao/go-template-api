help:: ## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | sort | fgrep -v fgrep | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

db-up: ## Initialize db in Docker
	@cd docker && docker-compose up -d

db-down: ## Finalize db in Docker
	@cd docker && docker-compose down

db-down-volume: ## Finalize db in Docker and remove data
	@cd docker && docker-compose down -v

api-up: ## Run api - Teste
	@make db-up
	@go run main.go

test: ## Go test
	go test ./...

test_coverage: ## Go test with coverage file
	go test -race -coverprofile coverage.out ./...

get_dependencies: ## Go get dependencies
	 go get -v -t -d ./...