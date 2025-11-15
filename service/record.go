package service

import (
	"run/global"
	"run/models"
	"run/models/request"

	"go.uber.org/zap"
)

type RecordService struct{}

func (r *RecordService) CreateRecord(recordCreateReq request.RecordCreateReq) (err error) {
	var game models.Game
	err = global.Db.First(&game, recordCreateReq.GameId).Error
	if err != nil {
		global.Log.Error("game id invalid", zap.Error(err))
		return
	}

	tx := global.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	record := recordCreateReq.CreateRecord()
	if err = tx.Create(record).Error; err != nil {
		tx.Rollback()
		global.Log.Error("create record failed", zap.Any("record", record), zap.Error(err))
	}

	gameRecord := models.GameRecord{
		GameId:   game.Id,
		RecordId: record.Id,
	}
	if err = tx.Create(&gameRecord).Error; err != nil {
		tx.Rollback()
		global.Log.Error("create game record failed", zap.Any("gameRecord", gameRecord), zap.Error(err))
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		global.Log.Error("commit record failed", zap.Error(err))
	}

	global.Log.Info("record created", zap.Any("record", record), zap.Any("gameRecord", gameRecord))
	return
}

func (r *RecordService) UpdateRecord(recordUpdateReq request.RecordUpdateReq) (err error) {
	var repo models.Record
	if err = global.Db.First(&repo, recordUpdateReq.Id).Error; err != nil {
		global.Log.Error("record id invalid", zap.Any("record", recordUpdateReq), zap.Error(err))
		return
	}

	err = global.Db.Model(&repo).Updates(recordUpdateReq).Error
	if err != nil {
		global.Log.Error("update record failed", zap.Any("record", recordUpdateReq), zap.Error(err))
		return
	}
	return
}

func (r *RecordService) ListRecords(gameId uint) (records []models.Record, err error) {
	if err = global.Db.First(&models.Game{}, gameId).Error; err != nil {
		global.Log.Error("game not found", zap.Uint("game id", gameId), zap.Error(err))
		return
	}
	var gameRecordList []models.GameRecord
	if err = global.Db.Where("game_id = ?", gameId).Find(&gameRecordList).Error; err != nil {
		global.Log.Error("found game record error", zap.Uint("game id", gameId), zap.Error(err))
		return
	}

	var recordIds []uint
	for _, gameRecord := range gameRecordList {
		recordIds = append(recordIds, gameRecord.RecordId)
	}
	if err = global.Db.Find(&records, recordIds).Error; err != nil {
		global.Log.Error("found record error", zap.Any("gameRecordList", gameRecordList), zap.Error(err))
		return
	}
	return
}

func (r *RecordService) DeleteRecord(ids []uint) (err error) {
	tx := global.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Delete(&models.Record{}, ids).Error; err != nil {
		tx.Rollback()
		global.Log.Error("delete record failed", zap.Uints("record id", ids), zap.Error(err))
		return
	}

	if err = tx.Delete(&models.GameRecord{}, "record_id IN (?)", ids).Error; err != nil {
		tx.Rollback()
		global.Log.Error("found game record error", zap.Uints("record id", ids), zap.Error(err))
		return
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		global.Log.Error("commit record failed", zap.Error(err))
	}
	return
}
