package Repositories

import (
	"context"
	"log"

	"task_manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) domain.TaskRepository {
	return &taskRepository{collection: collection}
}

func (r *taskRepository) GetTasks() []domain.Task {
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Error fetching tasks:", err)
		return nil
	}
	defer cursor.Close(context.TODO())

	var tasks []domain.Task
	for cursor.Next(context.TODO()) {
		var task domain.Task
		if err := cursor.Decode(&task); err != nil {
			log.Println("Error decoding task:", err)
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func (r *taskRepository) GetTaskByID(id string) (domain.Task, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Task{}, err
	}
	
	var task domain.Task
	err = r.collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&task)
	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

func (r *taskRepository) AddTask(task domain.Task) (domain.Task, error) {
	result, err := r.collection.InsertOne(context.TODO(), task)
	if err != nil {
		return domain.Task{}, err
	}
	task.ID = result.InsertedID.(primitive.ObjectID)
	return task, nil
}

func (r *taskRepository) UpdateTask(id string, updatedTask domain.Task) (domain.Task, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Task{}, err
	}
	
	_, err = r.collection.UpdateOne(context.TODO(),
		bson.M{"_id": objectID},
		bson.M{"$set": updatedTask})
	if err != nil {
		return domain.Task{}, err
	}
	return updatedTask, nil
}

func (r *taskRepository) DeleteTask(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	
	_, err = r.collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	return nil
}
