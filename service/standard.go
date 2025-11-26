package service

import (
	"run/global"
	"run/models"
	"run/models/constant"

	"go.uber.org/zap"
)

type StandardService struct{}

func (s *StandardService) ListAllStandard() ([]models.Standard, error) {
	var standards []models.Standard
	if err := global.Db.Find(&standards).Error; err != nil {
		global.Log.Error(constant.QueryFail, zap.Error(err))
		return nil, err
	}
	return standards, nil
}
