package usecase

import "go-restapi/internal/domain/model"

type TaskUsecase interface {
	GetAllTasks() ([]model.Task, error)
	GetTaskByID(id int) (*model.Task, error)
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task) error
	DeleteTask(id int) error
}

type taskUsecase struct {
	taskRepo model.TaskRepository
}

func NewTaskUsecase(taskRepo model.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo}
}

func (u *taskUsecase) GetAllTasks() ([]model.Task, error) {
	return u.taskRepo.FindAll()
}

func (u *taskUsecase) GetTaskByID(id int) (*model.Task, error) {
	return u.taskRepo.FindByID(id)
}

func (u *taskUsecase) CreateTask(task *model.Task) error {
	return u.taskRepo.Create(task)
}

func (u *taskUsecase) UpdateTask(task *model.Task) error {
	return u.taskRepo.Update(task)
}

func (u *taskUsecase) DeleteTask(id int) error {
	return u.taskRepo.Delete(id)
}
