.PHONY: help build test clean install lint coverage run dev fmt vet deps release-local release bump-version

BINARY_NAME=httpyum
BINARY_PATH=./$(BINARY_NAME)
CMD_PATH=./cmd/httpyum
INSTALL_PATH=/usr/local/bin
VERSION?=dev
LDFLAGS=-ldflags "-s -w -X 'httpyum/internal/config.version=$(VERSION)'"

help:
	@echo "Available targets:"
	@echo "  make build          - Build the binary"
	@echo "  make test           - Run tests"
	@echo "  make test-verbose   - Run tests with verbose output"
	@echo "  make coverage       - Generate test coverage report"
	@echo "  make lint           - Run linters (go vet + staticcheck if available)"
	@echo "  make fmt            - Format code with go fmt"
	@echo "  make vet            - Run go vet"
	@echo "  make clean          - Remove build artifacts"
	@echo "  make install        - Install binary to $(INSTALL_PATH)"
	@echo "  make uninstall      - Remove binary from $(INSTALL_PATH)"
	@echo "  make run            - Build and run with example.http"
	@echo "  make dev            - Run in development mode (no build)"
	@echo "  make deps           - Download dependencies"
	@echo "  make tidy           - Tidy go modules"
	@echo "  make release-local  - Build release binaries for all platforms"
	@echo "  make bump-version   - Bump version and trigger auto-release (VERSION=v1.0.0 make bump-version)"
	@echo "  make release        - Create a GitHub release (requires gh CLI, VERSION=v1.0.0 make release)"
	@echo ""
	@echo "Environment variables:"
	@echo "  VERSION             - Set version (default: dev)"
	@echo "  INSTALL_PATH        - Installation path (default: /usr/local/bin)"

build:
	@echo "Building $(BINARY_NAME)..."
	go build $(LDFLAGS) -o $(BINARY_PATH) $(CMD_PATH)
	@echo "Build complete: $(BINARY_PATH)"

test:
	@echo "Running tests..."
	go test ./...

test-verbose:
	@echo "Running tests (verbose)..."
	go test -v ./...

test-race:
	@echo "Running tests with race detector..."
	go test -race ./...

coverage:
	@echo "Generating coverage report..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint: vet
	@echo "Running linters..."
	@if command -v staticcheck >/dev/null 2>&1; then \
		staticcheck ./...; \
	else \
		echo "staticcheck not installed. Install with: go install honnef.co/go/tools/cmd/staticcheck@latest"; \
	fi

fmt:
	@echo "Formatting code..."
	go fmt ./...

vet:
	@echo "Running go vet..."
	go vet ./...

clean:
	@echo "Cleaning build artifacts..."
	rm -f $(BINARY_PATH)
	rm -f coverage.out coverage.html coverage.txt
	rm -f httpyum-*
	rm -rf dist/
	@echo "Clean complete"

install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
	@if [ -w "$(INSTALL_PATH)" ]; then \
		cp $(BINARY_PATH) $(INSTALL_PATH)/$(BINARY_NAME); \
	else \
		sudo cp $(BINARY_PATH) $(INSTALL_PATH)/$(BINARY_NAME); \
	fi
	@echo "Installed: $(INSTALL_PATH)/$(BINARY_NAME)"

uninstall:
	@echo "Uninstalling $(BINARY_NAME) from $(INSTALL_PATH)..."
	@if [ -w "$(INSTALL_PATH)" ]; then \
		rm -f $(INSTALL_PATH)/$(BINARY_NAME); \
	else \
		sudo rm -f $(INSTALL_PATH)/$(BINARY_NAME); \
	fi
	@echo "Uninstalled"

run: build
	@echo "Running $(BINARY_NAME)..."
	$(BINARY_PATH) example.http

dev:
	@echo "Running in development mode..."
	go run $(CMD_PATH) example.http

deps:
	@echo "Downloading dependencies..."
	go mod download

tidy:
	@echo "Tidying go modules..."
	go mod tidy

release-local:
	@echo "Building release binaries for all platforms..."
	@mkdir -p dist
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-amd64 $(CMD_PATH)
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-arm64 $(CMD_PATH)
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-amd64 $(CMD_PATH)
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-arm64 $(CMD_PATH)
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-windows-amd64.exe $(CMD_PATH)
	GOOS=windows GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-windows-arm64.exe $(CMD_PATH)
	@echo "Release binaries built in dist/"
	@ls -lh dist/

bump-version:
	@if [ -z "$(VERSION)" ] || [ "$(VERSION)" = "dev" ]; then \
		echo "Error: VERSION must be set (e.g., VERSION=v1.0.0 make bump-version)"; \
		exit 1; \
	fi
	@echo "$(VERSION)" | grep -qE '^v[0-9]+\.[0-9]+\.[0-9]+$$' || (echo "Error: VERSION must be in format vX.Y.Z"; exit 1)
	@echo "Updating VERSION file to $(VERSION)..."
	@echo "$(VERSION)" > VERSION
	@echo "VERSION file updated. Commit and push to trigger auto-release:"
	@echo "  git add VERSION"
	@echo "  git commit -m 'chore: bump version to $(VERSION)'"
	@echo "  git push"

release:
	@if [ -z "$(VERSION)" ] || [ "$(VERSION)" = "dev" ]; then \
		echo "Error: VERSION must be set (e.g., VERSION=v1.0.0 make release)"; \
		exit 1; \
	fi
	@echo "$(VERSION)" | grep -qE '^v[0-9]+\.[0-9]+\.[0-9]+$$' || (echo "Error: VERSION must be in format vX.Y.Z"; exit 1)
	@echo "Creating release $(VERSION) using GitHub CLI..."
	@if ! command -v gh >/dev/null 2>&1; then \
		echo "Error: GitHub CLI (gh) not installed. Install with: brew install gh"; \
		exit 1; \
	fi
	@echo "Creating tag $(VERSION)..."
	@git tag -a "$(VERSION)" -m "Release $(VERSION)" || (echo "Tag may already exist"; exit 1)
	@echo "Pushing tag to GitHub..."
	@git push origin "$(VERSION)"
	@echo "✅ Tag pushed! GitHub Actions will now build and create the release automatically."
	@echo "Monitor progress at: https://github.com/$$(git config --get remote.origin.url | sed 's/.*github.com[:/]\(.*\).git/\1/')/actions"

all: clean deps fmt vet test build
