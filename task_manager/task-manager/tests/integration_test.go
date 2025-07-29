package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"
	"task_manager/config"
	"task_manager/Delivery/http/controllers"
	"task_manager/Infrastructure/database/repositories"
	"task_manager/Infrastructure/services"
	usecases "task_manager/Usecases"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setupTestApp() *gin.Engine {
	// Load environment variables from .env file in project root
	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		// Try current directory as fallback
		err = godotenv.Load()
		if err != nil {
			panic("Error loading .env file: " + err.Error())
		}
	}

	// Connect to test database
	_ = config.ConnectToMongo()
	// Note: We don't disconnect here because the client needs to stay connected for the tests
	// The client will be cleaned up when the test process ends

	// Initialize repositories
	userRepo := repositories.NewUserRepository(config.UserCollection)
	taskRepo := repositories.NewTaskRepository(config.TaskCollection)

	// Initialize services
	tokenService := services.NewJWTService()

	// Initialize use cases
	userUsecase := usecases.NewUserUsecase(userRepo, tokenService)
	taskUsecase := usecases.NewTaskUsecase(taskRepo)

	// Initialize controllers
	userController := controllers.NewUserController(userUsecase)
	taskController := controllers.NewTaskController(taskUsecase)

	// Setup Gin router
	gin.SetMode(gin.TestMode)
	r := gin.New()
	
	// Setup routes
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
	r.GET("/tasks", taskController.GetTasks)
	r.POST("/tasks", taskController.AddTask)

	return r
}

func TestRegisterIntegration(t *testing.T) {
	app := setupTestApp()

	// Test data with unique email using timestamp
	timestamp := time.Now().Unix()
	uniqueEmail := fmt.Sprintf("test%d@example.com", timestamp)
	
	userData := map[string]interface{}{
		"name":     "Test User",
		"email":    uniqueEmail,
		"password": "password123",
	}

	jsonData, _ := json.Marshal(userData)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	// Debug: Print response
	fmt.Printf("Register Status: %d\n", w.Code)
	fmt.Printf("Register Body: %s\n", w.Body.String())

	assert.Equal(t, http.StatusCreated, w.Code) // 201 for successful creation

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, uniqueEmail, response["email"])
	assert.Equal(t, "Test User", response["name"])
	assert.Equal(t, "admin", response["role"]) // First user becomes admin
}

func TestLoginIntegration(t *testing.T) {
	app := setupTestApp()

	// First register a user with unique email
	timestamp := time.Now().Unix()
	uniqueEmail := fmt.Sprintf("test%d@example.com", timestamp)
	
	userData := map[string]interface{}{
		"name":     "Test User",
		"email":    uniqueEmail,
		"password": "password123",
	}

	jsonData, _ := json.Marshal(userData)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	// Now test login
	loginData := map[string]interface{}{
		"email":    uniqueEmail,
		"password": "password123",
	}

	jsonData, _ = json.Marshal(loginData)
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.True(t, response["success"].(bool))
	assert.NotEmpty(t, response["token"])
}

func TestGetTasksIntegration(t *testing.T) {
	app := setupTestApp()

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response["tasks"])
} 