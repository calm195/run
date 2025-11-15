package core

import (
	"net/http"
	"run/global"
	"run/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	if gin.IsDebugging() {
		engine.Use(gin.Logger())
	}

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	registerRouters(engine)

	return engine
}

func registerRouters(engine *gin.Engine) {
	rootGroup := engine.Group(global.Config.System.RouterPrefix)

	gameRouter := router.RootGroup.Game
	gameRouter.InitGameRouter(rootGroup)
}
