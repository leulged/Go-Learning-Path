package Repositories

import (
	// "context"
	"testing"

	"task_manager/domain"
	// infrastructure "task_manager/Infrastructure"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAddTask(t *testing.T) {
	setupTestDB()
	repo := NewTaskRepository(taskCollection)

	dummyTask := domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
	}

	task, err := repo.AddTask(dummyTask)
	assert.NoError(t, err)
	assert.NotEqual(t, primitive.NilObjectID, task.ID)
}

func TestGetAllTasks(t *testing.T) {
	setupTestDB()
	repo := NewTaskRepository(taskCollection)

	task1 := domain.Task{
		Title:       "Test Task 1",
		Description: "Test Description 1",
		Status:      "pending",
	}
	_, err := repo.AddTask(task1)
	assert.NoError(t, err)

	task2 := domain.Task{
		Title:       "Test Task 2",
		Description: "Test Description 2",
		Status:      "pending",
	}
	_, err = repo.AddTask(task2)
	assert.NoError(t, err)

	tasks := repo.GetTasks()
	assert.GreaterOrEqual(t, len(tasks), 2)
}

func TestGetTaskByID(t *testing.T) {
	setupTestDB()
	repo := NewTaskRepository(taskCollection)

	dummyTask := domain.Task{
		Title:       "Test Task 1",
		Description: "Test Description 1",
		Status:      "pending",
	}
	inserted, err := repo.AddTask(dummyTask)
	assert.NoError(t, err)

	task, err := repo.GetTaskByID(inserted.ID.Hex())
	assert.NoError(t, err)
	assert.Equal(t, dummyTask.Title, task.Title)
}

func TestUpdateTask(t *testing.T) {
	setupTestDB()
	repo := NewTaskRepository(taskCollection)

	dummyTask := domain.Task{
		Title:       "Test Task 1",
		Description: "Test Description 1",
		Status:      "pending",
	}
	inserted, err := repo.AddTask(dummyTask)
	assert.NoError(t, err)

	updateTask := domain.Task{
		Title:       "Test Task 1 Updated",
		Description: "Test Description 1 Updated",
		Status:      "completed",
	}

	updated, err := repo.UpdateTask(inserted.ID.Hex(), updateTask)
	assert.NoError(t, err)
	assert.Equal(t, updateTask.Title, updated.Title)
	assert.Equal(t, updateTask.Description, updated.Description)
	assert.Equal(t, updateTask.Status, updated.Status)
}

func TestDeleteTask(t *testing.T) {
	setupTestDB()
	repo := NewTaskRepository(taskCollection)

	dummyTask := domain.Task{
		Title:       "Test Task 1",
		Description: "Test Description 1",
		Status:      "pending",
	}
	inserted, err := repo.AddTask(dummyTask)
	assert.NoError(t, err)

	err = repo.DeleteTask(inserted.ID.Hex())
	assert.NoError(t, err)

	deletedTask, err := repo.GetTaskByID(inserted.ID.Hex())
	assert.Error(t, err)
	assert.Equal(t, domain.Task{}, deletedTask)
}
