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

func (g *GameService) CreateGame(req request.GameCreateReq) (err error) {
	if constant.IfGameTypeNotExist(req.Type) {
		global.Log.Error(constant.GameTypeInvalid, zap.Int16("type", req.Type))
		return errors.New(constant.GameTypeInvalid)
	}

	gameModel := req.CreateGame()
	if gameModel == nil {
		err = errors.New("CreateGame returned nil")
		global.Log.Error("CreateGame returned nil", zap.Any("req", req))
		return err
	}

	err = global.Db.Create(gameModel).Error
	if err != nil {
		global.Log.Error(constant.CreateFail,
			zap.Int16("type", req.Type),
			zap.Any("game", gameModel),
			zap.Error(err))
	}
	return err
}

func (g *GameService) UpdateGame(req request.GameUpdateReq) error {
	if req.Id == 0 {
		err := errors.New("game id is required")
		global.Log.Error("game id is zero", zap.Any("req", req))
		return err
	}

	if constant.IfGameTypeNotExist(req.Type) {
		global.Log.Error(constant.GameTypeInvalid, zap.Int16("type", req.Type))
		return errors.New(constant.GameTypeInvalid)
	}

	result := global.Db.Model(&models.Game{}).
		Where("id = ?", req.Id).
		Omit("id").
		Updates(req)

	if result.Error != nil {
		global.Log.Error(constant.UpdateFail,
			zap.Uint("game_id", req.Id),
			zap.Int16("type", req.Type),
			zap.Error(result.Error))
		return result.Error
	}

	if result.RowsAffected == 0 {
		global.Log.Warn(constant.NotExist, zap.Uint("game_id", req.Id))
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (g *GameService) ListAllGames() ([]response.GameWebViewRsp, error) {
	var repos []models.Game
	if err := global.Db.Find(&repos).Error; err != nil {
		global.Log.Error(constant.QueryFail, zap.Error(err))
		return nil, err
	}

	rspList := make([]response.GameWebViewRsp, len(repos))
	for i, item := range repos {
		rspList[i].CreateWebViewRsp(item)
	}
	return rspList, nil
}

func (g *GameService) GetGameById(id uint) (response.GameWebViewRsp, error) {
	if id == 0 {
		return response.GameWebViewRsp{}, gorm.ErrRecordNotFound
	}

	var repo models.Game
	err := global.Db.First(&repo, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.Log.Warn(constant.NotExist, zap.Uint("game_id", id))
		} else {
			global.Log.Error(constant.FindFail, zap.Uint("game_id", id), zap.Error(err))
		}
		return response.GameWebViewRsp{}, err
	}

	var rsp response.GameWebViewRsp
	rsp.CreateWebViewRsp(repo)
	return rsp, nil
}

func (g *GameService) DeleteGame(ids []uint) error {
	if len(ids) == 0 {
		global.Log.Warn("attempt to delete games with empty id list")
		return nil
	}

	validIDs := make([]uint, 0, len(ids))
	for _, id := range ids {
		if id == 0 {
			global.Log.Warn("skip deleting game with id=0")
			continue
		}
		validIDs = append(validIDs, id)
	}
	if len(validIDs) == 0 {
		return nil
	}

	return global.Db.Transaction(func(tx *gorm.DB) error {
		var recordIDs []uint
		if err := tx.Model(&models.GameRecord{}).
			Where("game_id IN ?", validIDs).
			Pluck("record_id", &recordIDs).Error; err != nil {
			global.Log.Error(constant.FindFail, zap.Uints("game_ids", validIDs), zap.Error(err))
			return err
		}

		if err := tx.Where("game_id IN ?", validIDs).Delete(&models.GameRecord{}).Error; err != nil {
			global.Log.Error(constant.DeleteFail, zap.String("table", "game_records"), zap.Error(err))
			return err
		}

		if len(recordIDs) > 0 {
			if err := tx.Delete(&models.Record{}, recordIDs).Error; err != nil {
				global.Log.Error(constant.DeleteFail, zap.String("table", "records"), zap.Uints("record_ids", recordIDs), zap.Error(err))
				return err
			}
			global.Log.Info("records deleted due to game deletion", zap.Int("count", len(recordIDs)), zap.Uints("game_ids", validIDs))
		}

		result := tx.Delete(&models.Game{}, validIDs)
		if result.Error != nil {
			global.Log.Error(constant.DeleteFail, zap.String("table", "game"), zap.Error(result.Error))
			return result.Error
		}

		global.Log.Info("games deleted successfully",
			zap.Uints("game_ids", validIDs),
			zap.Int64("games_deleted", result.RowsAffected),
			zap.Int("records_deleted", len(recordIDs)))

		return nil
	})
}

func (g *GameService) GetRecordNum(gameID uint) (int64, error) {
	if gameID == 0 {
		return 0, errors.New("game id is required")
	}

	var count int64
	err := global.Db.Model(&models.GameRecord{}).
		Where("game_id = ?", gameID).
		Count(&count).Error

	if err != nil {
		global.Log.Error(constant.QueryFail, zap.Uint("game_id", gameID), zap.Error(err))
		return 0, err
	}

	return count, nil
}
