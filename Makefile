.PHONY: init init.enviroment precommit.rehooks test lint lint.fix run deps.upgrade deps.install deps.clean

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
