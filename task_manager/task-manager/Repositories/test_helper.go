// Repositories/test_helper.go
package Repositories

import (
    "context"
    "task_manager/Infrastructure"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

var (
    taskCollection *mongo.Collection
    userCollection *mongo.Collection
)

func setupTestDB() {
    godotenv.Load("../.env")
    Infrastructure.ConnectToMongo()

    // Clear both task and user collections
    Infrastructure.TaskCollection.DeleteMany(context.TODO(), bson.M{})
    Infrastructure.UserCollection.DeleteMany(context.TODO(), bson.M{})

	taskCollection = Infrastructure.TaskCollection
	userCollection = Infrastructure.UserCollection
}


