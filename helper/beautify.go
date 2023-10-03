package helper

import "github.com/geiqin/gotools/helper/xtime"

//美化时间
func BeautifyTime(strTime string, format string) string {
	if strTime == "" {
		return ""
	}
	t, err := xtime.TimeStr2Time(strTime)
	if err != nil {
		return ""
	}
	return xtime.DateFormat(t, format)
}
