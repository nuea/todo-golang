HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

tidy: ## A command  that is used to remove unused module dependencies from go.mod and go.sum files.
	go mod tidy

gen: ## Genarate google wire
	wire ./...

run:  ## Running a server to use configs in environment development.
	GIN_MODE=debug ENV=local go run cmd/todo-golang-api/main.go

run_prod:  ## Running a server to use configs in environment production.
	GO111MODULE=on GIN_MODE=release ENV=production go run main.go


help: ## Show this help
	@${HELP_CMD}

.PHONY: run run_dev help local_dummy dev_dummy