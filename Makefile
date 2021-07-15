REPO                  := github.com/PutskouDzmitry/DbTr
PHONY: help
help: ## makefile targets description
	@echo "Usage:"
	@egrep '^[a-zA-Z_-]+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##/#-/' | column -t -s "#"

.PHONY: fmt
fmt: ## automatically formats Go source code
	@echo "Running 'go fmt ...'"
	@go fmt -x "$(REPO)/..."

.PHONY: db
db: fmt
	@docker build -f ./docker/db/Dockerfile -t kvarc/db-task:latest .

.PHONY: image
image: db ## build image from Dockerfile ./docker/server/Dockerfile
	@docker build -t kvarc/task-dbtr:latest .

.PHONY: up
up : image ## up docker compose
	@docker-compose up -d

.PHONY: integration
integration: up
	@go test --tags=integration ./cmd/server/

.PHONY: down
down : integration
	@docker-compose down