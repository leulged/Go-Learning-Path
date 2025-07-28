package main

import (
	"context"
	"log"
	"os"
	"task_manager/Infrastructure"
	"task_manager/Repositories"
	"task_manager/Usecases"
	"task_manager/Delivery/controllers"
	"task_manager/Delivery/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Connect to MongoDB
	client := Infrastructure.ConnectToMongo()
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	// Initialize repositories
	userRepo := Repositories.NewUserRepository(Infrastructure.UserCollection)
	taskRepo := Repositories.NewTaskRepository(Infrastructure.TaskCollection)

	// Initialize usecases
	userUsecase := Usecases.NewUserUsecase(userRepo)
	taskUsecase := Usecases.NewTaskUsecase(taskRepo)

	// Initialize controllers
	userController := controllers.NewUserController(userUsecase)
	taskController := controllers.NewTaskController(taskUsecase)

	// Setup Gin router
	r := gin.Default()

	// Register middleware globally if needed
	// r.Use(middleware.AuthMiddleware()) // Or apply selectively in router groups

	// Setup routes
	routers.SetupRoutes(r, userController, taskController)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
