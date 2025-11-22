package service

import (
	"run/global"
	"run/models"
	"run/models/constant"
	"run/models/request"

	"go.uber.org/zap"
)

type RecordService struct{}

func (r *RecordService) CreateRecord(recordCreateReq request.RecordCreateReq) (err error) {
	var game models.Game
	err = global.Db.First(&game, recordCreateReq.GameId).Error
	if err != nil {
		global.Log.Error(constant.NotExist, zap.Uint("game id", recordCreateReq.GameId), zap.Error(err))
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
		global.Log.Error(constant.CreateFail, zap.Any("record", record), zap.Error(err))
	}

	gameRecord := models.GameRecord{
		GameId:   game.Id,
		RecordId: record.Id,
	}
	if err = tx.Create(&gameRecord).Error; err != nil {
		tx.Rollback()
		global.Log.Error(constant.CreateFail, zap.Any("gameRecord", gameRecord), zap.Error(err))
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		global.Log.Error(constant.CommitFail, zap.Error(err))
	}

	global.Log.Info(constant.CommitSuccess)
	return
}

func (r *RecordService) UpdateRecord(recordUpdateReq request.RecordUpdateReq) (err error) {
	var repo models.Record
	if err = global.Db.First(&repo, recordUpdateReq.Id).Error; err != nil {
		global.Log.Error(constant.NotExist, zap.Uint("record id", recordUpdateReq.Id), zap.Error(err))
		return
	}

	err = global.Db.Model(&repo).Updates(recordUpdateReq).Error
	return
}

func (r *RecordService) ListRecords(gameId uint) (records []models.Record, err error) {
	if err = global.Db.First(&models.Game{}, gameId).Error; err != nil {
		global.Log.Error(constant.NotExist, zap.Uint("game id", gameId), zap.Error(err))
		return
	}
	var gameRecordList []models.GameRecord
	if err = global.Db.Where("game_id = ?", gameId).Find(&gameRecordList).Error; err != nil {
		global.Log.Error(constant.NotExist, zap.Uint("game id", gameId), zap.Error(err))
		return
	}

	var recordIds []uint
	for _, gameRecord := range gameRecordList {
		recordIds = append(recordIds, gameRecord.RecordId)
	}
	if err = global.Db.Where("id IN ?", recordIds).Find(&records).Error; err != nil {
		global.Log.Error(constant.NotExist, zap.Uints("record ids", recordIds), zap.Error(err))
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
		global.Log.Error(constant.DeleteFail, zap.Uints("record id", ids), zap.Error(err))
		return
	}

	if err = tx.Delete(&models.GameRecord{}, "record_id IN (?)", ids).Error; err != nil {
		tx.Rollback()
		global.Log.Error(constant.DeleteFail, zap.Uints("record id", ids), zap.Error(err))
		return
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		global.Log.Error(constant.CommitFail, zap.Error(err))
	}
	global.Log.Info(constant.CreateSuccess)
	return
}
