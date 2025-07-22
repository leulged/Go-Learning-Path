package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	DueDate     time.Time          `bson:"due_date" json:"due_date"`
	Status      string             `bson:"status" json:"status"`
}

// TaskRepository interface defines task data access operations
type TaskRepository interface {
	GetTasks() []Task
	GetTaskByID(id string) (Task, error)
	AddTask(task Task) (Task, error)
	UpdateTask(id string, updatedTask Task) (Task, error)
	DeleteTask(id string) error
}
