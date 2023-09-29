package repository

import (
	"errors"
	"incentrick-restful-api/entity"
)

var ErrUserNotFound = errors.New("User not found")

type UserRepository interface {
	Get(id int) (*entity.User, error)
	// Create(in entity.User) (*entity.User, error)
	// Update(in entity.User) (*entity.User, error)
}
