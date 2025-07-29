package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"task_manager/Delivery/http/response"
	"task_manager/Domain/entities"
	"task_manager/Domain/errors"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock TaskUsecase
type MockTaskUsecase struct {
	mock.Mock
}

func (m *MockTaskUsecase) GetTasks() []entities.Task {
	args := m.Called()
	return args.Get(0).([]entities.Task)
}

func (m *MockTaskUsecase) GetTaskByID(id string) (entities.Task, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return entities.Task{}, args.Error(1)
	}
	return args.Get(0).(entities.Task), args.Error(1)
}

func (m *MockTaskUsecase) AddTask(task entities.Task) (entities.Task, error) {
	args := m.Called(task)
	if args.Get(0) == nil {
		return entities.Task{}, args.Error(1)
	}
	return args.Get(0).(entities.Task), args.Error(1)
}

func (m *MockTaskUsecase) UpdateTask(id string, updatedTask entities.Task) (entities.Task, error) {
	args := m.Called(id, updatedTask)
	if args.Get(0) == nil {
		return entities.Task{}, args.Error(1)
	}
	return args.Get(0).(entities.Task), args.Error(1)
}

func (m *MockTaskUsecase) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func setupTaskTestRouter(controller *TaskController) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/tasks", controller.AddTask)
	r.GET("/tasks", controller.GetTasks)
	r.GET("/tasks/:id", controller.GetTaskByID)
	r.PUT("/tasks/:id", controller.UpdateTask)
	r.DELETE("/tasks/:id", controller.DeleteTask)
	return r
}

func TestTaskController_AddTask_Success(t *testing.T) {
	// Setup
	mockUsecase := new(MockTaskUsecase)
	controller := NewTaskController(mockUsecase)
	router := setupTaskTestRouter(controller)

	// Mock data
	expectedTask := entities.NewTask("Test Task", "Test Description", time.Now())
	expectedTask.ID = "task123"

	// Mock expectations
	mockUsecase.On("AddTask", mock.AnythingOfType("entities.Task")).Return(expectedTask, nil)

	// Test data
	requestData := map[string]interface{}{
		"title":       "Test Task",
		"description": "Test Description",
		"status":      "Pending",
		"user_id":     "user123",
	}

	jsonData, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Test Task", response["title"])
	assert.Equal(t, "Test Description", response["description"])
	assert.Equal(t, "Pending", response["status"])
	assert.Equal(t, "task123", response["id"])

	mockUsecase.AssertExpectations(t)
}

func TestTaskController_AddTask_ValidationError(t *testing.T) {
	// Setup
	mockUsecase := new(MockTaskUsecase)
	controller := NewTaskController(mockUsecase)
	router := setupTaskTestRouter(controller)

	// Test data with invalid title (empty)
	requestData := map[string]interface{}{
		"title":       "",
		"description": "Test Description",
		"status":      "Pending",
		"user_id":     "user123",
	}

	jsonData, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "Key: 'CreateTaskInput.Title' Error:Field validation for 'Title' failed on the 'required' tag")
}

func TestTaskController_AddTask_InvalidStatus(t *testing.T) {
	// Setup
	mockUsecase := new(MockTaskUsecase)
	controller := NewTaskController(mockUsecase)
	router := setupTaskTestRouter(controller)

	// Mock data
	expectedTask := entities.NewTask("Test Task", "Test Description", time.Now())
	expectedTask.ID = "task123"
	expectedTask.SetStatus("Invalid Status")

	// Mock expectations
	mockUsecase.On("AddTask", mock.AnythingOfType("entities.Task")).Return(expectedTask, nil)

	// Test data with invalid status
	requestData := map[string]interface{}{
		"title":       "Test Task",
		"description": "Test Description",
		"status":      "Invalid Status",
		"user_id":     "user123",
	}

	jsonData, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code) // Status validation happens in use case, not binding

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Test Task", response["title"])
	assert.Equal(t, "Test Description", response["description"])
	assert.Equal(t, "Invalid Status", response["status"])

	mockUsecase.AssertExpectations(t)
}

