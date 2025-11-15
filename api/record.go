package api

import (
	"run/global"
	"run/models/request"
	"run/models/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RecordApi struct{}

func (r *RecordApi) CreateRecord(c *gin.Context) {
	var req request.RecordCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Log.Error("request invalid", zap.Error(err))
		response.FailWithDetailed(req, "request invalid", c)
		return
	}

	err := recordService.CreateRecord(req)
	if err != nil {
		global.Log.Error("create record failed", zap.Error(err))
		response.FailWithDetailed(req, "create record failed", c)
		return
	}

	global.Log.Info("create record success", zap.Any("req", req))
	response.OkWithMessage("create record successfully", c)
}

func (r *RecordApi) UpdateRecord(c *gin.Context) {
	var req request.RecordUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Log.Error("request invalid", zap.Error(err))
		response.FailWithDetailed(req, "request invalid", c)
		return
	}

	err := recordService.UpdateRecord(req)
	if err != nil {
		global.Log.Error("update record failed", zap.Error(err))
		response.FailWithDetailed(req, "update record failed", c)
		return
	}
	global.Log.Info("update record success", zap.Any("req", req))
	response.OkWithMessage("update record successfully", c)
}

func (r *RecordApi) DeleteRecord(c *gin.Context) {
	var ids []uint
	if err := c.ShouldBindJSON(&ids); err != nil {
		global.Log.Error("parse record id failed", zap.Error(err))
		response.FailWithMessage("parse record id failed", c)
		return
	}

	err := recordService.DeleteRecord(ids)
	if err != nil {
		global.Log.Error("delete record failed", zap.Error(err))
		response.Fail(c)
		return
	}
	global.Log.Info("delete record success", zap.Uints("ids", ids))
	response.OkWithMessage("delete record successfully", c)
}

func (r *RecordApi) ListRecords(c *gin.Context) {
	var gameId uint
	if err := c.ShouldBindJSON(&gameId); err != nil {
		global.Log.Error("request invalid", zap.Error(err))
		response.FailWithMessage("request invalid", c)
		return
	}

	records, err := recordService.ListRecords(gameId)
	if err != nil {
		global.Log.Error("get record failed", zap.Error(err))
		response.FailWithMessage("get record failed", c)
		return
	}
	global.Log.Info("get record success", zap.Any("records", records))
	response.OkWithDetailed(records, "list records successfully", c)
}
