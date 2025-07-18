# Go Controller Refactoring Rules

This document outlines the rules and patterns applied when refactoring Go controllers, using the auth controller as an example.

## General Principles

1. **Single Responsibility Principle**: Each function should have a single responsibility.
2. **DRY (Don't Repeat Yourself)**: Extract common code into reusable functions.
3. **Error Handling**: Centralize error handling for consistent responses.
4. **Clear Structure**: Organize code into logical sections with comments.
5. **Descriptive Names**: Use clear, descriptive names for functions and variables.

## Specific Rules

### 1. Controller Organization

- Add descriptive comments for each controller function
- Structure each controller function into clear sections:
  1. Request parsing and validation
  2. Business logic through service calls
  3. Response formatting and sending
- Keep controllers focused on HTTP concerns, delegating business logic to services

### 2. Reusable Components

- Create shared validator instances at package level instead of in each function
- Extract common request parsing and validation logic into helper functions
- Centralize error handling for service layer errors

### 3. Error Handling

- Use a dedicated function for handling service errors
- Map domain errors to appropriate HTTP status codes
- Provide clear error messages to clients
- Avoid exposing internal error details in production

### 4. Helper Functions

- Create helper functions for:
  - Request decoding and validation
  - Context value extraction
  - Response mapping
  - Error handling

### 5. Response Formatting

- Use consistent response structures
- Set appropriate HTTP status codes
- Include only necessary information in responses

### 6. Code Style

- Remove unnecessary debug print statements (e.g., fmt.Println)
- Use consistent naming conventions
- Add comments for non-obvious code sections
- Organize imports logically

## Examples

### Before Refactoring:

```go
func LoginUserController(w http.ResponseWriter, r *http.Request) {
    validate := validator.New()
    var req dtos.LoginRequestDto

    if error := json.NewDecoder(r.Body).Decode(&req); error != nil {
        shared.Error(w, http.StatusBadRequest, "Invalid JSON body", error.Error())
        return
    }
    if error := validate.Struct(req); error != nil {
        http.Error(w, "Validation failed: "+error.Error(), http.StatusBadRequest)
        return
    }

    token, err := LoginUserService(req.UsernameOrEmail, req.Password)
    if err != nil {
        switch {
        case errors.Is(err, ErrUserNotFound):
            shared.Error(w, http.StatusNotFound, err.Error(), err.Error())
        // More error cases...
        default:
            shared.Error(w, http.StatusInternalServerError, "", err.Error())
        }
        return
    }

    w.WriteHeader(http.StatusOK)
    shared.Success(w, "Login successful", map[string]string{
        "token": token,
    })
}
```

### After Refactoring:

```go
// LoginUserController handles user login requests
func LoginUserController(w http.ResponseWriter, r *http.Request) {
    // Parse and validate request body
    var req dtos.LoginRequestDto
    if err := decodeAndValidateRequest(w, r, &req); err != nil {
        return
    }

    // Authenticate user through service layer
    token, err := LoginUserService(req.UsernameOrEmail, req.Password)
    if err != nil {
        handleServiceError(w, err)
        return
    }

    // Return success response with token
    w.WriteHeader(http.StatusOK)
    shared.Success(w, "Login successful", map[string]string{
        "token": token,
    })
}
```

## Benefits

1. **Improved Readability**: Code is easier to understand and maintain
2. **Reduced Duplication**: Common patterns are extracted into reusable functions
3. **Consistent Error Handling**: Errors are handled uniformly across controllers
4. **Better Separation of Concerns**: HTTP handling is separated from business logic
5. **Easier Testing**: Smaller, focused functions are easier to test
