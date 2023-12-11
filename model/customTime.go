package model

import (
	"database/sql/driver"
	"github.com/geiqin/gotools/helper"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type CustomTime string

func (t CustomTime) GormDataType() string {
	return "time"
}
func (t CustomTime) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return "time"
}

func (t CustomTime) Value() (driver.Value, error) {
	if t == "" {
		return nil, nil
	}
	return []byte(t), nil
}

func (t *CustomTime) Scan(v interface{}) error {
	str := helper.ToString(v)
	*t = CustomTime(str)
	return nil
}
