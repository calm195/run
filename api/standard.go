package api

import (
	"run/global"
	"run/models/constant"
	"run/models/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StandardApi struct{}

func (s *StandardApi) ListAllStandard(c *gin.Context) {
	standards, err := standardService.ListAllStandard()
	if err != nil {
		global.Log.Error(constant.ListFail, zap.Error(err))
		response.FailWithMessage(constant.ListFail, c)
		return
	}
	global.Log.Info(constant.ListSuccess, zap.Any("standards", standards))
	response.OkWithDetailed(standards, constant.ListSuccess, c)
}
