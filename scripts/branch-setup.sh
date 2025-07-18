#!/bin/bash

# Script to set up branch-specific configurations

# Get the current branch name
BRANCH=$(git rev-parse --abbrev-ref HEAD)

# Default environment based on branch
case "$BRANCH" in
  production)
    ENV="production"
    ;;
  development)
    ENV="development"
    ;;
  testing)
    ENV="testing"
    ;;
  feature/* | fix/*)
    ENV="development"
    ;;
  hotfix/*)
    ENV="production"
    ;;
  *)
    ENV="development"
    ;;
esac

# If an argument is provided, override the default
if [ ! -z "$1" ]; then
  ENV="$1"
fi

# Validate environment
if [[ "$ENV" != "development" && "$ENV" != "production" && "$ENV" != "testing" ]]; then
  echo "Invalid environment: $ENV"
  echo "Usage: $0 [development|production|testing]"
  exit 1
fi

echo "Setting up environment for branch '$BRANCH' with environment '$ENV'..."

# Create or update .env file with appropriate settings
if [ -f ".env" ]; then
  # Update existing .env file
  sed -i '' "s/^APP_ENV=.*/APP_ENV=$ENV/" .env
  echo "Updated .env with APP_ENV=$ENV"
else
  # Create new .env file from template
  cp env.example .env
  sed -i '' "s/^APP_ENV=.*/APP_ENV=$ENV/" .env
  echo "Created new .env with APP_ENV=$ENV"
fi

# Create data directory if it doesn't exist
mkdir -p data

echo "Environment setup complete for branch '$BRANCH' with environment '$ENV'"
echo "You can now run 'make dev', 'make prod', or 'make test' to start the application" 