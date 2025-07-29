package usecases

import (
	"task_manager/Domain/entities"
	"task_manager/Domain/interfaces"
)

type TaskUsecase interface {
	GetTasks() []entities.Task
	GetTaskByID(id string) (entities.Task, error)
	AddTask(task entities.Task) (entities.Task, error)
	UpdateTask(id string, updatedTask entities.Task) (entities.Task, error)
	DeleteTask(id string) error
}

type taskUsecase struct {
	taskRepo interfaces.TaskRepository
}

func NewTaskUsecase(taskRepo interfaces.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: taskRepo}
}

func (u *taskUsecase) GetTasks() []entities.Task {
	return u.taskRepo.GetTasks()
}

func (u *taskUsecase) GetTaskByID(id string) (entities.Task, error) {
	return u.taskRepo.GetTaskByID(id)
}

func (u *taskUsecase) AddTask(task entities.Task) (entities.Task, error) {
	return u.taskRepo.AddTask(task)
}

func (u *taskUsecase) UpdateTask(id string, updatedTask entities.Task) (entities.Task, error) {
	return u.taskRepo.UpdateTask(id, updatedTask)
}

func (u *taskUsecase) DeleteTask(id string) error {
	return u.taskRepo.DeleteTask(id)
} 