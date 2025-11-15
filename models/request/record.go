package request

import (
	"run/models"
	"time"
)

type RecordCreateReq struct {
	GameId      uint      `json:"game_id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Hour        int16     `json:"hour" binding:"min=0"`
	Minute      int16     `json:"minute" binding:"min=0,max=59"`
	Second      int16     `json:"second" binding:"required,min=0,max=59"`
	Microsecond int16     `json:"microsecond" binding:"required,min=0,max=999"`
	Finish      time.Time `json:"finish"`
}

func (r RecordCreateReq) CreateRecord() *models.Record {
	return &models.Record{
		Name:        r.Name,
		Hour:        r.Hour,
		Minute:      r.Minute,
		Second:      r.Second,
		Microsecond: r.Microsecond,
		Finish:      r.Finish,
	}
}

type RecordUpdateReq struct {
	Id          uint      `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Hour        int16     `json:"hour" binding:"min=0"`
	Minute      int16     `json:"minute" binding:"min=0,max=59"`
	Second      int16     `json:"second" binding:"required,min=0,max=59"`
	Microsecond int16     `json:"microsecond" binding:"required,min=0,max=999"`
	Finish      time.Time `json:"finish"`
}
