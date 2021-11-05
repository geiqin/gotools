package ajax

import (
	"github.com/geiqin/gotools/helper"
	"reflect"
)

//主要针对Service层结果转换为Json数据
func JsonData(who interface{}, err error) string {
	ret := &ResultData{}
	pageData := &PageData{}

	if err != nil {
		ret.Code = 400
		ret.Message = err.Error()
	}

	obj := reflect.ValueOf(who)
	elem := obj.Elem()
	if elem.Kind() == reflect.Struct {

		errData := elem.FieldByName("Error")
		if errData.IsValid() {
			if !errData.IsNil() {
				errElem := errData.Elem()
				code := errElem.FieldByName("Code")
				msg := errElem.FieldByName("Message")
				ret.Code = code.Int()
				ret.Message = msg.String()
				return helper.JsonEncode(ret)
			}
		}

		data := elem.FieldByName("Data")
		if data.IsValid() {
			if !data.IsNil() {
				ret.Code = 1
				ret.Data = data.Interface()
				return helper.JsonEncode(ret)
			}
		}

		info := elem.FieldByName("Info")
		if info.IsValid() {
			if !info.IsNil() {
				ret.Code = 1
				ret.Data = info.Interface()
				return helper.JsonEncode(ret)
			}
		}

		entity := elem.FieldByName("Entity")
		if entity.IsValid() {
			if !entity.IsNil() {
				ret.Code = 1
				ret.Data = entity.Interface()
				return helper.JsonEncode(ret)
			}
		}

		params := elem.FieldByName("Params")
		if params.IsValid() {
			if !params.IsNil() {
				ret.Code = 1
				ret.Data = params.Interface()
				return helper.JsonEncode(ret)
			}
		}

		//分页数据和列表数据
		pager := elem.FieldByName("Pager")
		if pager.IsValid() {
			if !pager.IsNil() {
				ret.Code = 1
				pageData.Pager = pager.Interface()
			}
		}
		items := elem.FieldByName("Items")
		if items.IsValid() {
			if !items.IsNil() {
				ret.Code = 1
				pageData.Items = items.Interface()
			}
		}
		if pageData.Pager != nil {
			ret.Code = 1
			ret.Data = pageData
			return helper.JsonEncode(ret)
		}
		if pageData.Pager == nil && pageData.Items != nil {
			ret.Code = 1
			ret.Data = pageData.Items
			return helper.JsonEncode(ret)
		}
	}
	if ret.Code == 0 {
		ret.Code = 1
	}
	return helper.JsonEncode(ret)
}
