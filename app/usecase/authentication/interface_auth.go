package authentication

import (
	"incentrick-restful-api/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UseCase interface {
	Login(idToken string, token string, platform string, w http.ResponseWriter, r *http.Request) (*entity.UserWithAuth, error)
	OauthGoogleLogin(c *gin.Context)
	OauthGoogleCallback(c *gin.Context)
	// Register(*entity.User) (*entity.User, error)
	// Logout(*entity.User) (*entity.User, error)
	// CreateToken(*entity.User, int) string //(*entity_authentication.Authentication, error)
}
