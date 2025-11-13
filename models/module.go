package models

import "time"

type Record struct {
	Base
	Name        string    `gorm:"column:name" json:"name"`               //type:string      comment:成绩拥有者姓名,;           version:2025-10-13 11:49
	Hour        int16     `gorm:"column:hour" json:"hour"`               //type:int16       comment:小时,;                     version:2025-10-13 11:49
	Minute      int16     `gorm:"column:minute" json:"minute"`           //type:int16       comment:分钟,;                     version:2025-10-13 11:49
	Second      int16     `gorm:"column:second" json:"second"`           //type:int16       comment:秒,;                       version:2025-10-13 11:49
	Microsecond int16     `gorm:"column:microsecond" json:"microsecond"` //type:int16       comment:微秒,;                     version:2025-10-13 11:49
	Finish      time.Time `gorm:"column:finish" json:"finish"`           //type:time.Time   comment:本成绩创造时间,;           version:2025-10-13 14:19
}

type Game struct {
	Base
	Name string `gorm:"column:name" json:"name"` //type:string      comment:赛事名称或成绩解释,;           version:2025-10-13 14:12
	Type int16  `gorm:"column:type" json:"type"` //type:int16       comment:项目类型,;                     version:2025-10-13 14:12
}

type GameRecord struct {
	Base
	GameId   int64 `gorm:"column:game_id" json:"game_id"`     //type:int64       comment:项目/赛事主键id,;           version:2025-10-13 14:12
	RecordId int64 `gorm:"column:record_id" json:"record_id"` //type:int64       comment:单次数据主键id,;            version:2025-10-13 14:12
}
