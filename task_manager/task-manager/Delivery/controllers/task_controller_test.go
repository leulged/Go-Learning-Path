package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"task_manager/Delivery/controllers"
	"task_manager/Delivery/http/request"
	"task_manager/domain"
	"task_manager/mocks"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAddTask(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskUsecase := mocks.NewMockTaskRepository(ctrl)
	taskController := controllers.NewTaskController(mockTaskUsecase)

	// Test cases
	t.Run("Success - Valid Task Creation", func(t *testing.T) {
		// Arrange
		input := request.CreateTaskInput{
			Title:       "Test Task",
			Description: "Test Description",
			DueDate:     time.Now(),
			Status:      "pending",
		}

		expectedTask := domain.Task{
			ID:          primitive.NewObjectID(),
			Title:       input.Title,
			Description: input.Description,
			DueDate:     input.DueDate,
			Status:      input.Status,
		}

		mockTaskUsecase.EXPECT().AddTask(gomock.Any()).Return(expectedTask, nil)

		// Create HTTP request
		jsonData, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.POST("/tasks", taskController.AddTask)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusCreated, w.Code)

		var response domain.Task
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, expectedTask.Title, response.Title)
		assert.Equal(t, expectedTask.Description, response.Description)
		assert.Equal(t, expectedTask.Status, response.Status)
	})

	t.Run("Error - Invalid JSON", func(t *testing.T) {
		// Arrange
		invalidJSON := `{"title": "Test", "description": "Test", "invalid": json}`

		// Create HTTP request
		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBufferString(invalidJSON))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.POST("/tasks", taskController.AddTask)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["error"], "invalid")
	})

	t.Run("Error - Missing Required Field", func(t *testing.T) {
		// Arrange
		input := request.CreateTaskInput{
			Description: "Test Description",
			DueDate:     time.Now(),
			Status:      "pending",
			// Title is missing (required field)
		}

		// Create HTTP request
		jsonData, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.POST("/tasks", taskController.AddTask)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["error"], "required")
	})

	t.Run("Error - Use Case Returns Error", func(t *testing.T) {
		// Arrange
		input := request.CreateTaskInput{
			Title:       "Test Task",
			Description: "Test Description",
			DueDate:     time.Now(),
			Status:      "pending",
		}

		mockTaskUsecase.EXPECT().AddTask(gomock.Any()).Return(domain.Task{}, assert.AnError)

		// Create HTTP request
		jsonData, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.POST("/tasks", taskController.AddTask)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["error"], "error")
	})
}

func TestGetTasks(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskUsecase := mocks.NewMockTaskRepository(ctrl)
	taskController := controllers.NewTaskController(mockTaskUsecase)

	t.Run("Success - Get All Tasks", func(t *testing.T) {
		// Arrange
		expectedTasks := []domain.Task{
			{
				ID:          primitive.NewObjectID(),
				Title:       "Task 1",
				Description: "Description 1",
				DueDate:     time.Now(),
				Status:      "pending",
			},
			{
				ID:          primitive.NewObjectID(),
				Title:       "Task 2",
				Description: "Description 2",
				DueDate:     time.Now().Add(24 * time.Hour),
				Status:      "completed",
			},
		}

		mockTaskUsecase.EXPECT().GetTasks().Return(expectedTasks)

		// Create HTTP request
		req, _ := http.NewRequest("GET", "/tasks", nil)

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.GET("/tasks", taskController.GetTasks)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string][]domain.Task
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response["tasks"], 2)
		assert.Equal(t, expectedTasks[0].Title, response["tasks"][0].Title)
		assert.Equal(t, expectedTasks[1].Title, response["tasks"][1].Title)
	})

	t.Run("Success - Empty Tasks List", func(t *testing.T) {
		// Arrange
		mockTaskUsecase.EXPECT().GetTasks().Return([]domain.Task{})

		// Create HTTP request
		req, _ := http.NewRequest("GET", "/tasks", nil)

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.GET("/tasks", taskController.GetTasks)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string][]domain.Task
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response["tasks"], 0)
	})
}

func TestGetTaskByID(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskUsecase := mocks.NewMockTaskRepository(ctrl)
	taskController := controllers.NewTaskController(mockTaskUsecase)

	t.Run("Success - Get Task by Valid ID", func(t *testing.T) {
		// Arrange
		taskID := primitive.NewObjectID()
		expectedTask := domain.Task{
			ID:          taskID,
			Title:       "Test Task",
			Description: "Test Description",
			DueDate:     time.Now(),
			Status:      "pending",
		}

		mockTaskUsecase.EXPECT().GetTaskByID(taskID.Hex()).Return(expectedTask, nil)

		// Create HTTP request
		req, _ := http.NewRequest("GET", "/tasks/"+taskID.Hex(), nil)

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.GET("/tasks/:id", taskController.GetTaskByID)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response domain.Task
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, expectedTask.Title, response.Title)
		assert.Equal(t, expectedTask.Description, response.Description)
		assert.Equal(t, expectedTask.Status, response.Status)
	})

	t.Run("Error - Invalid Task ID", func(t *testing.T) {
		// Arrange
		invalidID := "invalid-id"

		// Create HTTP request
		req, _ := http.NewRequest("GET", "/tasks/"+invalidID, nil)

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.GET("/tasks/:id", taskController.GetTaskByID)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["error"], "Invalid task ID")
	})

	t.Run("Error - Task Not Found", func(t *testing.T) {
		// Arrange
		taskID := primitive.NewObjectID()
		mockTaskUsecase.EXPECT().GetTaskByID(taskID.Hex()).Return(domain.Task{}, assert.AnError)

		// Create HTTP request
		req, _ := http.NewRequest("GET", "/tasks/"+taskID.Hex(), nil)

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.GET("/tasks/:id", taskController.GetTaskByID)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusNotFound, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["error"], "error")
	})
}

