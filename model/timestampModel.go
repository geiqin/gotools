package model

import (
	"gorm.io/gorm"
	"time"
)

type Timestamps struct {
	CreatedAt LocalTime `json:"created_at" gorm:"type:datetime;autoCreateTime"`
	UpdatedAt LocalTime `json:"updated_at" gorm:"type:datetime;autoUpdateTime"`
}

type DelTimestamps struct {
	DeletedAt gorm.DeletedAt `sql:"index" json:"deleted_at" gorm:"type:datetime"`
}

type OldTimestamps struct {
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}

type AllTimestamps struct {
	CreatedAt time.Time      `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:datetime"`
	DeletedAt gorm.DeletedAt `sql:"index" json:"deleted_at" gorm:"type:datetime"`
}
