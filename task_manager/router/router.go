package router

import (
	"task_manager/controllers"
	"task_manager/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	taskService := services.NewTaskService()
	taskController := controllers.NewTaskController(taskService)

	router.GET("/task" , taskController.GetTasks)
	router.GET("/task/:id" , taskController.GetTaskByID)
	router.POST("/task" , taskController.AddTask)
	router.PUT("/task/:id" , taskController.UpdateTask)
	router.DELETE("/task/:id" , taskController.DeleteTask)
	return  router
}


