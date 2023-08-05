package domain

import (
	"go-restapi/internal/domain/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id int) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return &user, err
}
