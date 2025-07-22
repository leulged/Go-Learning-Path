package routers

import (
	"task_manager/Delivery/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController *controllers.UserController, taskController *controllers.TaskController) {
	// === Public Routes ===
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	// === Admin-only User Management ===
	adminUserRoutes := r.Group("/users")
	adminUserRoutes.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminUserRoutes.POST("/promote", userController.Promote)
	}

	// === Authenticated User Routes (Tasks) ===
	taskRoutes := r.Group("/tasks")
	taskRoutes.Use(middleware.AuthMiddleware())
	{
		taskRoutes.GET("/", taskController.GetTasks)
		taskRoutes.GET("/:id", taskController.GetTaskByID)
	}

	// === Admin-only Task Management ===
	adminTaskRoutes := r.Group("/tasks")
	adminTaskRoutes.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminTaskRoutes.POST("/", taskController.AddTask)
		adminTaskRoutes.PUT("/:id", taskController.UpdateTask)
		adminTaskRoutes.DELETE("/:id", taskController.DeleteTask)
	}
}
