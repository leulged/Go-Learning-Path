package Usecases

import (
	"task_manager/domain"
)


type TaskUsecase interface {
	GetTasks() []domain.Task
	GetTaskByID(id string) (domain.Task, error)
	AddTask(task domain.Task) (domain.Task, error)
	UpdateTask(id string, updatedTask domain.Task) (domain.Task, error)
	DeleteTask(id string) error
}

type taskUsecase struct {
	taskRepo domain.TaskRepository
}

func NewTaskUsecase(taskRepo domain.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: taskRepo}
}

func (u *taskUsecase) GetTasks() []domain.Task {
	return u.taskRepo.GetTasks()
}

func (u *taskUsecase) GetTaskByID(id string) (domain.Task, error) {
	return u.taskRepo.GetTaskByID(id)
}

func (u *taskUsecase) AddTask(task domain.Task) (domain.Task, error) {
	return u.taskRepo.AddTask(task)
}

func (u *taskUsecase) UpdateTask(id string, updatedTask domain.Task) (domain.Task, error) {
	return u.taskRepo.UpdateTask(id, updatedTask)
}

func (u *taskUsecase) DeleteTask(id string) error {
	return u.taskRepo.DeleteTask(id)
}

