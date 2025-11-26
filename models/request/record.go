package request

import (
	"run/models"
	"time"

	"github.com/go-playground/validator/v10"
)

// TimeComponentsNotAllZero
// 自定义验证函数
// 函数签名必须是 func(fl validator.FieldLevel) bool
func TimeComponentsNotAllZero(fl validator.FieldLevel) bool {
	// 获取当前正在验证的结构体实例
	record := fl.Parent().Interface().(RecordCreateReq)

	// 检查 Hour, Minute, Second, Microsecond 是否都为 0
	return record.Hour != 0 || record.Minute != 0 || record.Second != 0 || record.Microsecond != 0
}

type RecordCreateReq struct {
	GameId      uint      `json:"game_id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Hour        int16     `json:"hour" binding:"min=0"`
	Minute      int16     `json:"minute" binding:"min=0,max=59"`
	Second      int16     `json:"second" binding:"min=0,max=59"`
	Microsecond int16     `json:"microsecond" binding:"min=0,max=999"`
	Finish      time.Time `json:"finish"`
	// 在结构体标签上添加自定义验证
	_ struct{} `binding:"timeComponentsNotAllZero"` // 使用一个匿名字段来附加验证规则到整个结构体
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

// TimeComponentsNotZero
// 自定义验证函数
// 函数签名必须是 func(fl validator.FieldLevel) bool
func TimeComponentsNotZero(fl validator.FieldLevel) bool {
	// 获取当前正在验证的结构体实例
	record := fl.Parent().Interface().(RecordUpdateReq)

	// 检查 Hour, Minute, Second, Microsecond 是否都为 0
	return record.Hour != 0 || record.Minute != 0 || record.Second != 0 || record.Microsecond != 0
}

type RecordUpdateReq struct {
	Id          uint      `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Hour        int16     `json:"hour" binding:"min=0"`
	Minute      int16     `json:"minute" binding:"min=0,max=59"`
	Second      int16     `json:"second" binding:"min=0,max=59"`
	Microsecond int16     `json:"microsecond" binding:"min=0,max=999"`
	Finish      time.Time `json:"finish"`
	_           struct{}  `binding:"timeComponentsNotZero"`
}
