# Musicfy

Musicfy is a modern, scalable music streaming platform built with Go. It provides user authentication, profile management, and a foundation for music streaming features. The project follows clean architecture principles and is designed for extensibility and maintainability.

## Features

- User registration and authentication (JWT-based)
- Secure password hashing
- User profile management
- RESTful API (versioned)
- Modular, clean architecture
- PostgreSQL database integration
- Environment-based configuration (development, production, testing)
- Ready for deployment (e.g., Liara, Docker)

## Branch Structure

The repository is organized with three main branches:

- **production**: The main production branch, stable and ready for deployment
- **development**: The development branch for ongoing development
- **testing**: The testing branch for integration and system tests

## Project Structure

```
internal/
  auth/         # Authentication logic, controllers, services, models, DTOs
  config/       # Configuration management
  db/           # Database connection and initialization
  shared/       # Shared utilities and response formatting
config/         # Environment configuration files
scripts/        # Utility scripts
main.go         # Application entry point
```

## Tech Stack

- **Language:** Go 1.22+
- **Web Framework:** net/http, Gorilla Mux
- **Database:** PostgreSQL (via database/sql)
- **Auth:** JWT (github.com/golang-jwt/jwt)
- **Validation:** go-playground/validator

## Getting Started

### Prerequisites

- Go 1.22 or higher
- PostgreSQL database

### Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/musicfy.git
   cd musicfy
   ```
2. **Install dependencies:**
   ```sh
   go mod download
   ```
3. **Set up environment configuration:**
   Copy the example environment file and edit it with your settings:

   ```sh
   cp env.example .env
   ```

   Then edit the `.env` file with your specific configuration.

4. **Run the application:**

   ```sh
   # Run in development mode (default)
   make dev

   # Or run in production mode
   make prod

   # Or run in testing mode
   make test
   ```

## Environment Configuration

Musicfy uses a single `.env` file for configuration. The application behavior changes based on the `APP_ENV` setting:

- `APP_ENV=development` - Development mode with detailed logging and debug endpoints
- `APP_ENV=production` - Production mode with optimized settings
- `APP_ENV=testing` - Testing mode with isolated configuration

Key environment variables:

- `APP_ENV` - Application environment (development, production, testing)
- `APP_PORT` - Port to run the server on
- `APP_HOST` - Host to bind the server to
- `DATABASE_URL` - Database connection string
- `DB_MAX_CONNS` - Maximum database connections
- `DB_IDLE_CONNS` - Maximum idle database connections
- `JWT_SECRET` - Secret key for JWT signing
- `JWT_EXPIRY_HOURS` - JWT token expiry in hours

## API Endpoints

All endpoints are prefixed with `/api/v1`.

### Authentication

- **POST /api/v1/auth/register**
  - Register a new user.
  - Request body:
    ```json
    {
      "first_name": "John",
      "last_name": "Doe",
      "username": "johndoe",
      "password": "yourpassword",
      "email": "john@example.com",
      "age": "25"
    }
    ```
- **POST /api/v1/auth/login**
  - Login with username/email and password.
  - Request body:
    ```json
    {
      "username_or_email": "johndoe",
      "password": "yourpassword"
    }
    ```
  - Response:
    ```json
    {
      "is_success": true,
      "message": "Login successful",
      "data": { "token": "<jwt_token>" }
    }
    ```
- **GET /api/v1/auth/user/profile**
  - Get the authenticated user's profile.
  - Requires `Authorization: Bearer <token>` header.
  - Response:
    ```json
    {
      "is_success": true,
      "message": "User retrieved successfully",
      "data": {
        "first_name": "John",
        "last_name": "Doe",
        "username": "johndoe",
        "email": "john@example.com",
        "age": 25,
        "created_at": "2024-06-10T12:00:00Z",
        "updated_at": "2024-06-10T12:00:00Z"
      }
    }
    ```

## Running in Different Environments

You can run the application in different environments using the provided Makefile commands:

```sh
# Development mode
make dev

# Production mode
make prod

# Testing mode
make test
```

Or using the run script:

```sh
./scripts/run.sh [development|production|testing]
```

## Deployment

- The project is ready for deployment on platforms like Liara, Docker, or any cloud provider.
- Example `liara.json` is provided for Liara deployment.
- For production deployment, make sure to set all required environment variables.

## Contributing

1. Fork the repository
2. Create a new branch from the appropriate base branch:
   - For new features: `git checkout -b feature/your-feature development`
   - For bug fixes: `git checkout -b fix/your-fix development`
   - For hotfixes: `git checkout -b hotfix/your-hotfix production`
3. Commit your changes
4. Push to your branch
5. Open a Pull Request to the appropriate target branch

## License

This project is licensed under the MIT License.
