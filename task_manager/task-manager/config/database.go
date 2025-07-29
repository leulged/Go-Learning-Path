package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TaskCollection *mongo.Collection
var UserCollection *mongo.Collection

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	URI      string
	Database string
	Timeout  time.Duration
}

// NewDatabaseConfig creates a new database configuration
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		URI:      getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		Database: getEnv("DATABASE_NAME", "task_management_system"),
		Timeout:  10 * time.Second,
	}
}

// ConnectToMongo connects to MongoDB with configuration
func ConnectToMongo() *mongo.Client {
	config := NewDatabaseConfig()
	
	clientOptions := options.Client().ApplyURI(config.URI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("MongoDB client creation error:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping error:", err)
	}

	fmt.Println("✅ Connected to MongoDB!")

	db := client.Database(config.Database)
	TaskCollection = db.Collection("tasks")
	UserCollection = db.Collection("users")

	return client
} 