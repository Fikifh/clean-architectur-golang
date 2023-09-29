package auth_handler

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	Login(c *gin.Context)
	OauthGoogleLogin(c *gin.Context)
	OauthGoogleCallback(c *gin.Context)
}
