package model

import (
	"database/sql/driver"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

type LocalTime time.Time

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}

	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = LocalTime(now)
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t LocalTime) Value() (driver.Value, error) {
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

func (t *LocalTime) Scan(v interface{}) error {
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = LocalTime(tTime)
	return nil
}

func (t LocalTime) String() string {
	return time.Time(t).Format(TimeFormat)
}

// *****EXTEND*****

//set time for localTime
func (t *LocalTime) SetTime(tTime time.Time) {
	*t = LocalTime(tTime)
}

func (t LocalTime) Year() int {
	return time.Time(t).Year()
}

func (t LocalTime) Month() time.Month {
	return time.Time(t).Month()
}

func (t LocalTime) Day() int {
	return time.Time(t).Day()
}

func (t LocalTime) Hour() int {
	return time.Time(t).Hour()
}
func (t LocalTime) Second() int {
	return time.Time(t).Second()
}

func (t LocalTime) Minute() int {
	return time.Time(t).Minute()
}

func (t LocalTime) Unix() int64 {
	return time.Time(t).Unix()
}

func (t LocalTime) UnixNano() int64 {
	return time.Time(t).UnixNano()
}

//获取时间
//timeType值为1时只取年,
//timeType值为2时只取年和月,
//timeType值为3时取只年和月和日,

func (t LocalTime) GetTime() time.Time {
	return time.Time(t)
}

//获取时间 年-月-日
func (t LocalTime) GetTimeOfDate() time.Time {
	timeStr := time.Time(t).Format("2006-01-02")
	t.Scan(timeStr)
	return t.GetTime()
}

func (t LocalTime) Add(d time.Duration) time.Time {
	return time.Time(t).Add(d)
}

// AddDate returns the time corresponding to adding the
// given number of years, months, and days to t.
// For example, AddDate(-1, 2, 3) applied to January 1, 2011
// returns March 4, 2010.
//
// AddDate normalizes its result in the same way that Date does,
// so, for example, adding one month to October 31 yields
// December 1, the normalized form for November 31.
func (t LocalTime) AddDate(years int, months int, days int) time.Time {
	return time.Time(t).AddDate(years, months, days)
}

//添加天数
func (t LocalTime) AddDay(days int) time.Time {
	return time.Time(t).AddDate(0, 0, days)
}

//添加小时
func (t LocalTime) AddHour(hour int) time.Time {
	var dur time.Duration = time.Duration(hour) * time.Hour
	return time.Time(t).Add(dur)
}

//添加分钟
func (t LocalTime) AddMinute(minute int) time.Time {
	var dur time.Duration = time.Duration(minute) * time.Minute
	return time.Time(t).Add(dur)
}

// After reports whether the time instant t is after u.
func (t LocalTime) After(u time.Time) bool {
	return time.Time(t).After(u)
}

// Before reports whether the time instant t is before u.
func (t LocalTime) Before(u time.Time) bool {
	return time.Time(t).Before(u)
}

// Equal reports whether t and u represent the same time instant.
// Two times can be equal even if they are in different locations.
// For example, 6:00 +0200 and 4:00 UTC are Equal.
// See the documentation on the Time type for the pitfalls of using == with
// Time values; most code should use Equal instead.
func (t LocalTime) Equal(u time.Time) bool {
	return time.Time(t).Equal(u)
}
