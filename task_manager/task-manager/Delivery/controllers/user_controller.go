package controllers

import (
	"net/http"
	"task_manager/domain"
	usecases "task_manager/Usecases"
	"task_manager/Delivery/http/request"

	"github.com/gin-gonic/gin"
)

// UserController handles user-related HTTP requests
type UserController struct {
	Service usecases.UserUsecase
}

// NewUserController creates and returns a new UserController instance
func NewUserController(service usecases.UserUsecase) *UserController {
	return &UserController{
		Service: service,
	}
}

// Register new user
func (uc *UserController) Register(c *gin.Context) {
	var input request.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert RegisterInput to User
	user := domain.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	createdUser, err := uc.Service.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

// Login existing user
func (uc *UserController) Login(c *gin.Context) {
	var input request.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := uc.Service.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token": token,
	})
}

// Promote user to admin
func (uc *UserController) Promote(c *gin.Context) {
	var input request.PromoteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uc.Service.PromoteToAdmin(input.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin"})
}
