package main

import (
	"context"
	"fmt"
	"log"
	"task_manager/config"
	"task_manager/Delivery/http/controllers"
	"task_manager/Delivery/http/routers"
	"task_manager/Infrastructure/database/repositories"
	"task_manager/Infrastructure/services"
	usecases "task_manager/Usecases"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Load application configuration
	appConfig := config.NewAppConfig()

	// Connect to MongoDB using config
	client := config.ConnectToMongo()
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	// Initialize repositories with clean architecture
	userRepo := repositories.NewUserRepository(config.UserCollection)
	taskRepo := repositories.NewTaskRepository(config.TaskCollection)

	// Initialize services
	tokenService := services.NewJWTService()

	// Initialize use cases with clean dependencies
	userUsecase := usecases.NewUserUsecase(userRepo, tokenService)
	taskUsecase := usecases.NewTaskUsecase(taskRepo)

	// Initialize controllers
	userController := controllers.NewUserController(userUsecase)
	taskController := controllers.NewTaskController(taskUsecase)

	// Setup Gin router
	r := gin.Default()

	// Setup routes with clean middleware
	routers.SetupRoutes(r, userController, taskController, tokenService)

	// Start server with config
	port := fmt.Sprintf(":%s", appConfig.Port)
	log.Printf("Starting server on port %s\n", appConfig.Port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
} 