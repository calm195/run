package models

import "time"

type Base struct {
	Id         int64     `gorm:"column:id;comment:主键ID" json:"id"`
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`
}
