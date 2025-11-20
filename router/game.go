package router

import "github.com/gin-gonic/gin"

type GameRouter struct{}

func (g *GameRouter) InitGameRouter(router *gin.RouterGroup) {
	gameRouter := router.Group("/game")

	gameRouter.POST("/create", gameApi.CreateGame)
	gameRouter.PUT("/edit", gameApi.UpdateGame)
	gameRouter.GET("/list", gameApi.ListAllGames)
	gameRouter.GET("/get", gameApi.GetGameById)
	gameRouter.DELETE("/delete", gameApi.DeleteGame)
	gameRouter.GET("/num", gameApi.GetRecordNum)
}
