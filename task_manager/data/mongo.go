package data

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Exported MongoDB collections
var TaskCollection *mongo.Collection
var UserCollection *mongo.Collection

// ConnectToMongo initializes the MongoDB client and collections
func ConnectToMongo() {
	uri := "mongodb+srv://leulgedion224:YtxgbwYFwW9snTti@cluster0.gdmxw28.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("MongoDB ping error:", err)
	}

	fmt.Println("Connected to MongoDB!")

	db := client.Database("task_management_system")
	TaskCollection = db.Collection("tasks")
	UserCollection = db.Collection("users")  // users collection for authentication
}
