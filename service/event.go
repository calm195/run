package service

import (
	"run/global"
	"run/models"
	"run/models/constant"

	"go.uber.org/zap"
)

type EventService struct{}

func (e *EventService) ListAllEvents() ([]models.Event, error) {
	var events []models.Event
	if err := global.Db.Find(&events).Error; err != nil {
		global.Log.Error(constant.QueryFail, zap.Error(err))
		return nil, err
	}

	return events, nil
}
