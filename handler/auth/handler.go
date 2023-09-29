package auth_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"incentrick-restful-api/app/usecase/authentication"
	utils "incentrick-restful-api/handler/Utils"
)

type authHandler struct {
	authUseCase authentication.UseCase
}

func NewHandler(authUseCase authentication.UseCase) AuthHandler {
	return &authHandler{authUseCase: authUseCase}
}

func (h *authHandler) Login(c *gin.Context) {
	var request LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{Message: err.Error()})
		return
	}

	// user_agent.New(c)
	data, err := h.authUseCase.Login(request.TokenId, request.AccessToken, "android", c.Writer, c.Request)
	if err == nil {
		c.JSON(http.StatusOK, utils.SuccessResponse{Data: data})
	} else {
		c.JSON(http.StatusNotFound, utils.ErrorResponse{Message: err.Error()})
	}
}

func (h *authHandler) OauthGoogleLogin(c *gin.Context) {
	h.authUseCase.OauthGoogleLogin(c)
	return
}

func (h *authHandler) OauthGoogleCallback(c *gin.Context) {
	h.authUseCase.OauthGoogleCallback(c)
	return
}
