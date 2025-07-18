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

## Git Branch Structure

The repository is organized with three main branches that correspond to the environments:

- **production**: The main branch for production code (formerly 'main')

  - Contains stable, tested code ready for deployment
  - Protected branch - requires pull request and review to merge
  - Tagged for releases

- **development**: The primary development branch

  - Contains features that are complete but pending testing
  - Target branch for feature branches
  - Periodically merged into testing for integration testing

- **testing**: The branch for integration and system testing
  - Used for comprehensive testing before production
  - Contains features being tested together
  - When stable, merged into production

### Workflow

1. Feature development:

   - Create feature branches from `development`
   - Name format: `feature/feature-name`
   - Merge back to `development` when complete

2. Bug fixes:

   - Create fix branches from `development`
   - Name format: `fix/bug-name`
   - Merge back to `development` when complete

3. Hotfixes:

   - Create hotfix branches from `production`
   - Name format: `hotfix/issue-name`
   - Merge to both `production` and `development` when complete

4. Release process:
   - Test in `testing` branch
   - When ready, merge to `production`
   - Tag with version number
