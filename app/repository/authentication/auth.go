package repository_auth

import (
	"errors"
	"incentrick-restful-api/entity"
	entity_authentication "incentrick-restful-api/entity/Authentication"
)

var ErrUserNotFound = errors.New("User not found")

type AuthRepository interface {
	Login(*entity_authentication.GoogleClaims) (*entity.UserWithAuth, error)
	Register(*entity_authentication.GoogleClaims) (*entity.UserWithAuth, error)
	GetGooglePublicKey(string) (string, error)
	CreateOrUpdateAuth(string, string, string, int) (*entity.UserWithAuth, error)
	// Logout(*entity.User) (*entity.User, error)
	// CreateToken(*entity.User, int) string //(*entity_authentication.Authentication, error)
}
