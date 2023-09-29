package entity

import entity_authentication "incentrick-restful-api/entity/Authentication"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserWithAuth struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token *entity_authentication.Authentication
}
