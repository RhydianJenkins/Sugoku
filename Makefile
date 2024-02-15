help: ## Show this help
	@grep -E '^[a-z.A-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-50s\033[0m %s\n", $$1, $$2}'

build: ## Build into a binary
	@go build -o sugoku ./main.go

run: ## Run the server
	@go run ./main.go

watch: ## Use nodemon to watch go files for changes
	@nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run ./main.go

test: ## Run all tests
	@go test -v ./pkg/board
