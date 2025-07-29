package response

import (
	"task_manager/Domain/entities"
	"time"
)

// TaskResponse represents the task data sent in HTTP responses
type TaskResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

// ToTaskResponse converts domain Task to TaskResponse
func ToTaskResponse(task entities.Task) TaskResponse {
	return TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
		Status:      task.Status,
	}
}

// TaskListResponse represents a list of tasks
type TaskListResponse struct {
	Tasks []TaskResponse `json:"tasks"`
}

// ToTaskListResponse converts domain tasks to TaskListResponse
func ToTaskListResponse(tasks []entities.Task) TaskListResponse {
	var taskResponses []TaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, ToTaskResponse(task))
	}
	
	return TaskListResponse{
		Tasks: taskResponses,
	}
} 