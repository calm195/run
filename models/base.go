package models

import "time"

type Base struct {
	Id         int64     `gorm:"column:primaryKey;id" json:"id"`        //type:int64       comment:主键ID,;             version:2025-10-13 14:07
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` //type:time.Time   comment:创建时间,;           version:2025-10-13 14:07
}
