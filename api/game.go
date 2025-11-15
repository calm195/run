package api

import (
	"run/global"
	"run/models/constant"
	"run/models/request"
	"run/models/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GameApi struct{}

func (g *GameApi) CreateGame(c *gin.Context) {
	var req request.GameCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Log.Error(constant.RequestInvalid, zap.Error(err))
		response.FailWithDetailed(req, constant.RequestInvalid, c)
		return
	}

	err := gameService.CreateGame(req)
	if err != nil {
		global.Log.Error(constant.CreateFail, zap.Error(err))
		response.FailWithDetailed(req, constant.CreateFail, c)
		return
	}
	global.Log.Info(constant.CreateSuccess, zap.Any("req", req))
	response.OkWithMessage(constant.CreateSuccess, c)
}

func (g *GameApi) UpdateGame(c *gin.Context) {
	var req request.GameUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Log.Error(constant.RequestInvalid, zap.Error(err))
		response.FailWithMessage(constant.RequestInvalid, c)
		return
	}

	err := gameService.UpdateGame(req)
	if err != nil {
		global.Log.Error(constant.UpdateFail, zap.Error(err))
		response.FailWithDetailed(req, constant.UpdateFail, c)
		return
	}
	global.Log.Info(constant.UpdateSuccess, zap.Any("req", req))
	response.OkWithMessage(constant.UpdateSuccess, c)
}

func (g *GameApi) ListAllGames(c *gin.Context) {
	games, err := gameService.ListAllGames()
	if err != nil {
		global.Log.Error(constant.ListFail, zap.Error(err))
		response.FailWithMessage(constant.ListFail, c)
		return
	}
	global.Log.Info(constant.ListSuccess, zap.Any("games", games))
	response.OkWithDetailed(games, constant.ListSuccess, c)
}

func (g *GameApi) DeleteGame(c *gin.Context) {
	var ids []uint
	if err := c.ShouldBindJSON(&ids); err != nil {
		global.Log.Error(constant.RequestInvalid, zap.Error(err))
		response.FailWithMessage(constant.RequestInvalid, c)
		return
	}

	err := gameService.DeleteGame(ids)
	if err != nil {
		global.Log.Error(constant.DeleteFail, zap.Error(err))
		response.FailWithDetailed(ids, constant.DeleteFail, c)
		return
	}
	global.Log.Info(constant.DeleteSuccess, zap.Any("ids", ids))
	response.OkWithMessage(constant.DeleteSuccess, c)
}

func (g *GameApi) GetRecordNum(c *gin.Context) {
	var id uint
	if err := c.ShouldBindJSON(&id); err != nil {
		global.Log.Error(constant.RequestInvalid, zap.Error(err))
		response.FailWithMessage(constant.RequestInvalid, c)
		return
	}

	num, err := gameService.GetRecordNum(id)
	if err != nil {
		global.Log.Error(constant.CountFail, zap.Error(err))
		response.FailWithDetailed(id, constant.CountFail, c)
		return
	}
	global.Log.Info(constant.CreateSuccess, zap.Int64("num", num))
	response.OkWithDetailed(num, constant.CreateSuccess, c)
}
