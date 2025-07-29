package routers

import (
	"task_manager/Delivery/http/controllers"
	"task_manager/Delivery/http/middleware"
	"task_manager/Domain/interfaces"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController *controllers.UserController, taskController *controllers.TaskController, tokenService interfaces.TokenService) {
	// === Public Routes ===
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	// === Admin-only User Management ===
	adminUserRoutes := r.Group("/users")
	adminUserRoutes.Use(middleware.AuthMiddleware(tokenService), middleware.AdminMiddleware())
	{
		adminUserRoutes.POST("/promote", userController.Promote)
	}

	// === Authenticated User Routes (Tasks) ===
	taskRoutes := r.Group("/tasks")
	taskRoutes.Use(middleware.AuthMiddleware(tokenService))
	{
		taskRoutes.GET("/", taskController.GetTasks)
		taskRoutes.GET("/:id", taskController.GetTaskByID)
	}

	// === Admin-only Task Management ===
	adminTaskRoutes := r.Group("/tasks")
	adminTaskRoutes.Use(middleware.AuthMiddleware(tokenService), middleware.AdminMiddleware())
	{
		adminTaskRoutes.POST("/", taskController.AddTask)
		adminTaskRoutes.PUT("/:id", taskController.UpdateTask)
		adminTaskRoutes.DELETE("/:id", taskController.DeleteTask)
	}
} 