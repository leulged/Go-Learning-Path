package repositories

import (
	"context"
	"path/filepath"
	"testing"
	"time"
	"task_manager/Domain/entities"
	"task_manager/Infrastructure/database/models"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupTaskTestDB(t *testing.T) (*mongo.Collection, func()) {
	// Load environment variables
	err := godotenv.Load(filepath.Join("..", "..", "..", ".env"))
	if err != nil {
		// Try current directory as fallback
		err = godotenv.Load()
		if err != nil {
			panic("Error loading .env file: " + err.Error())
		}
	}

	// Connect to test database using MongoDB Atlas
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://leulgedion224:YtxgbwYFwW9snTti@cluster0.gdmxw28.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	require.NoError(t, err)

	// Use a test database
	db := client.Database("test_task_manager")
	collection := db.Collection("tasks")

	// Clean up before test
	_, err = collection.DeleteMany(context.Background(), bson.M{})
	require.NoError(t, err)

	// Return cleanup function
	cleanup := func() {
		collection.DeleteMany(context.Background(), bson.M{})
		client.Disconnect(context.Background())
	}

	return collection, cleanup
}

func TestTaskRepository_AddTask(t *testing.T) {
	collection, cleanup := setupTaskTestDB(t)
	defer cleanup()

	repo := NewTaskRepository(collection)

	// Test task
	task := entities.NewTask("Test Task", "Test Description", time.Now())

	// Test AddTask
	addedTask, err := repo.AddTask(task)
	assert.NoError(t, err)
	assert.NotEmpty(t, addedTask.ID)

	// Verify task was saved
	var doc models.TaskDocument
	err = collection.FindOne(context.Background(), bson.M{"title": "Test Task"}).Decode(&doc)
	assert.NoError(t, err)
	assert.Equal(t, "Test Task", doc.Title)
	assert.Equal(t, "Test Description", doc.Description)
	assert.Equal(t, "Pending", doc.Status)
}

func TestTaskRepository_GetTaskByID(t *testing.T) {
	collection, cleanup := setupTaskTestDB(t)
	defer cleanup()

	repo := NewTaskRepository(collection)

	// Create test task
	task := entities.NewTask("Test Task", "Test Description", time.Now())
	addedTask, err := repo.AddTask(task)
	require.NoError(t, err)

	// Test GetTaskByID
	foundTask, err := repo.GetTaskByID(addedTask.ID)
	assert.NoError(t, err)
	assert.NotNil(t, foundTask)
	assert.Equal(t, "Test Task", foundTask.Title)
	assert.Equal(t, "Test Description", foundTask.Description)
	assert.Equal(t, "Pending", foundTask.Status)

	// Test GetTaskByID with non-existent ID
	notFoundTask, err := repo.GetTaskByID("nonexistentid")
	assert.Error(t, err)
	assert.Equal(t, entities.Task{}, notFoundTask)
}

func TestTaskRepository_GetTasks(t *testing.T) {
	collection, cleanup := setupTaskTestDB(t)
	defer cleanup()

	repo := NewTaskRepository(collection)

	// Create multiple tasks
	task1 := entities.NewTask("Task 1", "Description 1", time.Now())
	task2 := entities.NewTask("Task 2", "Description 2", time.Now())

	_, err := repo.AddTask(task1)
	require.NoError(t, err)
	_, err = repo.AddTask(task2)
	require.NoError(t, err)

	// Test GetTasks
	tasks := repo.GetTasks()
	assert.Len(t, tasks, 2)

	// Verify tasks
	titles := []string{tasks[0].Title, tasks[1].Title}
	assert.Contains(t, titles, "Task 1")
	assert.Contains(t, titles, "Task 2")
}

func TestTaskRepository_UpdateTask(t *testing.T) {
	collection, cleanup := setupTaskTestDB(t)
	defer cleanup()

	repo := NewTaskRepository(collection)

	// Create test task
	task := entities.NewTask("Test Task", "Test Description", time.Now())
	addedTask, err := repo.AddTask(task)
	require.NoError(t, err)

	// Update task
	updatedTask := entities.NewTask("Updated Task", "Updated Description", time.Now())
	updatedTask.SetStatus("Completed")
	
	resultTask, err := repo.UpdateTask(addedTask.ID, updatedTask)
	assert.NoError(t, err)

	// Verify update
	foundTask, err := repo.GetTaskByID(addedTask.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Task", foundTask.Title)
	assert.Equal(t, "Updated Description", foundTask.Description)
	assert.Equal(t, "Completed", foundTask.Status)
	assert.Equal(t, resultTask.Title, foundTask.Title)
}

func TestTaskRepository_DeleteTask(t *testing.T) {
	collection, cleanup := setupTaskTestDB(t)
	defer cleanup()

	repo := NewTaskRepository(collection)

	// Create test task
	task := entities.NewTask("Test Task", "Test Description", time.Now())
	addedTask, err := repo.AddTask(task)
	require.NoError(t, err)

	// Test DeleteTask
	err = repo.DeleteTask(addedTask.ID)
	assert.NoError(t, err)

	// Verify task was deleted
	foundTask, err := repo.GetTaskByID(addedTask.ID)
	assert.Error(t, err)
	assert.Equal(t, entities.Task{}, foundTask)
}

func TestTaskDocument_Mapping(t *testing.T) {
	// Test TaskToDomain
	doc := models.TaskDocument{
		ID:          primitive.NewObjectID(),
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "Pending",
		DueDate:     time.Now(),
	}

	task := models.TaskToDomain(doc)
	assert.Equal(t, doc.ID.Hex(), task.ID)
	assert.Equal(t, doc.Title, task.Title)
	assert.Equal(t, doc.Description, task.Description)
	assert.Equal(t, doc.Status, task.Status)

	// Test TaskFromDomain
	originalTask := entities.NewTask("Original Task", "Original Description", time.Now())
	originalTask.ID = "507f1f77bcf86cd799439011"

	convertedDoc, err := models.TaskFromDomain(originalTask)
	assert.NoError(t, err)
	assert.Equal(t, originalTask.ID, convertedDoc.ID.Hex())
	assert.Equal(t, originalTask.Title, convertedDoc.Title)
	assert.Equal(t, originalTask.Description, convertedDoc.Description)
	assert.Equal(t, originalTask.Status, convertedDoc.Status)
} 