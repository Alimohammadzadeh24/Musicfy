# Auth Module

This module handles user authentication and authorization in the Musicfy application, following clean architecture principles.

## Structure

The auth module follows a clean architecture approach with the following layers:

```
auth/
├── domain/           # Core business logic and interfaces
│   ├── entities/     # Domain entities
│   ├── repositories/ # Repository interfaces
│   ├── usecases/     # Use cases and business logic
│   └── errors.go     # Domain-specific errors
├── data/             # Data layer implementation
│   ├── repositories/ # Repository implementations
│   └── services/     # Service implementations
├── presentation/     # Presentation layer
│   ├── controllers/  # HTTP controllers
│   ├── dtos/         # Data Transfer Objects
│   ├── middleware/   # HTTP middleware
│   └── routes/       # Route definitions
└── module.go         # Module entry point
```

## Clean Architecture Guidelines

### Domain Layer

- Contains core business logic and entities
- Has no dependencies on external frameworks or libraries
- Defines interfaces that will be implemented by outer layers
- Contains use cases that orchestrate the flow of data to and from entities

### Data Layer

- Implements repository interfaces defined in the domain layer
- Handles data persistence and external service communication
- Contains concrete implementations of data sources

### Presentation Layer

- Handles HTTP requests and responses
- Contains controllers that use domain use cases
- Defines DTOs for data transfer between the API and domain
- Implements middleware for request processing

## Development Rules

1. **Dependency Rule**: Dependencies should only point inward. Domain layer should not depend on data or presentation layers.

2. **Entity Rule**: Entities are the core business objects and should be kept clean from framework code.

3. **Interface Segregation**: Define interfaces in the domain layer and implement them in outer layers.

4. **Use Case Rule**: Use cases should orchestrate the flow of data to and from entities and should contain business rules specific to the application.

5. **Naming Conventions**:

   - Interfaces should be named without prefixes (e.g., `UserRepository`)
   - Implementations should be named with descriptive suffixes (e.g., `UserRepositoryImpl`)
   - Use cases should be named with action verbs (e.g., `RegisterUser`, `LoginUser`)

6. **Error Handling**: Domain-specific errors should be defined in the domain layer and handled appropriately in outer layers.

7. **Testing**: Each layer should be testable in isolation. Use mocks for dependencies.

## Adding New Features

When adding new features to the auth module:

1. Start by defining domain entities and use cases
2. Implement repository interfaces in the data layer
3. Create DTOs and controllers in the presentation layer
4. Register new routes in the routes file
5. Update the module.go file if necessary

## Authentication Flow

1. User registers with email, username, and password
2. User logs in with username/email and password
3. Server validates credentials and returns JWT token
4. Client includes JWT token in Authorization header for protected routes
5. JWT middleware validates the token and adds user info to request context
