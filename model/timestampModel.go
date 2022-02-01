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

//创建者
type Creator struct {
	CreatorId   int64  `json:"creator_id" gorm:"default:0"` //创建者ID
	CreatorName string `json:"creator_name" gorm:"size:50"` //创建者名称
}

//修改者
type Modifier struct {
	ModifierId   int64  `json:"modifier_id" gorm:"default:0"` //修改者ID
	ModifierName string `json:"modifier_name" gorm:"size:50"` //修改者名称
}
