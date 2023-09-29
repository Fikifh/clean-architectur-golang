package user_repository

import (
	"time"
)

type UserModel struct {
	Id        int        `gorm:"primary_key;column:id;type:biginteger"`
	Name      string     `gorm:"column:name"`
	Email     string     `gorm:"column:email"`
	Password  string     `gorm:"column:password"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

type UserWithAuthModel struct {
	Id        int        `gorm:"primary_key;column:id;type:biginteger"`
	Name      string     `gorm:"column:name"`
	Email     string     `gorm:"column:email"`
	Password  string     `gorm:"column:password"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	Token     AuthModel  `gorm:"foreignKey:user_id;references:id"`
}

type AuthModel struct {
	Id           int        `gorm:"primary_key;column:id;biginteger"`
	UserId       int        `gorm:"primary_key;column:id;biginteger"`
	AccessToken  string     `gorm:"column:access_token"`
	RefreshToken string     `gorm:"column:refresh_token"`
	IsRevoked    string     `gorm:"column:is_revoked"`
	CreatedAt    *time.Time `gorm:"column:created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at"`
}

// func (UserModel) FromUserEntity(v entity.User) *UserModel {
// 	return &UserModel{
// 		Id:    v.Id,
// 		Name:  v.Name,
// 		Email: v.Email,
// 	}
// }

// func (m *UserModel) ToUserEntity() *entity.User {
// 	return &entity.User{
// 		Id:    m.Id,
// 		Name:  m.Name,
// 		Email: m.Email,
// 	}
// }
