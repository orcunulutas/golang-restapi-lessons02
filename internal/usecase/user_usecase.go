package usecase

import "go-restapi/internal/domain/model"

type UserUsecase interface {
	CreateUser(user *model.User) error
	GetUserByID(id int) (*model.User, error)
	// Define user related use cases here
}

type userUsecase struct {
	userRepo model.UserRepository
}

func NewUserUsecase(userRepo model.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) CreateUser(user *model.User) error {
	return u.userRepo.Create(user)
}

func (u *userUsecase) GetUserByID(id int) (*model.User, error) {
	return u.userRepo.FindByID(id)
}

// Define user related use cases here
