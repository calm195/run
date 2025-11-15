package models

import "time"

type Record struct {
	Base
	Name        string    `gorm:"column:name;comment:成绩拥有者姓名" json:"name"`
	Hour        int16     `gorm:"column:hour;comment:小时 " json:"hour"`
	Minute      int16     `gorm:"column:minute;comment:分钟" json:"minute"`
	Second      int16     `gorm:"column:second;comment:秒" json:"second"`
	Microsecond int16     `gorm:"column:microsecond;comment:微秒" json:"microsecond"`
	Finish      time.Time `gorm:"column:finish;comment:本成绩完赛时间" json:"finish"`
}

func (Record) TableName() string {
	return "record"
}

type Game struct {
	Base
	Name string `gorm:"column:name;comment:赛事名称或成绩解释" json:"name"`
	Type int16  `gorm:"column:type;comment:项目类型" json:"type"`
}

func (Game) TableName() string {
	return "game"
}

type GameRecord struct {
	Base
	GameId   uint `gorm:"column:game_id;comment:项目/赛事主键id" json:"game_id"`
	RecordId uint `gorm:"column:record_id;comment:单次数据主键id" json:"record_id"`
}

func (GameRecord) TableName() string {
	return "game_record"
}
