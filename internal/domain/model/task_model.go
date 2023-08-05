package model

import "time"

type Task struct {
	ID          int       `gorm:"primary_key" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskRepository interface {
	FindAll() ([]Task, error)
	FindByID(id int) (*Task, error)
	Create(task *Task) error
	Update(task *Task) error
	Delete(id int) error
}
