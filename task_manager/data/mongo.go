package data

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TaskCollection *mongo.Collection

func ConnectToMongo() {
	uri := "mongodb+srv://leulgedion224:YtxgbwYFwW9snTti@cluster0.gdmxw28.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	clientOptions := options.Client().ApplyURI(uri)

	// Fix typo: use "context" not "connect"
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Use consistent and lowercase naming for collection
	db := client.Database("task_management_system")
	TaskCollection = db.Collection("tasks")
}
