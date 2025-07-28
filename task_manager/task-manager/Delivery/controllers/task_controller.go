package controllers

import (
	"net/http"
	"task_manager/domain"
	usecases "task_manager/Usecases"
	"task_manager/Delivery/http/request"

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
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
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
	c.JSON(http.StatusOK, task)
}

// AddTask handles POST /tasks
func (tc *TaskController) AddTask(c *gin.Context) {
	var input request.CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := domain.Task{
		Title:       input.Title,
		Description: input.Description,
		DueDate:     input.DueDate,
		Status:      input.Status,
	}

	newTask, err := tc.Service.AddTask(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTask)
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

	updatedTask := domain.Task{
		Title:       input.Title,
		Description: input.Description,
		DueDate:     input.DueDate,
		Status:      input.Status,
	}

	task, err := tc.Service.UpdateTask(id.Hex(), updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
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
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
