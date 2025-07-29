package controllers

import (
	"net/http"
	"task_manager/Domain/entities"
	"task_manager/Delivery/http/request"
	"task_manager/Delivery/http/response"
	"task_manager/Usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskController handles task-related HTTP requests
type TaskController struct {
	Service usecases.TaskUsecase
}

// NewTaskController creates and returns a new TaskController instance
func NewTaskController(service usecases.TaskUsecase) *TaskController {
	return &TaskController{
		Service: service,
	}
}

// GetTasks handles GET /tasks
func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks := tc.Service.GetTasks()
	
	// Return response DTO
	response := response.ToTaskListResponse(tasks)
	c.JSON(http.StatusOK, response)
}

// GetTaskByID handles GET /tasks/:id
func (tc *TaskController) GetTaskByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	
	task, err := tc.Service.GetTaskByID(id.Hex())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	
	// Return response DTO
	response := response.ToTaskResponse(task)
	c.JSON(http.StatusOK, response)
}

// AddTask handles POST /tasks
func (tc *TaskController) AddTask(c *gin.Context) {
	var input request.CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert CreateTaskInput to domain Task
	task := entities.NewTask(input.Title, input.Description, input.DueDate)
	task.SetStatus(input.Status)

	newTask, err := tc.Service.AddTask(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return response DTO
	response := response.ToTaskResponse(newTask)
	c.JSON(http.StatusCreated, response)
}

// UpdateTask handles PUT /tasks/:id
func (tc *TaskController) UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var input request.UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert UpdateTaskInput to domain Task
	updatedTask := entities.NewTask(input.Title, input.Description, input.DueDate)
	updatedTask.SetStatus(input.Status)

	task, err := tc.Service.UpdateTask(id.Hex(), updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Return response DTO
	response := response.ToTaskResponse(task)
	c.JSON(http.StatusOK, response)
}

// DeleteTask handles DELETE /tasks/:id
func (tc *TaskController) DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	
	err = tc.Service.DeleteTask(id.Hex())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	
	// Return response DTO
	response := response.ToMessageResponse("Task deleted successfully")
	c.JSON(http.StatusOK, response)
} 