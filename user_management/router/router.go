package router

import (
	"user_management/controllers"
	"user_management/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	userService := service.NewUserService()
	userController := controllers.NewUserController(userService)

	router.GET("/users" , userController.GetUsers)
	router.POST("/users" , userController.CreateUser)
	router.GET("/users/:id" , userController.GetUserById)
	router.PUT("/users/:id" , userController.UpdateUser)
	router.DELETE("/users/:id" , userController.DeleteUser)
	return router
}

