package helper

import (
	"fmt"
	"github.com/geiqin/gotools/helper/xtime"
)

//美化时间(如: YYYY-MM-DD hh:mm)
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

//美化图片路径(常用的size有200,800,1200 等)【针对七牛云】
func BeautyImageUrl(url string, size int, quality ...int) string {
	var qualityVal = 75
	if quality != nil {
		qualityVal = quality[0]
	}
	if HasURL(url) {
		url = url + fmt.Sprintf("?imageMogr2/auto-orient/thumbnail/%dx%d>/blur/1x0/quality/%d", size, size, qualityVal)
	}
	return url
}
