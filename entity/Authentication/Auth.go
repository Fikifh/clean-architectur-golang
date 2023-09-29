package entity_authentication

import "github.com/golang-jwt/jwt"

type Authentication struct {
	Id           int    `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IsRevoked    bool   `json:"is_revoked"`
	Platform     string `json:"platform"`
}

type GoogleClaims struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Picture   string `json:"picture"`
	jwt.StandardClaims
}
