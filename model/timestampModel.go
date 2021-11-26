package model

import (
	"gorm.io/gorm"
	"time"
)

type Timestamps struct {
	CreatedAt LocalTime `json:"created_at" gorm:"type:datetime"`
	UpdatedAt LocalTime `json:"updated_at" gorm:"type:datetime"`
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

/*
//MyTime 自定义时间
type MyTime time.Time

func (t *MyTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = MyTime(t1)
	return err
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t MyTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t *MyTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = MyTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *MyTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}

*/
