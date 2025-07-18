# Repository Pattern Implementation

This document explains the implementation of the repository pattern in the Musicfy application, replacing direct ORM (GORM) usage with SQL queries.

## Overview

The repository pattern is a design pattern that separates the data access logic from the business logic. It provides a clean API for the business logic to access data without knowing the underlying data storage details.

## Benefits

1. **Separation of Concerns**: Business logic is separated from data access logic
2. **Testability**: Repositories can be easily mocked for testing
3. **Flexibility**: The underlying data storage can be changed without affecting the business logic
4. **Query Optimization**: Direct SQL queries can be optimized for specific use cases

## Implementation

### 1. Repository Interface

```go
type UserRepository interface {
    Create(user *models.User) error
    FindByUsername(username string) (*models.User, error)
    FindByEmail(email string) (*models.User, error)
    FindByUsernameOrEmail(usernameOrEmail string) (*models.User, error)
    FindByID(id uuid.UUID) (*models.User, error)
    Update(user *models.User) error
}
```

### 2. PostgreSQL Implementation

```go
type PostgresUserRepository struct {
    db *sql.DB
}
```

### 3. SQL Queries

Instead of using GORM's fluent API:

```go
// GORM approach (old)
dbInstance.Where("username = ?", username).Or("email = ?", email).First(&existingUser)
```

We now use direct SQL queries:

```go
// SQL approach (new)
query := `
    SELECT id, first_name, last_name, username, email, age, password_hash, created_at, updated_at
    FROM users
    WHERE username = $1 OR email = $1
`
row := r.db.QueryRow(query, usernameOrEmail)
```

### 4. Database Migrations

Instead of using GORM's AutoMigrate:

```go
// GORM approach (old)
database.AutoMigrate(&models.User{})
```

We now use SQL migration files:

```sql
-- SQL approach (new)
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    -- other fields
);
```

## SQL Commands Reference

### Create Table

```sql
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    username VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    age INTEGER NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
```

### Create Indexes

```sql
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
```

### Insert User

```sql
INSERT INTO users (id, first_name, last_name, username, email, age, password_hash, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
```

### Find User by Username or Email

```sql
SELECT id, first_name, last_name, username, email, age, password_hash, created_at, updated_at
FROM users
WHERE username = $1 OR email = $1
```

### Update User

```sql
UPDATE users
SET first_name = $1, last_name = $2, username = $3, email = $4,
    age = $5, password_hash = $6, updated_at = $7
WHERE id = $8
```

## Migration System

The migration system:

1. Tracks applied migrations in a `migrations` table
2. Reads SQL files from the `migrations` directory
3. Applies migrations in alphabetical order
4. Runs each migration in a transaction
5. Records successful migrations in the `migrations` table

## Usage in Service Layer

The service layer now uses the repository interface instead of direct database access:

```go
// Initialize repository
var userRepo = repository.NewUserRepository()

// Use repository in service functions
func GetUserByIDService(id uuid.UUID) (*models.User, error) {
    user, err := userRepo.FindByID(id)
    if err != nil {
        return nil, err
    }
    if user == nil {
        return nil, ErrUserNotFound
    }
    return user, nil
}
```
