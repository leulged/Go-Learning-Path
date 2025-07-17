package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global collection reference for the "user" collection
var UserCollection *mongo.Collection

func ConnectToMongo() {
	uri := "mongodb+srv://leulgedion224:YtxgbwYFwW9snTti@cluster0.gdmxw28.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	// Correct capitalization: `options.Client()` not `options.client()`
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	// Check the connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("MongoDB ping error:", err)
	}

	fmt.Println("MongoDB connected successfully")

	// Use correct collection initialization
	db := client.Database("user_management")
	UserCollection = db.Collection("user")
	
}
