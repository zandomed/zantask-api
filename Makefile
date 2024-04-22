.PHONY: init init.enviroment precommit.rehooks test lint lint.fix run run.docker run.dev deps.upgrade deps.install deps.clean

DOCKER_COMPOSE_DEV_FILE = ./docker/docker-compose.dev.yaml
DOCKER_PROJECT_NAME = zantask

# Default target
init: deps.install

# Install envirments dependencies
init.enviroment:
	@echo "== 👩‍🌾 init =="
	which go || brew install go
	which node || brew install node
	which pre-commit || brew install pre-commit
	which golangci-lint || brew install golangci-lint && brew upgrade golangci-lint

	@echo "== pre-commit setup =="
	pre-commit install --install-hooks --hook-type commit-msg

# Update pre-commit hooks
precommit.rehooks:
	@echo "== 🛠️ Updating pre-commit hooks =="
	pre-commit autoupdate
	pre-commit install --install-hooks
	pre-commit install --hook-type commit-msg

# Run tests
test:
	@echo "== 🦸‍️ Tests =="
	go test -v ./...

# Linting
lint:
	@echo "== 🙆 Linter =="
	golangci-lint run -v ./...

lint.fix:
	@echo "== 🙆 Linter Fix =="
	golangci-lint run -v ./... --fix

# Run the application
run:
	@echo "== 🏃‍♂️ Run =="
	go run cmd/main.go

run.dev:
	@echo "== 🏃‍♂️ Run (Development) =="
	go install github.com/cosmtrek/air@v1.51.0
	air

run.docker:
	@echo "== 🏃‍♂️ Run Docker (Development) =="
	docker-compose -f $(DOCKER_COMPOSE_DEV_FILE) -p $(DOCKER_PROJECT_NAME) build && docker-compose -f $(DOCKER_COMPOSE_DEV_FILE) -p $(DOCKER_PROJECT_NAME) up -d

# Dependency management
deps.upgrade:
	@echo "== 📦 Updating packages =="
	go get -u ./...
	go mod tidy
	go mod verify

deps.install:
	@echo "== 📦 Installing packages =="
	go mod download
	go mod tidy
	go mod verify

deps.clean:
	@echo "== 🧹 Cleaning up =="
	go mod tidy
	go mod verify
	go clean -modcache

deps.verify:
	@echo "== 🧹 Verifying dependencies =="
	go mod tidy
	go mod verify
