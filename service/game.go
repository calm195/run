package service

import (
	"errors"
	"run/global"
	"run/models"
	"run/models/constant"
	"run/models/request"
	"run/models/response"

	"go.uber.org/zap"
)

type GameService struct{}

func (g *GameService) CreateGame(game request.GameCreateReq) (err error) {
	if constant.IfGameTypeNotExist(game.Type) {
		global.Log.Error("game type not exist")
		return errors.New("game type not exist")
	}
	err = global.Db.Create(game.CreateGame()).Error
	return err
}

func (g *GameService) UpdateGame(game request.GameUpdateReq) (err error) {
	if constant.IfGameTypeNotExist(game.Type) {
		global.Log.Error("game type not exist")
		return errors.New("game type not exist")
	}

	var repo models.Game
	err = global.Db.First(&repo, game.Id).Error
	if err != nil {
		global.Log.Error("the game record not exits", zap.Error(err))
		return errors.New("the game record not exits")
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

func (g *GameService) DeleteGame(ids []uint) (err error) {
	err = global.Db.Delete(&models.Game{}, ids).Error
	return err
}

func (g *GameService) GetRecordNum(id uint) (num int64, err error) {
	var game models.Game
	err = global.Db.First(&game, id).Error
	if err != nil {
		global.Log.Error("get game record num failed", zap.Error(err))
		num = -1
		return
	}

	err = global.Db.Model(&models.GameRecord{}).Where("game_id=?", game.Id).Count(&num).Error
	return
}
