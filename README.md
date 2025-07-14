# Musicfy

Musicfy is a modern, scalable music streaming platform built with Go. It provides user authentication, profile management, and a foundation for music streaming features. The project follows clean architecture principles and is designed for extensibility and maintainability.

## Features

- User registration and authentication (JWT-based)
- Secure password hashing
- User profile management
- RESTful API (versioned)
- Modular, clean architecture
- PostgreSQL database integration
- Environment-based configuration
- Ready for deployment (e.g., Liara, Docker)

## Project Structure

```
internal/
  auth/         # Authentication logic, controllers, services, models, DTOs
  db/           # Database connection and initialization
  shared/       # Shared utilities and response formatting
main.go         # Application entry point
```

## Tech Stack

- **Language:** Go 1.22+
- **Web Framework:** net/http, Gorilla Mux
- **Database:** PostgreSQL (via GORM)
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
3. **Set environment variables:**
   Create a `.env` file in the root directory with the following:
   ```env
   DATABASE_URL=postgres://user:password@localhost:5432/musicfy?sslmode=disable
   JWT_SECRET=your_jwt_secret
   ```
4. **Run database migrations:**
   The app auto-migrates the User model on startup.
5. **Start the server:**
   ```sh
   go run main.go
   ```
   The server will run on `http://localhost:8080` by default.

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

## Environment Variables

- `DATABASE_URL` - PostgreSQL connection string
- `JWT_SECRET` - Secret key for JWT signing

## Deployment

- The project is ready for deployment on platforms like Liara, Docker, or any cloud provider.
- Example `liara.json` is provided for Liara deployment.

## Contributing

1. Fork the repository
2. Create a new branch (`git checkout -b feature/your-feature`)
3. Commit your changes
4. Push to your branch
5. Open a Pull Request

## License

This project is licensed under the MIT License. 