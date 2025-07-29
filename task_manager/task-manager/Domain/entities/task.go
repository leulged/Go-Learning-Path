package entities

import "time"

// Task is the core domain entity - pure business logic
type Task struct {
	ID          string
	Title       string
	Description string
	DueDate     time.Time
	Status      string
}

// NewTask creates a new task with validation
func NewTask(title, description string, dueDate time.Time) Task {
	return Task{
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Status:      "Pending", // default status
	}
}

// IsCompleted checks if task is completed
func (t Task) IsCompleted() bool {
	return t.Status == "Completed"
}

// IsOverdue checks if task is overdue
func (t Task) IsOverdue() bool {
	return time.Now().After(t.DueDate) && !t.IsCompleted()
}

// SetStatus sets the task status
func (t *Task) SetStatus(status string) {
	t.Status = status
}

// ValidStatuses returns valid task statuses
func ValidStatuses() []string {
	return []string{"Pending", "In Progress", "Completed"}
} 