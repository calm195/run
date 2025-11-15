package response

import (
	"run/models"
	"run/models/constant"
	"time"
)

type GameWebViewRsp struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Type      int16     `json:"type"`
	TypeName  string    `json:"type_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *GameWebViewRsp) CreateWebViewRsp(game models.Game) {
	r.Id = game.Id
	r.Name = game.Name
	r.Type = game.Type
	r.TypeName = constant.GetGameTypeName(game.Type)
	r.CreatedAt = game.CreatedAt
	r.UpdatedAt = game.UpdatedAt
}
