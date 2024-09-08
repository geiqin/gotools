package model

import (
	"gorm.io/gorm"
	"time"
)

type Timestamps struct {
	CreatedAt LocalTime `json:"created_at" gorm:"type:datetime;autoCreateTime;comment:创建时间"`
	UpdatedAt LocalTime `json:"updated_at" gorm:"type:datetime;autoUpdateTime;comment:修改时间"`
}

type DelTimestamps struct {
	DeletedAt gorm.DeletedAt `sql:"index" json:"deleted_at" gorm:"type:datetime;comment:删除时间"`
}

type OldTimestamps struct {
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}

type AllTimestamps struct {
	CreatedAt time.Time      `json:"created_at" gorm:"type:datetime;comment:创建时间"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:datetime;comment:修改时间"`
	DeletedAt gorm.DeletedAt `sql:"index" json:"deleted_at" gorm:"type:datetime;comment:删除时间"`
}
