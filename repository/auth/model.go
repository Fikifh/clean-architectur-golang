package auth

import (
	"time"

	"incentrick-restful-api/entity"
	entity_authentication "incentrick-restful-api/entity/Authentication"
)

type UserWithAuthModel struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Token     *entity_authentication.Authentication
}

type AuthModel struct {
	Id           int
	AccessToken  string
	RefreshToken string
	IsRevoked    bool
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

func (user *UserWithAuthModel) ToAuthUser() *entity.UserWithAuth {
	return &entity.UserWithAuth{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Token: user.Token,
	}
}

func (m *AuthModel) ToAuthEntity() *entity_authentication.Authentication {
	return &entity_authentication.Authentication{
		Id:           m.Id,
		AccessToken:  m.AccessToken,
		RefreshToken: m.RefreshToken,
		IsRevoked:    m.IsRevoked,
	}
}
