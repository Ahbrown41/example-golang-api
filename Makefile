VERSION=`git rev-parse HEAD`
BUILD=`date +%FT%T%z`
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

.PHONY: help
help: ## - Show help message
	@printf "\033[32m\xE2\x9c\x93 usage: make [target]\n\n\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build:
	@printf "\033[32m\xE2\x9c\x93 Build image\n\033[0m"
	@export DOCKER_CONTENT_TRUST=1 && docker build -f Dockerfile -t api-reference-golang .

.PHONY: build-no-cache
build-no-cache:
	@printf "\033[32m\xE2\x9c\x93 Build image no cache\n\033[0m"
	@export DOCKER_CONTENT_TRUST=1 && docker build --no-cache -f Dockerfile -t api-reference-golang .

.PHONY: backend_deploy
backend_deploy: ## - Deploy backend stack
	@docker stack deploy --compose-file ./local/backend-compose-stack.yml backend-compose-stack

.PHONY: backend_undeploy
backend_undeploy: ## - Undeploy backend stack
	@docker stack rm backend-compose-stack

.PHONY: ls
ls: ## - List 'api-reference-golang' docker images
	@docker image ls api-reference-golang

.PHONY: test
test: ## - Tests project recursively
	@go test ./...

.PHONY: docs
docs: ## - Generates swagger documentation
	@swag init --parseDependency

.PHONY: run
run:
	@docker run --network host -p 8080:8080 --name api-reference-golang --env-file ./.env api-reference-golang:latest