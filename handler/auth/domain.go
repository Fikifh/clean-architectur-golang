package auth_handler

type LoginRequest struct {
	AccessToken string `form:"access_token" binding:"required"`
	TokenId     string `form:"token_id" binding:"required"`
}
