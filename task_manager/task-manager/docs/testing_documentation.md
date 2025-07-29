# Testing Documentation

## Overview

This document provides comprehensive testing documentation for the Task Management API. The project follows a layered testing approach with unit tests, integration tests, repository tests, and utility tests using the `testify` library.

## 🧪 Test Architecture

### Test Structure Overview

```
task-manager/
├── Domain/
│   ├── entities/                # Core entities (user.go, task.go)
│   ├── errors/                  # Domain error types
│   └── interfaces/              # Repository and service interfaces
├── Usecases/
│   ├── user_usecase.go          # User business logic
│   ├── task_usecase.go          # Task business logic
│   ├── user_usecase_test.go     # User use case unit tests
│   └── task_usecase_test.go     # Task use case unit tests
├── Delivery/
│   └── http/
│       ├── controllers/
│       │   ├── user_controller.go
│       │   ├── task_controller.go
│       │   ├── user_controller_test.go   # User controller integration tests
│       │   └── task_controller_test.go   # Task controller integration tests
│       ├── middleware/          # HTTP middleware
│       ├── request/             # Request DTOs
│       ├── response/            # Response DTOs
│       └── routers/             # Route definitions
├── Infrastructure/
│   ├── database/
│   │   └── repositories/
│   │       ├── user_repository.go
│   │       ├── task_repository.go
│   │       ├── user_repository_test.go   # User repository integration tests
│   │       └── task_repository_test.go   # Task repository integration tests
│   └── services/                # JWT and other services
├── utils/
│   ├── validation.go
│   ├── hash.go
│   ├── validation_test.go       # Utility tests
│   └── hash_test.go             # Utility tests
├── tests/
│   └── integration_test.go      # End-to-end integration tests
├── mocks/                       # Generated mocks for testing
├── docs/                        # Documentation
└── main.go                      # Application entry point
```

## 📊 Test Coverage Status

| Layer            | Status          | Coverage | Test Type                  |
| ---------------- | --------------- | -------- | -------------------------- |
| **Use Cases**    | ✅ Complete     | High     | Unit Tests with Mocking    |
| **Controllers**  | ✅ Complete     | High     | Integration Tests          |
| **Repositories** | ✅ Complete     | High     | Database Integration Tests |
| **Utils**        | ✅ Complete     | High     | Utility Unit Tests         |
| **Domain**       | ❌ Not Required | N/A      | Entities/Interfaces Only   |

## 🚀 Running Tests

### Basic Test Commands

```bash
# Run all tests in the project
 go test ./... -v

# Run tests with coverage
 go test -cover ./...

# Run tests with coverage report
 go test -coverprofile=coverage.out ./...
 go tool cover -html=coverage.out

# Run tests with race condition detection
 go test -race ./...
```

### Package-Specific Testing

```bash
# Test Use Cases layer
 go test ./Usecases -v

# Test Controllers layer
 go test ./Delivery/http/controllers -v

# Test Repositories layer
 go test ./Infrastructure/database/repositories -v

# Test Utilities
 go test ./utils -v

# Test integration tests
 go test ./tests -v

# Run a specific test file
 go test ./Usecases/user_usecase_test.go -v
```

### Advanced Testing Options

```bash
# Run tests with verbose output and coverage
 go test -v -cover ./...

# Run tests with timeout
 go test -timeout 30s ./...

# Run tests with specific tags
 go test -tags=integration ./...

# Generate coverage report in HTML
 go test -coverprofile=coverage.out ./...
 go tool cover -html=coverage.out -o coverage.html
```

## 📋 Test Categories

### 1. Unit Tests (Use Cases, Utils)

**Location**: `Usecases/user_usecase_test.go`, `Usecases/task_usecase_test.go`, `utils/validation_test.go`, `utils/hash_test.go`

**Purpose**: Test business logic and utility functions in isolation using mocked dependencies.

### 2. Integration Tests (Controllers, End-to-End)

**Location**: `Delivery/http/controllers/user_controller_test.go`, `Delivery/http/controllers/task_controller_test.go`, `tests/integration_test.go`

**Purpose**: Test HTTP endpoints and full request/response cycles, including middleware and authentication.

### 3. Repository Tests (Database)

**Location**: `Infrastructure/database/repositories/user_repository_test.go`, `Infrastructure/database/repositories/task_repository_test.go`

**Purpose**: Test database operations with real MongoDB connection.

### 4. Utility Tests

**Location**: `utils/validation_test.go`, `utils/hash_test.go`

**Purpose**: Test utility functions for correctness and edge cases.

## 🛠️ Test Configuration

### Test Environment Setup

```bash
# Required environment variables for testing
export MONGODB_URI=mongodb://localhost:27017
export DATABASE_NAME=task_management_test
export JWT_SECRET=test_jwt_secret
export PORT=8081
```

### Test Database Configuration

- Repository tests and integration tests use a dedicated test database.
- Clean up test data before and after each test to ensure independence.

## 🧑‍💻 Test Best Practices

- Each test should be independent and not rely on other tests
- Use `t.Parallel()` for concurrent test execution where appropriate
- Clean up test data after each test
- Use descriptive test names
- Follow the Arrange-Act-Assert (AAA) pattern
- Use mocks for dependencies in unit tests
- Cover edge cases and error scenarios

## 🐛 Common Test Issues and Solutions

- **Database Connection Issues**: Ensure MongoDB is running and environment variables are set.
- **JWT Token Issues**: Use a consistent JWT secret for tests.
- **Concurrent Test Issues**: Use unique database names or collections for each test if needed.

## 📊 Test Metrics and Reporting

- Generate coverage report:
  ```bash
  go test -coverprofile=coverage.out ./...
  go tool cover -func=coverage.out
  go tool cover -html=coverage.out -o coverage.html
  ```
- Run tests with timing and memory profiling:
  ```bash
  go test -v -timeout 30s ./...
  go test -memprofile=mem.out ./...
  go tool pprof mem.out
  ```

## 🔍 Test Debugging

- Run specific test with verbose output:
  ```bash
  go test -v -run TestSpecificFunction
  ```
- Add logging to tests for debugging:
  ```go
  t.Log("Starting test...")
  // ...
  t.Log("Test completed successfully")
  ```

## 🚀 Continuous Integration Testing

- See `.github/workflows/test.yml` for GitHub Actions setup.
- Ensure MongoDB service is available in CI environment.

## 📚 Test Examples

- See test files in `Usecases/`, `Delivery/http/controllers/`, `Infrastructure/database/repositories/`, `utils/`, and `tests/` for real examples.

## 🤝 Contributing to Tests

When adding new features, ensure to:

1. **Write tests first** (TDD approach)
2. **Cover all edge cases**
3. **Test error scenarios**
4. **Maintain test independence**
5. **Update this documentation** if needed

## 📞 Support

For testing-related issues:

- Check the test examples in this document
- Review existing test files for patterns
- Ensure test environment is properly configured
- Verify MongoDB connection for repository tests
- Check JWT configuration for authentication tests
