package interfaces

import "task_manager/Domain/entities"

// UserRepository interface defines user data access operations
type UserRepository interface {
	GetUserByEmail(email string) (entities.User, error)
	CountDocuments(email string) (int64, error)
	InsertOne(user entities.User) (entities.User, error)
	UpdateOne(email string, user entities.User) (entities.User, error)
	UpdateRole(email, role string) error
}

// TaskRepository interface defines task data access operations
type TaskRepository interface {
	GetTasks() []entities.Task
	GetTaskByID(id string) (entities.Task, error)
	AddTask(task entities.Task) (entities.Task, error)
	UpdateTask(id string, updatedTask entities.Task) (entities.Task, error)
	DeleteTask(id string) error
} 