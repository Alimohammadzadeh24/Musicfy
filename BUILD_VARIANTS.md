# Build Variants

This document explains how to use the different build variants (development, production, testing) in the Musicfy application.

## Overview

The application supports three different environments:

1. **Development**: Used for local development with debug features enabled
2. **Production**: Used for production deployment with optimized settings
3. **Testing**: Used for running tests with isolated configuration

## Environment Configuration

The application uses a single `.env` file for configuration. The environment is determined by the `APP_ENV` variable, which can be set to:

- `development` (default)
- `production`
- `testing`

## Example .env File

```env
# Application Environment (development, production, testing)
APP_ENV=development
APP_PORT=8080
APP_HOST=localhost

# Database Configuration
DATABASE_URL=postgres://postgres:postgres@localhost:5432/musicfy_dev
DB_MAX_CONNS=10
DB_IDLE_CONNS=5

# JWT Configuration
JWT_SECRET=your_secret_key_here
JWT_EXPIRY_HOURS=24
```

## Running in Different Environments

### Using Make Commands

The simplest way to run the application in different environments is using the provided Make commands:

```bash
# Run in development mode
make dev

# Run in production mode
make prod

# Run in testing mode
make test

# Run in a specific environment
make run env=development
make run env=production
make run env=testing
```

### Using the Run Script

You can also use the provided shell script:

```bash
# Run in development mode (default)
./scripts/run.sh

# Run in production mode
./scripts/run.sh production

# Run in testing mode
./scripts/run.sh testing
```

### Setting Environment Variable Manually

You can set the `APP_ENV` environment variable manually:

```bash
# Run in development mode
APP_ENV=development go run main.go

# Run in production mode
APP_ENV=production go run main.go

# Run in testing mode
APP_ENV=testing go run main.go
```

## Building for Different Environments

To build the application for a specific environment:

```bash
# Build for production (default)
make build

# Build for a specific environment
make build env=development
make build env=production
make build env=testing
```

This will create an executable in the `bin/` directory with environment-specific optimizations.

## Environment Features

### Development Environment

- Detailed logging
- Additional debug endpoints
- Environment indicator in response headers
- Hot reloading (if configured)

### Production Environment

- Optimized performance
- Minimal logging (errors only)
- No debug endpoints
- Security headers

### Testing Environment

- Isolated database
- Mock external services
- Comprehensive logging for test results
- Shorter token expiration times

## Git Branches

The repository is organized with three main branches corresponding to the environments:

- `production`: The main production branch, stable and ready for deployment
- `development`: The development branch for ongoing development
- `testing`: The testing branch for integration and system tests

Each branch may contain environment-specific configurations and features.
