package usecases_test

import (
	"testing"
	"time"
	"task_manager/Domain/entities"
	"task_manager/Domain/errors"
	"task_manager/mocks"
	usecase "task_manager/Usecases"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetTaskByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskRepo := mocks.NewMockTaskRepository(ctrl)

	// sample data
	taskID := "507f1f77bcf86cd799439011"
	expected := entities.Task{
		ID:          taskID,
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "Pending",
		DueDate:     time.Now(),
	}

	mockTaskRepo.EXPECT().GetTaskByID(taskID).Return(expected, nil)

	taskUsecase := usecase.NewTaskUsecase(mockTaskRepo)
	task, err := taskUsecase.GetTaskByID(taskID)

	assert.NoError(t, err)
	assert.Equal(t, expected, task)
}

func TestGetTaskByIDNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskRepo := mocks.NewMockTaskRepository(ctrl)

	taskID := "507f1f77bcf86cd799439011"
	mockTaskRepo.EXPECT().GetTaskByID(taskID).Return(entities.Task{}, errors.TaskNotFoundError{})

	taskUsecase := usecase.NewTaskUsecase(mockTaskRepo)
	task, err := taskUsecase.GetTaskByID(taskID)

	assert.Error(t, err)
	assert.IsType(t, errors.TaskNotFoundError{}, err)
	assert.Equal(t, entities.Task{}, task)
}

func TestAddTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskRepo := mocks.NewMockTaskRepository(ctrl)

	task := entities.NewTask("Test Task", "Test Description", time.Now())
	mockTaskRepo.EXPECT().AddTask(gomock.Any()).Return(task, nil)
	
	taskUsecase := usecase.NewTaskUsecase(mockTaskRepo)
	result, err := taskUsecase.AddTask(task)

	assert.NoError(t, err)
	assert.Equal(t, task.Title, result.Title)
	assert.Equal(t, task.Description, result.Description)
	assert.Equal(t, "Pending", result.Status)
}

func TestGetTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskRepo := mocks.NewMockTaskRepository(ctrl)

	// sample data
	expected := []entities.Task{
		{
			ID:          "1",
			Title:       "Test Task",
			Description: "Test Description",
			Status:      "Pending",
			DueDate:     time.Now(),
		},
		{
			ID:          "2",
			Title:       "Test Task 2",
			Description: "Test Description 2",
			Status:      "Completed",
			DueDate:     time.Now(),
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

	taskID := "507f1f77bcf86cd799439011"
	task := entities.Task{
		ID:          taskID,
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      "In Progress",
		DueDate:     time.Now(),
	}
	mockTaskRepo.EXPECT().UpdateTask(taskID, task).Return(task, nil)

	taskUsecase := usecase.NewTaskUsecase(mockTaskRepo)
	result, err := taskUsecase.UpdateTask(taskID, task)

	assert.NoError(t, err)
	assert.Equal(t, task, result)
}

func TestUpdateTaskNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskRepo := mocks.NewMockTaskRepository(ctrl)

	taskID := "507f1f77bcf86cd799439011"
	task := entities.Task{
		ID:          taskID,
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      "In Progress",
		DueDate:     time.Now(),
	}
	mockTaskRepo.EXPECT().UpdateTask(taskID, task).Return(entities.Task{}, errors.TaskNotFoundError{})

	taskUsecase := usecase.NewTaskUsecase(mockTaskRepo)
	result, err := taskUsecase.UpdateTask(taskID, task)

	assert.Error(t, err)
	assert.IsType(t, errors.TaskNotFoundError{}, err)
	assert.Equal(t, entities.Task{}, result)
}

func TestDeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskRepo := mocks.NewMockTaskRepository(ctrl)				

	taskID := "507f1f77bcf86cd799439011"
	mockTaskRepo.EXPECT().DeleteTask(taskID).Return(nil)

	taskUsecase := usecase.NewTaskUsecase(mockTaskRepo)
	err := taskUsecase.DeleteTask(taskID)

	assert.NoError(t, err)
}

func TestDeleteTaskNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskRepo := mocks.NewMockTaskRepository(ctrl)				

	taskID := "507f1f77bcf86cd799439011"
	mockTaskRepo.EXPECT().DeleteTask(taskID).Return(errors.TaskNotFoundError{})

	taskUsecase := usecase.NewTaskUsecase(mockTaskRepo)
	err := taskUsecase.DeleteTask(taskID)

	assert.Error(t, err)
	assert.IsType(t, errors.TaskNotFoundError{}, err)
} 