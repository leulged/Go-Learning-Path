# Task Management REST API Documentation

**Base URL:** `http://localhost:8080`

---

## Endpoints

### 1. Get All Tasks

- **URL:** `/tasks`
- **Method:** `GET`
- **Description:** Retrieve all tasks in the system.

#### Success Response
- **Status:** `200 OK`
```json
{
  "tasks": [
    {
      "id": 1,
      "title": "Task 1",
      "description": "First task description",
      "due_date": "2025-07-20T00:00:00Z",
      "status": "Pending"
    },
    ...
  ]
}
````

---

### 2. Get Task by ID

* **URL:** `/tasks/{id}`
* **Method:** `GET`
* **Description:** Get details of a specific task by its ID.

#### URL Parameter

* `id` (integer): The task ID

#### Success Response

* **Status:** `200 OK`

```json
{
  "id": 1,
  "title": "Task 1",
  "description": "First task description",
  "due_date": "2025-07-20T00:00:00Z",
  "status": "Pending"
}
```

#### Error Response

* **Status:** `404 Not Found`

```json
{
  "error": "Task not found"
}
```

---

### 3. Create a New Task

* **URL:** `/tasks`
* **Method:** `POST`
* **Description:** Create a new task.

#### Request Body

```json
{
  "title": "New Task",
  "description": "Task description here",
  "due_date": "2025-07-25T00:00:00Z",
  "status": "Pending"
}
```

#### Success Response

* **Status:** `201 Created`

```json
{
  "id": 3,
  "title": "New Task",
  "description": "Task description here",
  "due_date": "2025-07-25T00:00:00Z",
  "status": "Pending"
}
```

#### Error Response

* **Status:** `400 Bad Request`

```json
{
  "error": "Invalid input data"
}
```

---

### 4. Update Task by ID

* **URL:** `/tasks/{id}`
* **Method:** `PUT`
* **Description:** Update an existing task by its ID.

#### URL Parameter

* `id` (integer): The task ID

#### Request Body

```json
{
  "title": "Updated Task Title",
  "description": "Updated description",
  "due_date": "2025-07-30T00:00:00Z",
  "status": "In Progress"
}
```

#### Success Response

* **Status:** `200 OK`

```json
{
  "id": 1,
  "title": "Updated Task Title",
  "description": "Updated description",
  "due_date": "2025-07-30T00:00:00Z",
  "status": "In Progress"
}
```

#### Error Responses

* **Status:** `400 Bad Request`

```json
{
  "error": "Invalid input data"
}
```

* **Status:** `404 Not Found`

```json
{
  "error": "Task not found"
}
```

---

### 5. Delete Task by ID

* **URL:** `/tasks/{id}`
* **Method:** `DELETE`
* **Description:** Delete a task by its ID.

#### URL Parameter

* `id` (integer): The task ID

#### Success Response

* **Status:** `200 OK`

```json
{
  "message": "Task deleted successfully"
}
```

#### Error Response

* **Status:** `404 Not Found`

```json
{
  "error": "Task not found"
}
```

---

## Error Handling

The API uses standard HTTP response codes:

| Code | Meaning               | Description                           |
| ---- | --------------------- | ------------------------------------- |
| 200  | OK                    | Request succeeded                     |
| 201  | Created               | Task was successfully created         |
| 400  | Bad Request           | Malformed input or validation failure |
| 404  | Not Found             | Task with given ID does not exist     |
| 500  | Internal Server Error | Unexpected error on the server        |

All error responses return a JSON object with an `"error"` key.

---

## Notes

* All dates must be in **ISO 8601 format**: `YYYY-MM-DDTHH:MM:SSZ`
* The `id` must be int
* The `status` field must be one of: `"Pending"`, `"In Progress"`, `"Completed"`
* All communication is in JSON format
* This API uses **in-memory storage** — data resets when the server restarts

---

## Example cURL Commands

### List all tasks

```bash
curl -X GET http://localhost:8080/tasks
```

### Get task by ID

```bash
curl -X GET http://localhost:8080/tasks/1
```

### Create a new task

```bash
curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{
  "title": "New Task",
  "description": "Task description",
  "due_date": "2025-07-25T00:00:00Z",
  "status": "Pending"
}'
```

### Update a task

```bash
curl -X PUT http://localhost:8080/tasks/1 \
-H "Content-Type: application/json" \
-d '{
  "title": "Updated Title",
  "description": "Updated description",
  "due_date": "2025-07-30T00:00:00Z",
  "status": "In Progress"
}'
```

### Delete a task

```bash
curl -X DELETE http://localhost:8080/tasks/1
```

---

