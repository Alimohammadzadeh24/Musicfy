.PHONY: dev prod test run build clean setup branch-setup

# Default target
all: dev

# Setup the project
setup:
	@echo "Setting up project..."
	@cp -n env.example .env || echo ".env file already exists"
	@mkdir -p data
	@echo "Setup complete. Edit .env file with your configuration."

# Set up branch-specific environment
branch-setup:
	@./scripts/branch-setup.sh

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

# Git branch helpers
git-dev:
	@git checkout development
	@./scripts/branch-setup.sh development

git-prod:
	@git checkout production
	@./scripts/branch-setup.sh production

git-test:
	@git checkout testing
	@./scripts/branch-setup.sh testing

# Help command
help:
	@echo "Musicfy Makefile commands:"
	@echo "  make setup        - Set up the project (copy env.example to .env)"
	@echo "  make branch-setup - Set up environment based on current branch"
	@echo "  make dev          - Run in development mode"
	@echo "  make prod         - Run in production mode"
	@echo "  make test         - Run in testing mode"
	@echo "  make run env=X    - Run in specified environment (development, production, testing)"
	@echo "  make build        - Build for production"
	@echo "  make build env=X  - Build for specified environment"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make git-dev      - Switch to development branch and set up environment"
	@echo "  make git-prod     - Switch to production branch and set up environment"
	@echo "  make git-test     - Switch to testing branch and set up environment" 