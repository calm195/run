package api

import (
	"run/global"
	"run/models/request"
	"run/models/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GameApi struct{}

func (g *GameApi) CreateGame(c *gin.Context) {
	var req request.GameCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Log.Error("request invalid", zap.Error(err))
		response.FailWithDetailed(req, "request invalid", c)
		return
	}

	err := gameService.CreateGame(req)
	if err != nil {
		global.Log.Error("create game failed", zap.Error(err))
		response.FailWithMessage("create game failed", c)
		return
	}
	global.Log.Info("create game success", zap.Any("req", req))
	response.OkWithMessage("create game successfully", c)
}

func (g *GameApi) UpdateGame(c *gin.Context) {
	var req request.GameUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Log.Error("update game failed", zap.Error(err))
		response.FailWithDetailed(req, "update game failed", c)
		return
	}

	err := gameService.UpdateGame(req)
	if err != nil {
		global.Log.Error("update game failed", zap.Error(err))
		response.FailWithDetailed(req, "update game failed", c)
		return
	}
	global.Log.Info("update game success", zap.Any("req", req))
	response.OkWithMessage("update game successfully", c)
}

func (g *GameApi) ListAllGames(c *gin.Context) {
	games, err := gameService.ListAllGames()
	if err != nil {
		global.Log.Error("list games failed", zap.Error(err))
		response.FailWithMessage("list games failed", c)
		return
	}
	response.OkWithData(games, c)
}

func (g *GameApi) DeleteGame(c *gin.Context) {
	var ids []uint
	if err := c.ShouldBindJSON(&ids); err != nil {
		global.Log.Error("delete game failed", zap.Error(err))
		response.FailWithMessage("delete game failed", c)
		return
	}

	err := gameService.DeleteGame(ids)
	if err != nil {
		global.Log.Error("delete game failed", zap.Error(err))
		response.FailWithMessage("delete game failed", c)
		return
	}
	global.Log.Info("delete game success", zap.Any("ids", ids))
	response.OkWithMessage("delete game successfully", c)
}

func (g *GameApi) GetRecordNum(c *gin.Context) {
	var id uint
	if err := c.ShouldBindJSON(&id); err != nil {
		global.Log.Error("game id invalid", zap.Error(err))
		response.FailWithMessage("game id invalid", c)
		return
	}

	num, err := gameService.GetRecordNum(id)
	if err != nil {
		global.Log.Error("get game record num failed", zap.Error(err))
		response.FailWithMessage("get game record num failed", c)
		return
	}
	global.Log.Info("get game record num success", zap.Int64("num", num))
	response.OkWithDetailed(num, "get game record num successfully", c)
}
