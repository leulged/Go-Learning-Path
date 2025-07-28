package usecases_test

import (
    "testing"

    "task_manager/domain"
    usecase "task_manager/Usecases"
    "task_manager/mocks"

    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/assert"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetTasksByID(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockTaskRepo := mocks.NewMockTaskRepository(ctrl)

    // sample data
    taskID := "1"
    objectID, _ := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")
    expected := domain.Task{
        ID:          objectID,
        Title:       "Test Task",
        Description: "Test Description",
        Status:      "pending",
    }

    mockTaskRepo.EXPECT().GetTaskByID(taskID).Return(expected, nil)

    taskUsecase := usecase.NewTaskUsecase(mockTaskRepo)
    task, err := taskUsecase.GetTaskByID(taskID)

    assert.NoError(t, err)
    assert.Equal(t, expected, task)
}

func TestAddTask(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockTaskRepo := mocks.NewMockTaskRepository(ctrl)

	task := domain.Task{
		Title: "Test Task",
		Description: "Test Description",
		Status: "pending",
	}
	mockTaskRepo.EXPECT().AddTask(task).Return(task, nil)
	
	taskUsecase := usecase.NewTaskUsecase(mockTaskRepo)
	result, err := taskUsecase.AddTask(task)

	assert.NoError(t, err)
	assert.Equal(t, task, result)
}

func TestGetTasks(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockTaskRepo := mocks.NewMockTaskRepository(ctrl)

    // sample data
    expected := []domain.Task{
        {
            // ID:          objectID,
            Title:       "Test Task",
            Description: "Test Description",
            Status:      "pending",
        },
        {
            // ID:          objectID,
            Title:       "Test Task 2",
            Description: "Test Description 2",
            Status:      "completed",
        },
    }

    mockTaskRepo.EXPECT().GetTasks().Return(expected)

    taskUsecase := usecase.NewTaskUsecase(mockTaskRepo)
    tasks := taskUsecase.GetTasks()

    assert.Equal(t, expected, tasks)
}

func TestUpdateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskRepo := mocks.NewMockTaskRepository(ctrl)

	task := domain.Task{
		Title: "Test Task",
		Description: "Test Description",
		Status: "pending",
	}
	mockTaskRepo.EXPECT().UpdateTask(task.ID.Hex(), task).Return(task, nil)

	taskUsecase := usecase.NewTaskUsecase(mockTaskRepo)
	result, err := taskUsecase.UpdateTask(task.ID.Hex(), task)

	assert.NoError(t, err)
	assert.Equal(t, task, result)
}

func TestDeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskRepo := mocks.NewMockTaskRepository(ctrl)				

	taskID := "1"
	mockTaskRepo.EXPECT().DeleteTask(taskID).Return(nil)

	taskUsecase := usecase.NewTaskUsecase(mockTaskRepo)
	err := taskUsecase.DeleteTask(taskID)

	assert.NoError(t, err)
} 