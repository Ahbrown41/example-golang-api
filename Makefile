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
	@export DOCKER_CONTENT_TRUST=1 && docker build -f Dockerfile -t reference-golang-svc .

.PHONY: build-no-cache
build-no-cache:
	@printf "\033[32m\xE2\x9c\x93 Build image no cache\n\033[0m"
	@export DOCKER_CONTENT_TRUST=1 && docker build --no-cache -f Dockerfile -t reference-golang-svc .

.PHONY: ls
ls: ## - List 'reference-golang-svc' docker images
	@docker image ls reference-golang-svc

.PHONY: run
run:
	@docker run --network host -p 8080:8080 --name reference-golang-svc --env-file ./.env reference-golang-svc:latest