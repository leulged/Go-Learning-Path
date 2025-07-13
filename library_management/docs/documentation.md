
# Console-Based Library Management System

## Overview

This is a simple console-based Library Management System implemented in Go. It demonstrates core Go concepts including structs, interfaces, methods, slices, and maps. The system supports adding, removing, borrowing, and returning books, as well as listing available and borrowed books.

---

## Features

* Add new books to the library
* Remove existing books
* Borrow books by members
* Return borrowed books
* List all available books
* List all books borrowed by a specific member

---

## Folder Structure

```
library_management/
├── main.go                      # Application entry point
├── controllers/
│   └── library_controller.go    # Handles user input and invokes service methods
├── models/
│   ├── book.go                  # Book struct definition
│   └── member.go                # Member struct definition
├── services/
│   └── library_service.go       # Business logic and data management
├── docs/
│   └── documentation.md         # Additional documentation (optional)
├── go.mod                       # Module definition
└── README.md                    # Project documentation
```

---

## Getting Started

### Prerequisites

* [Go](https://golang.org/dl/) 1.18 or higher installed
* Basic understanding of Go programming



## How It Works

* **Models** define the core data structures: `Book` and `Member`.
* **Services** contain the core business logic managing books and members.
* **Controllers** manage user input/output and communicate with services.
* Data is stored in memory using maps and slices for simplicity.

