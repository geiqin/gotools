package model

import (
	"gorm.io/gorm"
)

type Timestamps struct {
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime:milli;comment:创建时间"` //创建时间
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime:milli;comment:修改时间"` //修改时间
}

type DelTimestamps struct {
	DeletedAt gorm.DeletedAt `sql:"index" json:"deleted_at" gorm:"type:datetime;comment:删除时间"`
}

type AllTimestamps struct {
	UpdatedAt int64          `json:"updated_at" gorm:"autoUpdateTime:milli;comment:创建时间"` //创建时间
	CreatedAt int64          `json:"created_at" gorm:"autoCreateTime:milli;comment:修改时间"` //修改时间
	DeletedAt gorm.DeletedAt `sql:"index" json:"deleted_at" gorm:"type:datetime;comment:删除时间"`
}
