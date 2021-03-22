package helper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func JsonDecode(jsonStr string, out interface{}) interface{} {
	json.Unmarshal([]byte(jsonStr), &out)
	return &out
}

func JsonEncode(obj interface{}) string {
	str, _ := json.Marshal(obj)
	return string(str)
}

func JsonToMSS(s string) map[string]string {
	if s == "" {
		return nil
	}
	msi := make(map[string]interface{})
	err := json.Unmarshal([]byte(s), &msi)
	if err != nil {
		return nil
	}
	mss := make(map[string]string)
	for k, v := range msi {
		mss[k] = convertAnyToStr(v)
	}
	return mss
}

// 将任意类型转string
func convertAnyToStr(v interface{}) string {
	if v == nil {
		return ""
	}
	switch d := v.(type) {
	case string:
		return d
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(reflect.ValueOf(v).Int(), 10)
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatUint(reflect.ValueOf(v).Uint(), 10)
	case []byte:
		return string(d)
	case float32, float64:
		return strconv.FormatFloat(reflect.ValueOf(v).Float(), 'f', -1, 64)
	case bool:
		return strconv.FormatBool(d)
	default:
		return fmt.Sprint(v)
	}
}