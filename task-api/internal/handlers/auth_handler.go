package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"task-api/internal/auth"
	"task-api/internal/config"
	"task-api/internal/dto"
)

// Login godoc
// @Summary Login user
// @Description Authenticates user with username and password. Returns JWT token if credentials are valid.
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body dto.LoginRequest true "Login credentials"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.LoginBadRequestResponse
// @Failure 401 {object} dto.LoginUnauthorizedResponse
// @Failure 500 {object} dto.LoginCouldNotGenerateTokenResponse
// @Router /login [post]
func Login(c *gin.Context, cfg *config.Config) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.LoginBadRequestResponse{
			Message: "invalid request",
		})
		return
	}

	if req.Username == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, dto.LoginBadRequestResponse{
			Message: "invalid request",
		})
		return
	}

	if req.Username != cfg.Auth.AdminUsername || req.Password != cfg.Auth.AdminPassword {
		c.JSON(http.StatusUnauthorized, dto.LoginUnauthorizedResponse{
			Message: "invalid credentials",
		})
		return
	}

	token, err := auth.GenerateToken(cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.LoginCouldNotGenerateTokenResponse{
			Message: "could not generate token",
		})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		Token: token,
	})
}
