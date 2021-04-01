package ajax

import "github.com/geiqin/gotools/helper"

//结果数据
type ResultData struct {
	Code    int64       `json:"code"`              //错误代码: 成功：1 ，其它数字为失败
	Message string      `json:"message,omitempty"` //错误消息
	Data    interface{} `json:"data"`              //成功数据
}

//分页数据
type PageData struct {
	Pager interface{} `json:"pager"`
	Items interface{} `json:"items"`
}

//输出错误json数据
func Failed(message string, errCode ...int64) string {
	var code int64 = 400
	if errCode != nil {
		c := errCode[0]
		if c > 1 {
			code = c
		}
	}
	ret := &ResultData{
		Code:    code,
		Message: message,
	}
	return helper.JsonEncode(ret)
}

//输出成功json数据
func Succeed(data interface{}) string {
	ret := &ResultData{
		Code: 1,
		Data: data,
	}
	return helper.JsonEncode(ret)
}