func TestUpdateTask(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskUsecase := mocks.NewMockTaskRepository(ctrl)
	taskController := controllers.NewTaskController(mockTaskUsecase)

	t.Run("Success - Update Task", func(t *testing.T) {
		// Arrange
		taskID := primitive.NewObjectID()
		input := request.UpdateTaskInput{
			Title:       "Updated Task",
			Description: "Updated Description",
			DueDate:     time.Now().Add(24 * time.Hour),
			Status:      "completed",
		}

		expectedTask := domain.Task{
			ID:          taskID,
			Title:       input.Title,
			Description: input.Description,
			DueDate:     input.DueDate,
			Status:      input.Status,
		}

		mockTaskUsecase.EXPECT().UpdateTask(taskID.Hex(), gomock.Any()).Return(expectedTask, nil)

		// Create HTTP request
		jsonData, _ := json.Marshal(input)
		req, _ := http.NewRequest("PUT", "/tasks/"+taskID.Hex(), bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.PUT("/tasks/:id", taskController.UpdateTask)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response domain.Task
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, expectedTask.Title, response.Title)
		assert.Equal(t, expectedTask.Description, response.Description)
		assert.Equal(t, expectedTask.Status, response.Status)
	})

	t.Run("Error - Invalid Task ID", func(t *testing.T) {
		// Arrange
		invalidID := "invalid-id"
		input := request.UpdateTaskInput{
			Title: "Updated Task",
		}

		// Create HTTP request
		jsonData, _ := json.Marshal(input)
		req, _ := http.NewRequest("PUT", "/tasks/"+invalidID, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.PUT("/tasks/:id", taskController.UpdateTask)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["error"], "Invalid task ID")
	})

	t.Run("Error - Invalid JSON", func(t *testing.T) {
		// Arrange
		taskID := primitive.NewObjectID()
		invalidJSON := `{"title": "Updated", "invalid": json}`

		// Create HTTP request
		req, _ := http.NewRequest("PUT", "/tasks/"+taskID.Hex(), bytes.NewBufferString(invalidJSON))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.PUT("/tasks/:id", taskController.UpdateTask)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["error"], "invalid")
	})

	t.Run("Error - Task Not Found", func(t *testing.T) {
		// Arrange
		taskID := primitive.NewObjectID()
		input := request.UpdateTaskInput{
			Title: "Updated Task",
		}

		mockTaskUsecase.EXPECT().UpdateTask(taskID.Hex(), gomock.Any()).Return(domain.Task{}, assert.AnError)

		// Create HTTP request
		jsonData, _ := json.Marshal(input)
		req, _ := http.NewRequest("PUT", "/tasks/"+taskID.Hex(), bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.PUT("/tasks/:id", taskController.UpdateTask)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusNotFound, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["error"], "error")
	})
}

func TestDeleteTask(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskUsecase := mocks.NewMockTaskRepository(ctrl)
	taskController := controllers.NewTaskController(mockTaskUsecase)

	t.Run("Success - Delete Task", func(t *testing.T) {
		// Arrange
		taskID := primitive.NewObjectID()
		mockTaskUsecase.EXPECT().DeleteTask(taskID.Hex()).Return(nil)

		// Create HTTP request
		req, _ := http.NewRequest("DELETE", "/tasks/"+taskID.Hex(), nil)

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.DELETE("/tasks/:id", taskController.DeleteTask)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["message"], "Task deleted successfully")
	})

	t.Run("Error - Invalid Task ID", func(t *testing.T) {
		// Arrange
		invalidID := "invalid-id"

		// Create HTTP request
		req, _ := http.NewRequest("DELETE", "/tasks/"+invalidID, nil)

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.DELETE("/tasks/:id", taskController.DeleteTask)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["error"], "Invalid task ID")
	})

	t.Run("Error - Task Not Found", func(t *testing.T) {
		// Arrange
		taskID := primitive.NewObjectID()
		mockTaskUsecase.EXPECT().DeleteTask(taskID.Hex()).Return(assert.AnError)

		// Create HTTP request
		req, _ := http.NewRequest("DELETE", "/tasks/"+taskID.Hex(), nil)

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.DELETE("/tasks/:id", taskController.DeleteTask)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusNotFound, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["error"], "error")
	})
} 