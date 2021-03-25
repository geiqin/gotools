package helper


import (
	"reflect"
)

type ajaxData struct {
	Code    int64       `json:"code"`              //错误代码: 成功：1 ，其它数字为失败
	Message string      `json:"message,omitempty"` //错误消息
	Data    interface{} `json:"data,omitempty"`    //成功数据
}

type pageData struct {
	Pager interface{} `json:"pager,omitempty"`
	Items interface{} `json:"items,omitempty"`
}

func JsonData(who interface{}, err error) string {
	ajax := &ajaxData{}
	pageData := &pageData{}

	if err != nil {
		ajax.Code = 1000
		ajax.Message = err.Error()
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
				ajax.Code = code.Int()
				ajax.Message = msg.String()
				return JsonEncode(ajax)
			}
		}

		info := elem.FieldByName("Info")
		if info.IsValid() {
			if !info.IsNil() {
				ajax.Code = 1
				ajax.Data = info.Interface()
				return JsonEncode(ajax)
			}
		}

		entity := elem.FieldByName("Entity")
		if entity.IsValid() {
			if !entity.IsNil() {
				ajax.Code = 1
				ajax.Data = entity.Interface()
				return JsonEncode(ajax)
			}
		}

		//分页数据和列表数据
		pager := elem.FieldByName("Pager")
		if pager.IsValid() {
			if !pager.IsNil() {
				ajax.Code = 1
				pageData.Pager = pager.Interface()
			}
		}
		items := elem.FieldByName("Items")
		if items.IsValid() {
			if !items.IsNil() {
				ajax.Code = 1
				pageData.Items = items.Interface()
			}
		}
		if pageData.Pager != nil {
			ajax.Code = 1
			ajax.Data = pageData
			return JsonEncode(ajax)
		}
		if pageData.Pager == nil && pageData.Items != nil {
			ajax.Code = 1
			ajax.Data = pageData.Items
			return JsonEncode(ajax)
		}
	}
	return JsonEncode(ajax)
}

