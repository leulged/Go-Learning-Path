package controllers

import (
	"net/http"
	"strconv"
	"task_manager/models"
	"task_manager/services"

	"github.com/gin-gonic/gin"
)

// TaskController handles task-related HTTP requests
type TaskController struct {
	Service services.TaskService
}

// NewTaskController creates and returns a new TaskController instance
func NewTaskController(service services.TaskService) *TaskController {
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
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := tc.Service.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// AddTask handles POST /tasks
func (tc *TaskController) AddTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.Service.UpdateTask(id, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask handles DELETE /tasks/:id
func (tc *TaskController) DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	err = tc.Service.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
