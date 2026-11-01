.PHONY: help run build fmt vet check clean

APP := bin/app

help: ## Show available targets
	@grep -E '^[a-z-]+:.*## ' $(MAKEFILE_LIST) | awk -F':.*## ' '{printf "  %-8s %s\n", $$1, $$2}'

run: ## Run the server locally (reads .env if present)
	@set -a; [ -f .env ] && . ./.env; set +a; go run .

build: ## Build the binary into bin/
	CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o $(APP) .

fmt: ## Format Go code
	gofmt -w .

vet: ## Run go vet
	go vet ./...

check: vet ## Fail if code is not formatted
	@test -z "$$(gofmt -l .)" || (gofmt -l .; exit 1)

clean: ## Remove build output
	rm -rf bin
