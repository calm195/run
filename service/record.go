package service

import (
	"errors"
	"run/global"
	"run/models"
	"run/models/constant"
	"run/models/request"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RecordService struct{}

func (r *RecordService) CreateRecord(recordCreateReq request.RecordCreateReq) (err error) {
	var game models.Game
	if err = global.Db.First(&game, recordCreateReq.GameId).Error; err != nil {
		global.Log.Error(constant.NotExist, zap.Uint("game id", recordCreateReq.GameId), zap.Error(err))
		return
	}

	tx := global.Db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else if r := recover(); r != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
			if err != nil {
				global.Log.Error(constant.CommitFail, zap.Error(err))
			} else {
				global.Log.Info(constant.CommitSuccess)
			}
		}
	}()

	record := recordCreateReq.CreateRecord()
	if err = tx.Create(record).Error; err != nil {
		global.Log.Error(constant.CreateFail, zap.Any("record", record), zap.Error(err))
		return
	}

	gameRecord := models.GameRecord{
		GameId:   game.Id,
		RecordId: record.Id,
	}
	if err = tx.Create(&gameRecord).Error; err != nil {
		global.Log.Error(constant.CreateFail, zap.Any("gameRecord", gameRecord), zap.Error(err))
		return
	}

	return
}

func (r *RecordService) UpdateRecord(req request.RecordUpdateReq) error {
	result := global.Db.Model(&models.Record{}).
		Where("id = ?", req.Id).
		Updates(req)
	if result.Error != nil {
		global.Log.Error(constant.UpdateFail, zap.Uint("record id", req.Id), zap.Error(result.Error))
		return result.Error
	}
	if result.RowsAffected == 0 {
		global.Log.Warn(constant.NotExist, zap.Uint("record id", req.Id))
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *RecordService) ListRecords(gameId uint) (records []models.Record, err error) {
	var game models.Game
	err = global.Db.Select("id").Where("id = ?", gameId).Take(&game).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 不存在
			global.Log.Warn(constant.NotExist, zap.Uint("game id", gameId))
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	err = global.Db.
		Table(models.Record{}.TableName()).
		Joins("INNER JOIN game_record ON record.id = game_record.record_id").
		Where("game_record.game_id = ?", gameId).
		Find(&records).Error

	if err != nil {
		global.Log.Error(constant.QueryFail, zap.Uint("game id", gameId), zap.Error(err))
		return nil, err
	}

	return records, nil
}

func (r *RecordService) DeleteRecord(ids []uint) (err error) {
	if len(ids) == 0 {
		return nil
	}

	tx := global.Db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else {
			err = tx.Commit().Error
			if err != nil {
				global.Log.Error(constant.CommitFail, zap.Error(err))
			} else {
				global.Log.Info(constant.DeleteSuccess)
			}
		}
	}()

	if err = tx.Where("id IN ?", ids).Delete(&models.Record{}).Error; err != nil {
		global.Log.Error(constant.DeleteFail, zap.Uints("record ids", ids), zap.Error(err))
		return
	}

	if err = tx.Where("record_id IN ?", ids).Delete(&models.GameRecord{}).Error; err != nil {
		global.Log.Error(constant.DeleteFail, zap.Uints("record ids", ids), zap.Error(err))
		return
	}

	return
}
