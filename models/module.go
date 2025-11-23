package models

import (
	"run/models/types"
	"time"
)

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

type Event struct {
	Base
	Name     string `gorm:"column:name;comment:运动项目名称" json:"name"`
	Distance int32  `gorm:"column:distance;comment:距离" json:"distance"`
}

func (Event) TableName() string {
	return "event"
}

type Standard struct {
	Base
	EventID        uint                 `gorm:"not null;index:idx_standards_lookup,priority:1;comment:项目ID" json:"event_id"`
	Gender         types.Gender         `gorm:"type:smallint;not null;index:idx_standards_lookup,priority:2;comment:性别 1=男 2=女" json:"gender"`
	Level          types.Level          `gorm:"type:smallint;not null;index:idx_standards_lookup,priority:4;comment:等级 1=健将 2=一级 3=二级 4=三级 5=参与级" json:"level"`
	Threshold      float64              `gorm:"type:double precision;not null;comment:达标成绩（秒），成绩 ≤ 该值即达标" json:"threshold"`
	StandardSystem types.StandardSystem `gorm:"type:smallint;not null;default:2;index:idx_standards_lookup,priority:3;comment:标准体系 1=体测 2=中国 3=国际" json:"standard_system"`
}

func (Standard) TableName() string {
	return "standard"
}
