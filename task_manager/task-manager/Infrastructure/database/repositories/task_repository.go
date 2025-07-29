package repositories

import (
	"context"
	"log"
	"task_manager/Domain/entities"
	"task_manager/Domain/errors"
	"task_manager/Domain/interfaces"
	"task_manager/Infrastructure/database/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) interfaces.TaskRepository {
	return &taskRepository{collection: collection}
}

func (r *taskRepository) GetTasks() []entities.Task {
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Error fetching tasks:", err)
		return nil
	}
	defer cursor.Close(context.TODO())

	var tasks []entities.Task
	for cursor.Next(context.TODO()) {
		var doc models.TaskDocument
		if err := cursor.Decode(&doc); err != nil {
			log.Println("Error decoding task:", err)
			continue
		}
		tasks = append(tasks, models.TaskToDomain(doc))
	}
	return tasks
}

func (r *taskRepository) GetTaskByID(id string) (entities.Task, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entities.Task{}, errors.InvalidTaskIDError{}
	}
	
	var doc models.TaskDocument
	err = r.collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entities.Task{}, errors.TaskNotFoundError{}
		}
		return entities.Task{}, err
	}
	
	return models.TaskToDomain(doc), nil
}

func (r *taskRepository) AddTask(task entities.Task) (entities.Task, error) {
	doc, err := models.TaskFromDomain(task)
	if err != nil {
		return entities.Task{}, err
	}

	result, err := r.collection.InsertOne(context.TODO(), doc)
	if err != nil {
		return entities.Task{}, errors.TaskCreationError{Message: "failed to create task"}
	}

	doc.ID = result.InsertedID.(primitive.ObjectID)
	return models.TaskToDomain(doc), nil
}

func (r *taskRepository) UpdateTask(id string, updatedTask entities.Task) (entities.Task, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entities.Task{}, errors.InvalidTaskIDError{}
	}
	
	doc, err := models.TaskFromDomain(updatedTask)
	if err != nil {
		return entities.Task{}, err
	}
	
	_, err = r.collection.UpdateOne(context.TODO(),
		bson.M{"_id": objectID},
		bson.M{"$set": doc})
	if err != nil {
		return entities.Task{}, errors.TaskUpdateError{Message: "failed to update task"}
	}
	
	return models.TaskToDomain(doc), nil
}

func (r *taskRepository) DeleteTask(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.InvalidTaskIDError{}
	}
	
	_, err = r.collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return errors.TaskUpdateError{Message: "failed to delete task"}
	}
	
	return nil
} 