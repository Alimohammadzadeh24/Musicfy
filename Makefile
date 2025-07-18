.PHONY: dev prod test run build clean

# Default target
all: dev

# Run in development mode
dev:
	@echo "Starting in development mode..."
	@APP_ENV=development go run main.go

# Run in production mode
prod:
	@echo "Starting in production mode..."
	@APP_ENV=production go run main.go

# Run in testing mode
test:
	@echo "Starting in testing mode..."
	@APP_ENV=testing go run main.go

# Run with specified environment
run:
	@./scripts/run.sh $(env)

# Build for specified environment (default: production)
build:
	@echo "Building for $(or $(env),production) environment..."
	@APP_ENV=$(or $(env),production) go build -o bin/musicfy main.go

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@go clean

# Help command
help:
	@echo "Musicfy Makefile commands:"
	@echo "  make dev        - Run in development mode"
	@echo "  make prod       - Run in production mode"
	@echo "  make test       - Run in testing mode"
	@echo "  make run env=X  - Run in specified environment (development, production, testing)"
	@echo "  make build      - Build for production"
	@echo "  make build env=X - Build for specified environment"
	@echo "  make clean      - Clean build artifacts" 