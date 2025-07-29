package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"task_manager/Delivery/http/response"
	"task_manager/Domain/entities"
	"task_manager/Domain/errors"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock UserUsecase
type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Register(user entities.User) (entities.User, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return entities.User{}, args.Error(1)
	}
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockUserUsecase) Login(email, password string) (string, error) {
	args := m.Called(email, password)
	return args.String(0), args.Error(1)
}

func (m *MockUserUsecase) GetUserByEmail(email string) (entities.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return entities.User{}, args.Error(1)
	}
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockUserUsecase) PromoteToAdmin(email string) error {
	args := m.Called(email)
	return args.Error(0)
}

func setupTestRouter(controller *UserController) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	return r
}

func TestUserController_Register_Success(t *testing.T) {
	// Setup
	mockUsecase := new(MockUserUsecase)
	controller := NewUserController(mockUsecase)
	router := setupTestRouter(controller)

	// Mock data
	expectedUser := entities.NewUser("Test User", "test@example.com", "hashedpassword")
	expectedUser.SetRole("admin")
	expectedUser.ID = "user123"

	// Mock expectations
	mockUsecase.On("Register", mock.AnythingOfType("entities.User")).Return(expectedUser, nil)

	// Test data
	requestData := map[string]interface{}{
		"name":     "Test User",
		"email":    "test@example.com",
		"password": "password123",
	}

	jsonData, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", response["email"])
	assert.Equal(t, "Test User", response["name"])
	assert.Equal(t, "admin", response["role"])
	assert.Equal(t, "user123", response["id"])

	mockUsecase.AssertExpectations(t)
}

func TestUserController_Register_ValidationError(t *testing.T) {
	// Setup
	mockUsecase := new(MockUserUsecase)
	controller := NewUserController(mockUsecase)
	router := setupTestRouter(controller)

	// Test data with invalid email
	requestData := map[string]interface{}{
		"name":     "Test User",
		"email":    "invalid-email",
		"password": "password123",
	}

	jsonData, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "Field validation for 'Email' failed on the 'email' tag")
}

func TestUserController_Register_EmailAlreadyExists(t *testing.T) {
	// Setup
	mockUsecase := new(MockUserUsecase)
	controller := NewUserController(mockUsecase)
	router := setupTestRouter(controller)

	// Mock expectations
	mockUsecase.On("Register", mock.AnythingOfType("entities.User")).Return(entities.User{}, errors.EmailAlreadyExistsError{})

	// Test data
	requestData := map[string]interface{}{
		"name":     "Test User",
		"email":    "test@example.com",
		"password": "password123",
	}

	jsonData, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "email already exists")

	mockUsecase.AssertExpectations(t)
}

func TestUserController_Register_InvalidJSON(t *testing.T) {
	// Setup
	mockUsecase := new(MockUserUsecase)
	controller := NewUserController(mockUsecase)
	router := setupTestRouter(controller)

	// Test data with invalid JSON
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "invalid character 'i' looking for beginning of value")
}

func TestUserController_Login_Success(t *testing.T) {
	// Setup
	mockUsecase := new(MockUserUsecase)
	controller := NewUserController(mockUsecase)
	router := setupTestRouter(controller)

	// Mock expectations
	mockUsecase.On("Login", "test@example.com", "password123").Return("jwt-token-123", nil)

	// Test data
	requestData := map[string]interface{}{
		"email":    "test@example.com",
		"password": "password123",
	}

	jsonData, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.True(t, response["success"].(bool))
	assert.Equal(t, "jwt-token-123", response["token"])

	mockUsecase.AssertExpectations(t)
}

func TestUserController_Login_InvalidCredentials(t *testing.T) {
	// Setup
	mockUsecase := new(MockUserUsecase)
	controller := NewUserController(mockUsecase)
	router := setupTestRouter(controller)

	// Mock expectations
	mockUsecase.On("Login", "test@example.com", "wrongpassword").Return("", errors.InvalidCredentialsError{})

	// Test data
	requestData := map[string]interface{}{
		"email":    "test@example.com",
		"password": "wrongpassword",
	}

	jsonData, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "invalid credentials")

	mockUsecase.AssertExpectations(t)
}

func TestUserController_Login_InvalidJSON(t *testing.T) {
	// Setup
	mockUsecase := new(MockUserUsecase)
	controller := NewUserController(mockUsecase)
	router := setupTestRouter(controller)

	// Test data with invalid JSON
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "invalid character 'i' looking for beginning of value")
}

func TestUserResponse_ToMap(t *testing.T) {
	// Test user response mapping
	user := entities.NewUser("Test User", "test@example.com", "hashedpassword")
	user.SetRole("admin")
	user.ID = "user123"

	userResp := response.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	// Convert to map using JSON marshaling
	jsonData, err := json.Marshal(userResp)
	assert.NoError(t, err)
	
	var userMap map[string]interface{}
	err = json.Unmarshal(jsonData, &userMap)
	assert.NoError(t, err)
	
	assert.Equal(t, "user123", userMap["id"])
	assert.Equal(t, "Test User", userMap["name"])
	assert.Equal(t, "test@example.com", userMap["email"])
	assert.Equal(t, "admin", userMap["role"])
} 