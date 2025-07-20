# Task Management REST API Documentation

**Base URL:** `http://localhost:8080`

---

## Authentication & Authorization

This API uses JWT (JSON Web Tokens) for authentication. All protected endpoints require a valid JWT token in the Authorization header.

### Authentication Flow

1. Register a new user account
2. Login to receive a JWT token
3. Include the token in subsequent requests

### Token Format

```
Authorization: Bearer <your_jwt_token>
```

---

## MongoDB Configuration

- Database: `task_management_system`
- Collections: `users` and `tasks`
- Tasks use a **custom integer ID** instead of MongoDB's default `_id`.

---

## Authentication Endpoints

### 1. User Registration

- **URL:** `/register`
- **Method:** `POST`
- **Description:** Create a new user account. The first user registered becomes an admin automatically.

#### Request Body

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

#### Success Response

```json
{
  "id": "507f1f77bcf86cd799439011",
  "name": "John Doe",
  "email": "john@example.com",
  "role": "admin"
}
```

#### Error Response

```json
{
  "error": "Email already exists"
}
```

---

### 2. User Login

- **URL:** `/login`
- **Method:** `POST`
- **Description:** Authenticate user and receive JWT token.

#### Request Body

```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

#### Success Response

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### Error Response

```json
{
  "error": "Invalid credentials"
}
```

---

### 3. Promote User to Admin

- **URL:** `/user/promote`
- **Method:** `POST`
- **Authentication:** Required (Admin only)
- **Description:** Promote a regular user to admin role.

#### Request Body

```json
{
  "email": "user@example.com"
}
```

#### Success Response

```json
{
  "message": "User promoted to admin"
}
```

#### Error Response

```json
{
  "error": "User not found"
}
```

---

## Task Management Endpoints

### 1. Get All Tasks

- **URL:** `/task`
- **Method:** `GET`
- **Authentication:** Required
- **Description:** Retrieve all tasks (accessible by all authenticated users).

#### Headers

```
Authorization: Bearer <your_jwt_token>
```

#### Success Response

```json
{
  "tasks": [
    {
      "id": 1,
      "title": "Task 1",
      "description": "Do homework",
      "due_date": "2025-07-20T00:00:00Z",
      "status": "Pending"
    },
    ...
  ]
}
```

---

### 2. Get Task by ID

- **URL:** `/task/{id}`
- **Method:** `GET`
- **Authentication:** Required
- **Description:** Retrieve a task by its custom integer ID (accessible by all authenticated users).

#### Path Parameter

- `id`: integer

#### Headers

```
Authorization: Bearer <your_jwt_token>
```

#### Success Response

```json
{
  "id": 1,
  "title": "Task 1",
  "description": "Do homework",
  "due_date": "2025-07-20T00:00:00Z",
  "status": "Pending"
}
```

#### Error Response

```json
{
  "error": "Task not found"
}
```

---

### 3. Create a New Task

- **URL:** `/task`
- **Method:** `POST`
- **Authentication:** Required (Admin only)
- **Description:** Create a new task (admin only).

#### Headers

```
Authorization: Bearer <your_jwt_token>
```

#### Request Body

```json
{
  "title": "New Task",
  "description": "Something to do",
  "due_date": "2025-07-25T00:00:00Z",
  "status": "Pending"
}
```

> `id` is **automatically generated** by the system (you do not send it in the request).

#### Success Response

```json
{
  "id": 5,
  "title": "New Task",
  "description": "Something to do",
  "due_date": "2025-07-25T00:00:00Z",
  "status": "Pending"
}
```

---

### 4. Update Task by ID

- **URL:** `/task/{id}`
- **Method:** `PUT`
- **Authentication:** Required (Admin only)
- **Description:** Update an existing task (admin only).

#### Path Parameter

- `id`: integer

#### Headers

```
Authorization: Bearer <your_jwt_token>
```

#### Request Body

```json
{
  "title": "Updated Task",
  "description": "Updated work",
  "due_date": "2025-07-30T00:00:00Z",
  "status": "Completed"
}
```

#### Success Response

```json
{
  "id": 5,
  "title": "Updated Task",
  "description": "Updated work",
  "due_date": "2025-07-30T00:00:00Z",
  "status": "Completed"
}
```

---

### 5. Delete Task by ID

- **URL:** `/task/{id}`
- **Method:** `DELETE`
- **Authentication:** Required (Admin only)
- **Description:** Delete a task (admin only).

#### Path Parameter

- `id`: integer

#### Headers

```
Authorization: Bearer <your_jwt_token>
```

#### Success Response

```json
{
  "message": "Task deleted successfully"
}
```

#### Error Response

```json
{
  "error": "Task not found"
}
```

---

## Error Responses

### Authentication Errors

```json
{
  "error": "Authorization header missing or invalid"
}
```

```json
{
  "error": "Invalid or expired token"
}
```

### Authorization Errors

```json
{
  "error": "Admin access required"
}
```

---

## Usage Examples

### 1. Register and Login Flow

```bash
# 1. Register a new user
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","password":"password123"}'

# 2. Login to get token
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"password123"}'

# 3. Use token for protected endpoints
curl -X GET http://localhost:8080/task \
  -H "Authorization: Bearer <your_token_here>"
```

### 2. Admin Operations

```bash
# Create a task (admin only)
curl -X POST http://localhost:8080/task \
  -H "Authorization: Bearer <admin_token>" \
  -H "Content-Type: application/json" \
  -d '{"title":"New Task","description":"Description","due_date":"2025-07-25T00:00:00Z","status":"Pending"}'

# Promote a user to admin
curl -X POST http://localhost:8080/user/promote \
  -H "Authorization: Bearer <admin_token>" \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com"}'
```

---

## Notes

- All `id` values are custom-generated integers (not ObjectIDs).
- Dates must follow ISO 8601 format: `"YYYY-MM-DDTHH:MM:SSZ"`
- Status must be one of: `"Pending"`, `"In Progress"`, `"Completed"`
- JWT tokens expire after 24 hours
- The first user to register automatically becomes an admin
- Only admins can create, update, and delete tasks
- All authenticated users can view tasks
- Passwords are securely hashed using bcrypt
