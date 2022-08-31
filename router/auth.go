package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yt-lite/auth/db"
	"github.com/yt-lite/auth/handler"
	"github.com/yt-lite/auth/repository"
	"github.com/yt-lite/auth/service"
)

func AuthRouter(r *gin.RouterGroup) {
	h := handler.NewAuthHandler(service.NewAuthService(repository.NewAuthRepo(db.DB)))
	r.POST("/register", h.Register())
	r.POST("/login", h.Login())
}
