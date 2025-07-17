package controllers

import (
	// "context"
	"net/http"
	"strconv"

	"user_management/models"
	"user_management/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		Service: service,
	}
}

func (uci *UserController) GetUsers(c *gin.Context) {
	users := uci.Service.GetUsers()
	c.JSON(http.StatusOK, users)
}
func (uci *UserController) CreateUser (c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest , err)
		return
	}

	newUser , err := uci.Service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest , err)
		return
	}
	c.JSON(http.StatusOK , newUser)
	
}

func (uci *UserController) GetUserById(c *gin.Context) {
	idParam := c.Param("id")
	id ,err := strconv.Atoi(idParam) 
	if err != nil {
		c.JSON(http.StatusBadRequest , err)
		return
	}
	task , err := uci.Service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest ,err)
		return
	}
	c.JSON(http.StatusOK , task)
}

func (uci *UserController) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var updateUser models.User
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	updated, err := uci.Service.UpdateUser(id, updateUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (uci *UserController) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest , err)
	}

	err = uci.Service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest , err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}