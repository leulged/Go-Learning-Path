# Task Management API

A comprehensive REST API for task management built with Go, following Clean Architecture principles. Features user authentication, role-based authorization, and full CRUD operations for tasks.

## 🏗️ Architecture

This project follows **Clean Architecture** with the following layers:

- **Domain**: Core entities and interfaces
- **Use Cases**: Business logic and application services
- **Delivery**: HTTP controllers and routing
- **Infrastructure**: Database connections and external services
- **Repositories**: Data access layer

## 🚀 Features

### 🔐 Authentication & Authorization

- **JWT-based authentication** with 24-hour token expiration
- **Role-based access control** (Admin/User roles)
- **Secure password hashing** using bcrypt
- **Automatic admin assignment** for the first registered user

### 👥 User Management

- User registration with email validation
- User login with JWT token generation
- Admin user promotion functionality
- Role-based middleware protection

### 📋 Task Management

- **Full CRUD operations** for tasks
- **Custom task IDs** with MongoDB ObjectID
- **Task status tracking** (Pending, In Progress, Completed)
- **Due date management** with ISO 8601 format
- **Admin-only task creation/modification**
- **Public task viewing** for all authenticated users

### 🛡️ Security Features

- **JWT middleware** for route protection
- **Admin middleware** for privileged operations
- **Password encryption** using bcrypt
- **Input validation** and error handling

### 📄 Database

- **MongoDB** integration with official driver
- **Environment-based configuration** with .env support
- **Connection pooling** and proper resource management

## 📁 Project Structure

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

## 🛠️ Technology Stack

- **Language**: Go 1.24.5
- **Web Framework**: Gin
- **Database**: MongoDB
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: bcrypt
- **Testing**: Testify
- **Environment**: godotenv
- **Architecture**: Clean Architecture

## 🚀 Getting Started

### Prerequisites

- Go 1.24.5 or higher
- MongoDB instance
- Git

### Installation

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd task-manager
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Set up environment variables**
   Create a `.env` file in the root directory:

   ```env
   MONGODB_URI=mongodb://localhost:27017
   DATABASE_NAME=task_management_system
   PORT=8080
   JWT_SECRET=your_secure_jwt_secret_key
   ```

4. **Run the application**
   ```bash
   go run Delivery/main.go
   ```

The server will start on `http://localhost:8080`

## 📚 API Documentation

See `docs/api_documentation.md` for full API details.

## 🧪 Testing

### Test Structure

- **Use Cases**: `Usecases/user_usecase_test.go`, `Usecases/task_usecase_test.go`
- **Controllers**: `Delivery/http/controllers/user_controller_test.go`, `Delivery/http/controllers/task_controller_test.go`
- **Repositories**: `Infrastructure/database/repositories/user_repository_test.go`, `Infrastructure/database/repositories/task_repository_test.go`
- **Utilities**: `utils/validation_test.go`, `utils/hash_test.go`
- **Integration**: `tests/integration_test.go`

### Running Tests

```bash
# Run all tests
 go test ./... -v

# Run with coverage
 go test -cover ./...

# Run specific packages
 go test ./Usecases -v
 go test ./Delivery/http/controllers -v
 go test ./Infrastructure/database/repositories -v
 go test ./utils -v
 go test ./tests -v

# Run tests with coverage report
 go test -coverprofile=coverage.out ./...
 go tool cover -html=coverage.out
```

### Test Coverage Status

- ✅ **Use Cases**: Comprehensive coverage with mocking
- ✅ **Controllers**: High coverage with HTTP testing
- ✅ **Repositories**: Integration tests with MongoDB
- ✅ **Utilities**: High coverage for helpers
- ❌ **Domain**: No tests needed (entities/interfaces only)

### Test Categories

1. **Unit Tests**: Use cases and utilities with dependency mocking
2. **Integration Tests**: Controllers and end-to-end HTTP request testing
3. **Repository Tests**: Database operations with test helpers
4. **Utility Tests**: Validation, hashing, and helpers

### Testing Best Practices

- **Test independence** and isolation
- **Comprehensive edge case** coverage
- **Proper mocking** with testify/mock
- **Clear test naming** and assertions
- **Database cleanup** after each test

## 🔧 Configuration

### Environment Variables

| Variable        | Description               | Default                     |
| --------------- | ------------------------- | --------------------------- |
| `MONGODB_URI`   | MongoDB connection string | `mongodb://localhost:27017` |
| `DATABASE_NAME` | Database name             | `task_management_system`    |
| `PORT`          | Server port               | `8080`                      |
| `JWT_SECRET`    | JWT signing secret        | `your_jwt_secret_key`       |

### Database Collections

- **users**: User accounts and authentication data
- **tasks**: Task management data

## 🔒 Security Features

### Authentication

- JWT tokens with 24-hour expiration
- Secure password hashing with bcrypt
- Role-based access control

### Authorization

- **Public routes**: Registration and login
- **Authenticated routes**: Task viewing
- **Admin-only routes**: Task management, user promotion

### Data Protection

- Password fields excluded from JSON responses
- Input validation and sanitization
- Error handling without sensitive data exposure

## 📊 API Response Examples

### Success Responses

```json
// User Registration
{
  "id": "507f1f77bcf86cd799439011",
  "name": "John Doe",
  "email": "john@example.com",
  "role": "admin"
}

// Login
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}

// Task List
{
  "tasks": [
    {
      "id": "507f1f77bcf86cd799439011",
      "title": "Complete Project",
      "description": "Finish the task management API",
      "due_date": "2025-07-25T00:00:00Z",
      "status": "Pending"
    }
  ]
}
```

### Error Responses

```json
// Authentication Error
{
  "error": "Invalid credentials"
}

// Authorization Error
{
  "error": "Admin access required"
}

// Validation Error
{
  "error": "Email already exists"
}
```

## 🚀 Deployment

### Local Development

```bash
 go run Delivery/main.go
```

### Production Build

```bash
 go build -o task-manager Delivery/main.go
 ./task-manager
```

### Docker (Optional)

```dockerfile
FROM golang:1.24-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o task-manager Delivery/main.go
EXPOSE 8080
CMD ["./task-manager"]
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## 📝 License

This project is licensed under the MIT License.

## 🆘 Support

For support and questions:

- Check the API documentation in `docs/api_documentation.md`
- Review the test files for usage examples
- Open an issue for bugs or feature requests
