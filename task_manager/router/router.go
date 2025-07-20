
package router

import (
	"task_manager/controllers"
	"task_manager/middleware"
	"task_manager/services"

	"github.com/gin-gonic/gin"
)
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Initialize task service/controller
	taskService := services.NewTaskService()
	taskController := controllers.NewTaskController(taskService)

	// Initialize user service/controller
	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	// Public routes for registration and login (no auth middleware)
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	// Admin-only user management routes (require auth + admin)
	adminUserRoutes := router.Group("/user")
	adminUserRoutes.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminUserRoutes.POST("/promote", userController.Promote)
	}

	// User-level task routes (require auth)
	taskRoutes := router.Group("/task")
	taskRoutes.Use(middleware.AuthMiddleware())
	{
		taskRoutes.GET("/", taskController.GetTasks)
		taskRoutes.GET("/:id", taskController.GetTaskByID)
	}

	// Admin-only task routes (require auth + admin)
	adminTaskRoutes := router.Group("/task")
	adminTaskRoutes.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminTaskRoutes.POST("/", taskController.AddTask)
		adminTaskRoutes.PUT("/:id", taskController.UpdateTask)
		adminTaskRoutes.DELETE("/:id", taskController.DeleteTask)
	}

	return router
}
