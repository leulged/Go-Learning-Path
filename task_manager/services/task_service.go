package services

import (
	"context"
	"fmt"
	"log"

	"task_manager/data"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TaskService defines the interface for task operations.
type TaskService interface {
	GetTasks() []models.Task
	GetTaskByID(id int) (models.Task, error)
	AddTask(task models.Task) (models.Task, error)
	UpdateTask(id int, updatedTask models.Task) (models.Task, error)
	DeleteTask(id int) error
}

// taskServiceImpl implements TaskService using an in-memory map.
type taskServiceImpl struct {
	collection *mongo.Collection
}

// NewTaskService initializes the service with an empty map.
func NewTaskService() TaskService {
	return &taskServiceImpl{
		collection : data.TaskCollection,
	}
}

// ✅ GetTasks returns all tasks as a slice
func (ts *taskServiceImpl) GetTasks() []models.Task {
	
	cursor , err := ts.collection.Find(context.TODO() , bson.M{})
	if err != nil {
		log.Println("Error Fetching Tasks")
		return nil
	}
	defer cursor.Close(context.TODO())
	var tasks []models.Task
	for cursor.Next(context.TODO()){
		var task models.Task

		if err := cursor.Decode(&task) ; err != nil {
			log.Println("Error Fetching Task :" , err)
			continue
		}
		tasks = append(tasks, task)


	}
	return tasks
}

// ✅ GetTaskByID finds a task by its ID or returns an error
func (ts *taskServiceImpl) GetTaskByID(id int) (models.Task, error) {
	var task models.Task
	err := ts.collection.FindOne(context.TODO() , bson.M{"id" : id}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{} , fmt.Errorf("task not found")
		}
		return models.Task{} , err
	}
	return task , nil
}

// ✅ AddTask adds a task to the map
func (ts *taskServiceImpl) AddTask(task models.Task) (models.Task, error) {
	_ , err := ts.collection.InsertOne(context.TODO() , task)
	if err != nil {
		return models.Task{} , err
	}
	return task , nil
}

// ✅ UpdateTask updates fields if the task exists
func (ts *taskServiceImpl) UpdateTask(id int, updatedTask models.Task) (models.Task, error) {
	filter := bson.M{"id" : id}
	update := bson.M{
		"$set" : bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"status":      updatedTask.Status,
			"due_date":    updatedTask.DueDate,
		},

	}
	result , err := ts.collection.UpdateOne(context.TODO() , filter , update)
	if err != nil {
		return models.Task{} , err
	}
	if result.MatchedCount == 0 {
		return models.Task{} , fmt.Errorf("task not found")
	}
	updatedTask.ID = id
	return updatedTask , nil
}

// ✅ DeleteTask removes task if it exists
func (ts *taskServiceImpl) DeleteTask(id int) error {
	result , err := ts.collection.DeleteOne(context.TODO() , bson.M{"id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return  fmt.Errorf("task not found")
	}
	
	return nil
}
