package domain

import (
	"go-restapi/internal/domain/model"

	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) model.TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) FindAll() ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) FindByID(id int) (*model.Task, error) {
	var task model.Task
	err := r.db.First(&task, id).Error
	return &task, err
}

func (r *taskRepository) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) Update(task *model.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) Delete(id int) error {
	return r.db.Delete(&model.Task{}, id).Error
}
