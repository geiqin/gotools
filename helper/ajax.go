package helper

import (
	"github.com/geiqin/gotools/ajax"
)

//该方法已弃用，请调用 ajax.JsonData 函数
func JsonData(who interface{}, err error) string {
	return ajax.JsonData(who,err)
}