func TestTaskController_GetTasks_Success(t *testing.T) {
	// Setup
	mockUsecase := new(MockTaskUsecase)
	controller := NewTaskController(mockUsecase)
	router := setupTaskTestRouter(controller)

	// Mock data
	task1 := entities.NewTask("Task 1", "Description 1", time.Now())
	task1.ID = "task1"
	task2 := entities.NewTask("Task 2", "Description 2", time.Now())
	task2.ID = "task2"

	expectedTasks := []entities.Task{task1, task2}

	// Mock expectations
	mockUsecase.On("GetTasks").Return(expectedTasks)

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response["tasks"])

	tasks := response["tasks"].([]interface{})
	assert.Len(t, tasks, 2)

	// Verify first task
	task1Data := tasks[0].(map[string]interface{})
	assert.Equal(t, "Task 1", task1Data["title"])
	assert.Equal(t, "Description 1", task1Data["description"])
	assert.Equal(t, "Pending", task1Data["status"])

	mockUsecase.AssertExpectations(t)
}

func TestTaskController_GetTaskByID_Success(t *testing.T) {
	// Setup
	mockUsecase := new(MockTaskUsecase)
	controller := NewTaskController(mockUsecase)
	router := setupTaskTestRouter(controller)

	// Mock data
	expectedTask := entities.NewTask("Test Task", "Test Description", time.Now())
	expectedTask.ID = "507f1f77bcf86cd799439011"

	// Mock expectations
	mockUsecase.On("GetTaskByID", "507f1f77bcf86cd799439011").Return(expectedTask, nil)

	req, _ := http.NewRequest("GET", "/tasks/507f1f77bcf86cd799439011", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Test Task", response["title"])
	assert.Equal(t, "Test Description", response["description"])
	assert.Equal(t, "Pending", response["status"])
	assert.Equal(t, "507f1f77bcf86cd799439011", response["id"])

	mockUsecase.AssertExpectations(t)
}

func TestTaskController_GetTaskByID_NotFound(t *testing.T) {
	// Setup
	mockUsecase := new(MockTaskUsecase)
	controller := NewTaskController(mockUsecase)
	router := setupTaskTestRouter(controller)

	// Mock expectations
	mockUsecase.On("GetTaskByID", "507f1f77bcf86cd799439012").Return(entities.Task{}, errors.TaskNotFoundError{})

	req, _ := http.NewRequest("GET", "/tasks/507f1f77bcf86cd799439012", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusNotFound, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "task not found")

	mockUsecase.AssertExpectations(t)
}

func TestTaskController_UpdateTask_Success(t *testing.T) {
	// Setup
	mockUsecase := new(MockTaskUsecase)
	controller := NewTaskController(mockUsecase)
	router := setupTaskTestRouter(controller)

	// Mock data
	expectedTask := entities.NewTask("Updated Task", "Updated Description", time.Now())
	expectedTask.SetStatus("Completed")
	expectedTask.ID = "507f1f77bcf86cd799439011"

	// Mock expectations
	mockUsecase.On("UpdateTask", "507f1f77bcf86cd799439011", mock.AnythingOfType("entities.Task")).Return(expectedTask, nil)

	// Test data
	requestData := map[string]interface{}{
		"title":       "Updated Task",
		"description": "Updated Description",
		"status":      "Completed",
	}

	jsonData, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("PUT", "/tasks/507f1f77bcf86cd799439011", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Task", response["title"])
	assert.Equal(t, "Updated Description", response["description"])
	assert.Equal(t, "Completed", response["status"])

	mockUsecase.AssertExpectations(t)
}

func TestTaskController_DeleteTask_Success(t *testing.T) {
	// Setup
	mockUsecase := new(MockTaskUsecase)
	controller := NewTaskController(mockUsecase)
	router := setupTaskTestRouter(controller)

	// Mock expectations
	mockUsecase.On("DeleteTask", "507f1f77bcf86cd799439011").Return(nil)

	req, _ := http.NewRequest("DELETE", "/tasks/507f1f77bcf86cd799439011", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Task deleted successfully", response["message"])

	mockUsecase.AssertExpectations(t)
}

func TestTaskResponse_ToMap(t *testing.T) {
	// Test task response mapping
	task := entities.NewTask("Test Task", "Test Description", time.Now())
	task.ID = "task123"

	taskResp := response.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}

	// Convert to map using JSON marshaling
	jsonData, err := json.Marshal(taskResp)
	assert.NoError(t, err)
	
	var taskMap map[string]interface{}
	err = json.Unmarshal(jsonData, &taskMap)
	assert.NoError(t, err)
	
	assert.Equal(t, "task123", taskMap["id"])
	assert.Equal(t, "Test Task", taskMap["title"])
	assert.Equal(t, "Test Description", taskMap["description"])
	assert.Equal(t, "Pending", taskMap["status"])
} 