# Task Management REST API Documentation

**Base URL:** `http://localhost:8080`

---

## MongoDB Configuration

- Database: `task_management_system`
- Collection: `task`
- Tasks use a **custom integer ID** instead of MongoDB's default `_id`.

---

## Endpoints

### 1. Get All Tasks

- **URL:** `/tasks`
- **Method:** `GET`
- **Description:** Retrieve all tasks.

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
````

---

### 2. Get Task by ID

* **URL:** `/tasks/{id}`
* **Method:** `GET`
* **Description:** Retrieve a task by its custom integer ID.

#### Path Parameter

* `id`: integer

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

* **URL:** `/tasks`
* **Method:** `POST`

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

* **URL:** `/tasks/{id}`
* **Method:** `PUT`

#### Path Parameter

* `id`: integer

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

* **URL:** `/tasks/{id}`
* **Method:** `DELETE`

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

## Notes

* All `id` values are custom-generated integers (not ObjectIDs).
* Dates must follow ISO 8601 format: `"YYYY-MM-DDTHH:MM:SSZ"`
* Status must be one of: `"Pending"`, `"In Progress"`, `"Completed"`

