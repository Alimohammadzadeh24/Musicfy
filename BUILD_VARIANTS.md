# Build Variants

This document explains how to use the different build variants (development, production, testing) in the Musicfy application.

## Overview

The application supports three different environments:

1. **Development**: Used for local development with debug features enabled
2. **Production**: Used for production deployment with optimized settings
3. **Testing**: Used for running tests with isolated configuration

## Environment-Specific Configuration

Each environment has its own configuration file in the `config/env` directory:

- `config/env/development.env`: Development environment configuration
- `config/env/production.env`: Production environment configuration
- `config/env/testing.env`: Testing environment configuration

These files contain environment-specific settings for:

- Application host and port
- Database connection parameters
- JWT configuration
- Logging levels
- Other environment-specific settings

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
