package helper

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/shomali11/util/xhashes"
	"log"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//结构转换，包括数组结构
func ConvertData(from interface{}, to interface{}) interface{} {
	json := JsonEncode(from)
	if json != "" {
		JsonDecode(json, &to)
	}
	return to
}

func GetVal(name string, mps map[string]interface{}) interface{} {
	v, ok := mps[name]
	if !ok {
		return nil
	}
	return v
}

//生成id含前缀和后缀字符串
func GetIdentityFlag(id int64, prefix string, suffix string) string {
	flag := fmt.Sprintf("%08d", id)
	flag = prefix + flag + suffix
	return flag
}

func GenerateSn(prefix ...string) string {
	sn := xhashes.FNV64(UniqueId())
	snStr := strconv.FormatUint(sn, 10)
	if prefix != nil {
		snStr = prefix[0] + snStr
	}
	return snStr
}

//获取随机数字串（验证码常用）
func GetRandomNumber(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// 获取随机字符串
//    length：字符串长度
func GetRandomString(length int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	var (
		result []byte
		b      []byte
		r      *rand.Rand
	)
	b = []byte(str)
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

/*
//把汉字转换未Pinyin
func ConvertPinyin(chinese string, mode ...int) string {
	strArr := pinyin.LazyConvert(chinese, nil)
	if strArr != nil {
		return strings.Join(strArr, "")
	}
	return ""
}


*/
//截取字符串 start 起点下标 end 终点下标(不包括)
func Substr(str string, start int, end int) (string, error) {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return "", errors.New("start is wrong:" + string(start))
	}

	if end < 0 || end > length {
		return "", errors.New("end is wrong:" + string(start))
	}
	return string(rs[start:end]), nil
}

//Map类型转换为Struct
func MapToStruct(fromMap interface{}, toStruct interface{}) interface{} {
	mapstructure.Decode(fromMap, toStruct)
	return toStruct
}

//判断字符是否在数组中
func InArray(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//string到int64
func StringToInt(val string) int {
	// string到int
	ret, _ := strconv.Atoi(val)
	return ret
}

func StringToInt32(val string) int32 {
	ret, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		log.Println("StringToInt32 convert failed ,value is :", val)
		return 0
	}
	return int32(ret)
}

func StringToInt64(val string) int64 {
	ret, _ := strconv.ParseInt(val, 10, 64)
	return ret
}
func IntToString(val int) string {
	// int到string
	ret := strconv.Itoa(val)
	return ret
}
func Int64ToString(val int64) string {
	// int64到string
	ret := strconv.FormatInt(val, 10)
	return ret
}

func Int32ToInt(value int32) int {
	str := Int64ToString(int64(value))
	return StringToInt(str)
}

func IntToInt32(value int) int32 {
	str := IntToString(value)
	return StringToInt32(str)
}

// 将任意类型转string
func ToString(v interface{}) string {
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

//把任意数字类型转换为int64
func ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}

func DisplaySize(raw float64) string {
	if raw < 1024 {
		return fmt.Sprintf("%.2fB", raw)
	}

	if raw < 1024*1024 {
		return fmt.Sprintf("%.2fK", raw/1024.0)
	}

	if raw < 1024*1024*1024 {
		return fmt.Sprintf("%.2fM", raw/1024.0/1024.0)
	}

	if raw < 1024*1024*1024*1024 {
		return fmt.Sprintf("%.2fG", raw/1024.0/1024.0/1024.0)
	}

	if raw < 1024*1024*1024*1024*1024 {
		return fmt.Sprintf("%.2fT", raw/1024.0/1024.0/1024.0/1024.0)
	}

	if raw < 1024*1024*1024*1024*1024*1024 {
		return fmt.Sprintf("%.2fP", raw/1024.0/1024.0/1024.0/1024.0/1024.0)
	}

	return "TooLarge"
}

// addslashes() 函数返回在预定义字符之前添加反斜杠的字符串。
// 预定义字符是：
// 单引号（'）
// 双引号（"）
// 反斜杠（\）
func AddSlashes(str string) string {
	tmpRune := []rune{}
	strRune := []rune(str)
	for _, ch := range strRune {
		switch ch {
		case []rune{'\\'}[0], []rune{'"'}[0], []rune{'\''}[0]:
			tmpRune = append(tmpRune, []rune{'\\'}[0])
			tmpRune = append(tmpRune, ch)
		default:
			tmpRune = append(tmpRune, ch)
		}
	}
	return string(tmpRune)
}

// stripslashes() 函数删除由 addslashes() 函数添加的反斜杠。
func StripSlashes(str string) string {
	dstRune := []rune{}
	strRune := []rune(str)
	strLenth := len(strRune)
	for i := 0; i < strLenth; i++ {
		if strRune[i] == []rune{'\\'}[0] {
			i++
		}
		dstRune = append(dstRune, strRune[i])
	}
	return string(dstRune)
}

// 数组是否包含某元素（string）
func HasContainString(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// 数组是否包含某元素（int32）
func HasContainInt32(items []int32, item int32) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// 数组是否包含某元素（int64）
func HasContainInt64(items []int64, item int64) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

//判断是否为URL
func HasURL(str string) bool {
	str = strings.ToLower(str)
	if strings.HasPrefix(str, "http://") || strings.HasPrefix(str, "https://") {
		return true
	}
	return false
}

/**
 * @Author Dong
 * @Description 获得当前月的初始和结束日期
 * @Date 16:29 2020/8/6
 * @Param  * @param null
 * @return
 **/
func GetMonthDay() (string, string) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	f := firstOfMonth.Unix()
	l := lastOfMonth.Unix()
	return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

/**
 * @Author Dong
 * @Description 获得当前周的初始和结束日期
 * @Date 16:32 2020/8/6
 * @Param  * @param null
 * @return
 **/
func GetWeekDay() (string, string) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if offset > 0 {
		offset = -6
	}

	lastoffset := int(time.Saturday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if lastoffset == 6 {
		lastoffset = -1
	}

	firstOfWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	lastOfWeeK := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, lastoffset+1)
	f := firstOfWeek.Unix()
	l := lastOfWeeK.Unix()
	return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

/**
 * @Author Dong
 * @Description //获得当前季度的初始和结束日期
 * @Date 16:33 2020/8/6
 * @Param  * @param null
 * @return
 **/
func GetQuarterDay() (string, string) {
	year := time.Now().Format("2006")
	month := int(time.Now().Month())
	var firstOfQuarter string
	var lastOfQuarter string
	if month >= 1 && month <= 3 {
		//1月1号
		firstOfQuarter = year + "-01-01 00:00:00"
		lastOfQuarter = year + "-03-31 23:59:59"
	} else if month >= 4 && month <= 6 {
		firstOfQuarter = year + "-04-01 00:00:00"
		lastOfQuarter = year + "-06-30 23:59:59"
	} else if month >= 7 && month <= 9 {
		firstOfQuarter = year + "-07-01 00:00:00"
		lastOfQuarter = year + "-09-30 23:59:59"
	} else {
		firstOfQuarter = year + "-10-01 00:00:00"
		lastOfQuarter = year + "-12-31 23:59:59"
	}
	return firstOfQuarter, lastOfQuarter
}
