package models

import (
	"task_manager/Domain/entities"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskDocument represents the MongoDB document structure
type TaskDocument struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	DueDate     time.Time          `bson:"due_date"`
	Status      string             `bson:"status"`
}

// TaskFromDomain converts domain Task to MongoDB TaskDocument
func TaskFromDomain(task entities.Task) (TaskDocument, error) {
	var objectID primitive.ObjectID
	var err error

	if task.ID != "" {
		objectID, err = primitive.ObjectIDFromHex(task.ID)
		if err != nil {
			return TaskDocument{}, err
		}
	}

	return TaskDocument{
		ID:          objectID,
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
		Status:      task.Status,
	}, nil
}

// TaskToDomain converts MongoDB TaskDocument to domain Task
func TaskToDomain(doc TaskDocument) entities.Task {
	return entities.Task{
		ID:          doc.ID.Hex(),
		Title:       doc.Title,
		Description: doc.Description,
		DueDate:     doc.DueDate,
		Status:      doc.Status,
	}
} 