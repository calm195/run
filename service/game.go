package service

import (
	"errors"
	"run/global"
	"run/models"
	"run/models/constant"
	"run/models/request"
	"run/models/response"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GameService struct{}

func (g *GameService) CreateGame(game request.GameCreateReq) (err error) {
	if constant.IfGameTypeNotExist(game.Type) {
		global.Log.Error(constant.GameTypeInvalid)
		return errors.New(constant.GameTypeInvalid)
	}
	err = global.Db.Create(game.CreateGame()).Error
	return err
}

func (g *GameService) UpdateGame(game request.GameUpdateReq) (err error) {
	if constant.IfGameTypeNotExist(game.Type) {
		global.Log.Error(constant.GameTypeInvalid)
		return errors.New(constant.GameTypeInvalid)
	}

	var repo models.Game
	if err = global.Db.First(&repo, game.Id).Error; err != nil {
		global.Log.Error(constant.NotExist, zap.Uint("game id", game.Id), zap.Error(err))
		return
	}

	err = global.Db.Model(&repo).Updates(game).Error
	return
}

func (g *GameService) ListAllGames() (gameWebViewRspList []response.GameWebViewRsp, err error) {
	var repos []models.Game
	err = global.Db.Find(&repos).Error
	for _, item := range repos {
		var gameWebViewRsp response.GameWebViewRsp
		gameWebViewRsp.CreateWebViewRsp(item)
		gameWebViewRspList = append(gameWebViewRspList, gameWebViewRsp)
	}
	return gameWebViewRspList, err
}

func (g *GameService) GetGameById(id uint) (gameWebViewRsp response.GameWebViewRsp, err error) {
	var repo models.Game
	err = global.Db.First(&repo, id).Error
	gameWebViewRsp.CreateWebViewRsp(repo)
	return gameWebViewRsp, err
}

func (g *GameService) DeleteGame(ids []uint) (err error) {
	var games []models.Game
	if err = global.Db.Find(&games, ids).Error; err != nil {
		global.Log.Error(constant.FindFail, zap.Uints("ids", ids), zap.Error(err))
		return
	}

	var gameRecordList []models.GameRecord
	if err = global.Db.Where("game_id in (?)", ids).Find(&gameRecordList).Error; err != nil {
		global.Log.Error(constant.FindFail, zap.Uints("game ids", ids), zap.Error(err))
	}

	var recordIds []uint
	for _, gameRecord := range gameRecordList {
		recordIds = append(recordIds, gameRecord.RecordId)
	}

	return global.Db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Delete(&games).Error; err != nil {
			global.Log.Error(constant.DeleteFail, zap.Error(err))
			return err
		}
		if len(gameRecordList) <= 0 {
			return nil
		}
		if err = tx.Delete(&gameRecordList).Error; err != nil {
			global.Log.Error(constant.DeleteFail, zap.Error(err))
			return err
		}
		if len(recordIds) <= 0 {
			return nil
		}
		if err = tx.Delete(&models.Record{}, recordIds).Error; err != nil {
			global.Log.Error(constant.DeleteFail, zap.Error(err))
			return err
		}
		return nil
	})
}

func (g *GameService) GetRecordNum(id uint) (num int64, err error) {
	var game models.Game
	if err = global.Db.First(&game, id).Error; err != nil {
		global.Log.Error(constant.FindFail, zap.Error(err))
		num = -1
		return
	}

	err = global.Db.Model(&models.GameRecord{}).Where("game_id=?", game.Id).Count(&num).Error
	return
}
