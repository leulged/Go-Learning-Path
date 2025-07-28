package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"task_manager/Delivery/controllers"
	"task_manager/Delivery/http/request"
	"task_manager/domain"
	"task_manager/mocks"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestRegister(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUsecase := mocks.NewMockUserUsecase(ctrl)
	userController := controllers.NewUserController(mockUserUsecase)

	t.Run("Success - Valid User Registration", func(t *testing.T) {
		// Arrange
		input := request.RegisterInput{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password123",
		}

		expectedUser := domain.User{
			ID:       primitive.NewObjectID(),
			Name:     input.Name,
			Email:    input.Email,
			Password: "", // Password should not be returned
			Role:     "user",
		}

		mockUserUsecase.EXPECT().Register(gomock.Any()).Return(expectedUser, nil)

		// Create HTTP request
		jsonData, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.POST("/register", userController.Register)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusCreated, w.Code)

		var response domain.User
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, expectedUser.Name, response.Name)
		assert.Equal(t, expectedUser.Email, response.Email)
		assert.Equal(t, expectedUser.Role, response.Role)
		assert.Empty(t, response.Password) // Password should not be returned
	})

	t.Run("Error - Invalid JSON", func(t *testing.T) {
		// Arrange
		invalidJSON := `{"name": "John", "email": "john@example.com", "invalid": json}`

		// Create HTTP request
		req, _ := http.NewRequest("POST", "/register", bytes.NewBufferString(invalidJSON))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.POST("/register", userController.Register)

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
		input := request.RegisterInput{
			Name: "John Doe",
			// Email is missing (required field)
		}

		// Create HTTP request
		jsonData, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.POST("/register", userController.Register)

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
		input := request.RegisterInput{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password123",
		}

		mockUserUsecase.EXPECT().Register(gomock.Any()).Return(domain.User{}, assert.AnError)

		// Create HTTP request
		jsonData, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.POST("/register", userController.Register)

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

func TestLogin(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUsecase := mocks.NewMockUserUsecase(ctrl)
	userController := controllers.NewUserController(mockUserUsecase)

	t.Run("Success - Valid Login", func(t *testing.T) {
		// Arrange
		input := request.LoginInput{
			Email:    "john@example.com",
			Password: "password123",
		}

		expectedToken := "jwt-token-here"

		mockUserUsecase.EXPECT().Login(input.Email, input.Password).Return(expectedToken, nil)

		// Create HTTP request
		jsonData, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.POST("/login", userController.Login)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, true, response["success"])
		assert.Equal(t, expectedToken, response["token"])
	})

	t.Run("Error - Invalid JSON", func(t *testing.T) {
		// Arrange
		invalidJSON := `{"email": "john@example.com", "invalid": json}`

		// Create HTTP request
		req, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(invalidJSON))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.POST("/login", userController.Login)

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
		input := request.LoginInput{
			Email: "john@example.com",
			// Password is missing (required field)
		}

		// Create HTTP request
		jsonData, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.POST("/login", userController.Login)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["error"], "required")
	})

	t.Run("Error - Invalid Credentials", func(t *testing.T) {
		// Arrange
		input := request.LoginInput{
			Email:    "john@example.com",
			Password: "wrongpassword",
		}

		mockUserUsecase.EXPECT().Login(input.Email, input.Password).Return("", assert.AnError)

		// Create HTTP request
		jsonData, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.POST("/login", userController.Login)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusUnauthorized, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["error"], "error")
	})
}

func TestPromote(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUsecase := mocks.NewMockUserUsecase(ctrl)
	userController := controllers.NewUserController(mockUserUsecase)

	t.Run("Success - Promote User to Admin", func(t *testing.T) {
		// Arrange
		input := request.PromoteInput{
			Email: "john@example.com",
		}

		mockUserUsecase.EXPECT().PromoteToAdmin(input.Email).Return(nil)

		// Create HTTP request
		jsonData, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/users/promote", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.POST("/users/promote", userController.Promote)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["message"], "User promoted to admin")
	})

	t.Run("Error - Invalid JSON", func(t *testing.T) {
		// Arrange
		invalidJSON := `{"email": "john@example.com", "invalid": json}`

		// Create HTTP request
		req, _ := http.NewRequest("POST", "/users/promote", bytes.NewBufferString(invalidJSON))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.POST("/users/promote", userController.Promote)

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
		input := request.PromoteInput{
			// Email is missing (required field)
		}

		// Create HTTP request
		jsonData, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/users/promote", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.POST("/users/promote", userController.Promote)

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
		input := request.PromoteInput{
			Email: "john@example.com",
		}

		mockUserUsecase.EXPECT().PromoteToAdmin(input.Email).Return(assert.AnError)

		// Create HTTP request
		jsonData, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/users/promote", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Setup router
		router := gin.New()
		router.POST("/users/promote", userController.Promote)

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