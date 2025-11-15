package request

import "run/models"

type GameCreateReq struct {
	Name string `json:"name" binding:"required"`
	Type int16  `json:"type" binding:"required"`
}

func (req GameCreateReq) CreateGame() *models.Game {
	return &models.Game{
		Name: req.Name,
		Type: req.Type,
	}
}

type GameUpdateReq struct {
	Id   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Type int16  `json:"type" binding:"required"`
}

func (req GameUpdateReq) CreateGame() *models.Game {
	return &models.Game{
		Name: req.Name,
		Type: req.Type,
	}
}
