package api

import (
	"run/global"
	"run/models/constant"
	"run/models/request"
	"run/models/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RecordApi struct{}

func (r *RecordApi) CreateRecord(c *gin.Context) {
	var req request.RecordCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Log.Error(constant.RequestInvalid, zap.Error(err))
		response.FailWithMessage(constant.RequestInvalid, c)
		return
	}

	err := recordService.CreateRecord(req)
	if err != nil {
		global.Log.Error(constant.CreateFail, zap.Error(err))
		response.FailWithDetailed(req, constant.CreateFail, c)
		return
	}

	global.Log.Info(constant.CreateSuccess, zap.Any("req", req))
	response.OkWithMessage(constant.CreateSuccess, c)
}

func (r *RecordApi) UpdateRecord(c *gin.Context) {
	var req request.RecordUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Log.Error(constant.RequestInvalid, zap.Error(err))
		response.FailWithMessage(constant.RequestInvalid, c)
		return
	}

	err := recordService.UpdateRecord(req)
	if err != nil {
		global.Log.Error(constant.UpdateFail, zap.Error(err))
		response.FailWithDetailed(req, constant.UpdateFail, c)
		return
	}
	global.Log.Info(constant.UpdateSuccess, zap.Any("req", req), zap.Any("req", req))
	response.OkWithMessage(constant.UpdateSuccess, c)
}

func (r *RecordApi) DeleteRecord(c *gin.Context) {
	var ids []uint
	if err := c.ShouldBindJSON(&ids); err != nil {
		global.Log.Error(constant.RequestInvalid, zap.Error(err))
		response.FailWithMessage(constant.RequestInvalid, c)
		return
	}

	err := recordService.DeleteRecord(ids)
	if err != nil {
		global.Log.Error(constant.DeleteFail, zap.Error(err))
		response.FailWithDetailed(ids, constant.DeleteFail, c)
		return
	}
	global.Log.Info(constant.DeleteSuccess, zap.Uints("ids", ids))
	response.OkWithMessage(constant.DeleteSuccess, c)
}

func (r *RecordApi) ListRecords(c *gin.Context) {
	gameIdString := c.Query("id")
	if gameIdString == "" {
		global.Log.Error(constant.RequestInvalid, zap.Any("query", c))
		response.FailWithMessage(constant.RequestInvalid, c)
		return
	}

	gameIdUint64, err := strconv.ParseUint(gameIdString, 10, 64)
	if err != nil {
		global.Log.Error(constant.RequestInvalid, zap.Error(err))
		response.Fail(c)
		return
	}
	gameId := uint(gameIdUint64)

	records, err := recordService.ListRecords(gameId)
	if err != nil {
		global.Log.Error(constant.ListFail, zap.Error(err))
		response.FailWithDetailed(gameId, constant.ListFail, c)
		return
	}
	global.Log.Info(constant.ListSuccess, zap.Any("records", records))
	response.OkWithDetailed(records, constant.ListSuccess, c)
}
