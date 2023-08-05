package model

import "time"

type User struct {
	ID        int       `gorm:"primary_key" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(user *User) error
	FindByID(id int) (*User, error)
	// Define user related methods here
}
