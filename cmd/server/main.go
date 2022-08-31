package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yt-lite/auth/config"
	"github.com/yt-lite/auth/db"
	"github.com/yt-lite/libs/logger"
)

func main() {
	// ...
	gin.SetMode(gin.DebugMode)
	config.Init()
	db.Init()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	logger.Fatal(run(r).Error())
}

func run(r *gin.Engine) error {

	logger.Infof("Server started on port %s", config.Cfg.HttpAuthPort)
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Cfg.HttpAuthPort), r)
}
