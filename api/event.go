package api

import (
	"run/global"
	"run/models/constant"
	"run/models/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EventApi struct{}

func (e *EventApi) ListAllEvents(c *gin.Context) {
	events, err := eventService.ListAllEvents()
	if err != nil {
		global.Log.Error(constant.ListFail, zap.Error(err))
		response.FailWithMessage(constant.ListFail, c)
		return
	}
	global.Log.Info(constant.ListSuccess, zap.Any("events", events))
	response.OkWithDetailed(events, constant.ListSuccess, c)
}
