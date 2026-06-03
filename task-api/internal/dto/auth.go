package dto

type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"admin"`
}

type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6Ik..."`
}

type LoginBadRequestResponse struct {
	Message string `json:"message" example:"invalid request"`
}

type LoginUnauthorizedResponse struct {
	Message string `json:"message" example:"invalid credentials"`
}

type LoginCouldNotGenerateTokenResponse struct {
	Message string `json:"message" example:"could not generate token"`
}
