package services

import (
	"fmt"
	"task_manager/models"
)

// TaskService defines the interface for task operations.
type TaskService interface {
	GetTasks() []models.Task
	GetTaskByID(id int) (models.Task, error)
	AddTask(task models.Task) (models.Task, error)
	UpdateTask(id int, updatedTask models.Task) (models.Task, error)
	DeleteTask(id int) error
}

// taskServiceImpl implements TaskService using an in-memory map.
type taskServiceImpl struct {
	tasks map[int]models.Task
}

// NewTaskService initializes the service with an empty map.
func NewTaskService() TaskService {
	return &taskServiceImpl{
		tasks: make(map[int]models.Task),
	}
}

// ✅ GetTasks returns all tasks as a slice
func (ts *taskServiceImpl) GetTasks() []models.Task {
	result := []models.Task{}
	for _, task := range ts.tasks {
		result = append(result, task)
	}
	return result
}

// ✅ GetTaskByID finds a task by its ID or returns an error
func (ts *taskServiceImpl) GetTaskByID(id int) (models.Task, error) {
	task, exists := ts.tasks[id]
	if !exists {
		return models.Task{}, fmt.Errorf("task not found")
	}
	return task, nil
}

// ✅ AddTask adds a task to the map
func (ts *taskServiceImpl) AddTask(task models.Task) (models.Task, error) {
	ts.tasks[task.ID] = task // Make sure ID is int
	return task, nil
}

// ✅ UpdateTask updates fields if the task exists
func (ts *taskServiceImpl) UpdateTask(id int, updatedTask models.Task) (models.Task, error) {
	task, exists := ts.tasks[id]
	if !exists {
		return models.Task{}, fmt.Errorf("task not found")
	}

	// Apply updates
	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.Status = updatedTask.Status
	task.DueDate = updatedTask.DueDate

	ts.tasks[id] = task // Save back to map
	return task, nil
}

// ✅ DeleteTask removes task if it exists
func (ts *taskServiceImpl) DeleteTask(id int) error {
	_, exists := ts.tasks[id]
	if !exists {
		return fmt.Errorf("task not found")
	}
	delete(ts.tasks, id)
	return nil
}
