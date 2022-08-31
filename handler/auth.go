package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yt-lite/auth/dto"
	"github.com/yt-lite/auth/service"
	"github.com/yt-lite/libs/errs"
	h "github.com/yt-lite/libs/handler"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return AuthHandler{authService}
}

func (ah *AuthHandler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.RegisterAuthRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			h.WriteError(c.Writer, errs.NewBadRequestError(err.Error()))
			return
		}
		if err := ah.authService.Register(req); err != nil {
			h.WriteError(c.Writer, err)
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
		})
	}
}

func (ah *AuthHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.LoginAuthRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			h.WriteError(c.Writer, errs.NewBadRequestError(err.Error()))
			return
		}
		auth, err := ah.authService.Login(req)
		if err != nil {
			h.WriteError(c.Writer, err)
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
			"data":    auth,
		})
	}
}
