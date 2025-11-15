package models

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	Id        uint           `gorm:"column:id;comment:主键ID" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
