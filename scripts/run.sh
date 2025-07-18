#!/bin/bash

# Script to run the application in different environments

# Default to development environment
ENV=${1:-development}

# Validate environment
if [[ "$ENV" != "development" && "$ENV" != "production" && "$ENV" != "testing" ]]; then
  echo "Invalid environment: $ENV"
  echo "Usage: $0 [development|production|testing]"
  exit 1
fi

echo "Starting application in $ENV environment..."

# Set environment variable
export APP_ENV=$ENV

# Run the application
go run main.go 