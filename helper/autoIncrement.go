package helper

import (
	"errors"
	"fmt"
	"github.com/geiqin/gotools/helper/checker"
	"log"
	"strings"
)

//生成自动增长ID
type AutoIncrement struct {
	prefix      int64
	parentId    int64
	isCheck     bool
	padLength   int
	customIndex int64
	suffixList    []string
	shortParentId int64
	depth         int
	format        string
}

func NewAutoIncrement(prefix int64, format ...string) *AutoIncrement {
	formatVal := "2-2-2"
	if format != nil {
		formatVal = format[0]
	}
	return &AutoIncrement{prefix: prefix, format: formatVal}
}

func (b *AutoIncrement) checkFormat(format string) error {
	suffixArr := strings.Split(format, "-")
	b.depth = len(suffixArr)
	b.padLength = len(ToString(b.prefix))
	if b.depth <= 0 || b.depth > 5 {
		return errors.New("编码深度超出范围，最多5层")
	}
	for _, v := range suffixArr {
		if checker.IsInteger(v) {
			num := StringToInt(v)
			if num < 1 || num > 9 {
				return errors.New("format 参数只能是数字+中横线，并且数字必须大于0小于10")
			}
			zeros := ""
			for i := 0; i < num; i++ {
				zeros += "0"
			}
			b.suffixList = append(b.suffixList, zeros)
			b.padLength = b.padLength + num
		} else {
			return errors.New("format 参数只能是数字+中横线，并且数字必须大于0小于10")
		}
	}
	return nil
}

func (b *AutoIncrement) Check(parentId int64, customIndex int64) error {
	b.parentId = parentId
	b.customIndex = customIndex
	if err := b.checkFormat(b.format); err != nil {
		return err
	}

	if b.prefix < 0 {
		return errors.New("编号前缀数字不得小于0，值为0时顶级最多为9条数据")
	}
	var shortParentIdStr string
	if parentId > 0 {
		shortParentIdStr, _ = b.trimRight(ToString(parentId), 0)
		if len(shortParentIdStr) >= b.padLength {
			return errors.New("已超出最大深度限制，请重新选择父级点")
		}
	} else {
		shortParentIdStr = ToString(b.prefix)
	}
	b.shortParentId = StringToInt64(shortParentIdStr)
	b.isCheck = true
	return nil
}

func (b *AutoIncrement) trimRight(idStr string, level int) (string, int) {
	level = level + 1
	last := b.depth - level
	if last >= 0 {
		idStr = strings.TrimSuffix(idStr, b.suffixList[last])
		if last-1 >= 0 {
			if strings.HasSuffix(idStr, b.suffixList[last-1]) {
				return b.trimRight(idStr, level)
			}
		}
	}
	return idStr, last
}
func (b *AutoIncrement) fillLastZero(val int64, suffixStr string) string {
	var str string
	sufLen := len(suffixStr)
	switch sufLen {
	case 1:
		str = fmt.Sprintf("%01d", val)
	case 2:
		str = fmt.Sprintf("%02d", val)
	case 3:
		str = fmt.Sprintf("%03d", val)
	case 4:
		str = fmt.Sprintf("%04d", val)
	case 5:
		str = fmt.Sprintf("%05d", val)
	case 6:
		str = fmt.Sprintf("%06d", val)
	case 7:
		str = fmt.Sprintf("%07d", val)
	case 8:
		str = fmt.Sprintf("%08d", val)
	case 9:
		str = fmt.Sprintf("%09d", val)
	}
	return str
}

func (b *AutoIncrement) MakeNewId(currentMaxId int64) (int64, error) {
	var autoFlag string
	if !b.isCheck {
		return 0, errors.New("检查未通过，不能生成ID,请确认调用check函数是否规范")
	}
	if currentMaxId > 0 {
		log.Println("aaa:", currentMaxId)
		currentMaxIdStr, _ := b.trimRight(ToString(currentMaxId), 0)
		currentMaxId = StringToInt64(currentMaxIdStr)
		if b.prefix == 0 && currentMaxId == 9 {
			return 0, errors.New("当prefix为0时，顶层编号第一位数不得大于9")
		}
		autoFlag = fmt.Sprintf("%d", currentMaxId+1)
	} else {

		if b.prefix == 0 {

			if b.parentId > 0 {
				pidStr, newLevel := b.trimRight(ToString(b.parentId), 0)
				pid := StringToInt64(pidStr)
				if pid > 0 {
					autoFlag = fmt.Sprintf("%d%s", pid, b.fillLastZero(1, b.suffixList[newLevel]))
				}
			} else {
				autoFlag = fmt.Sprintf("%d", 1)
			}

		} else {
			autoFlag = fmt.Sprintf("%d%s", b.shortParentId, b.fillLastZero(1, b.suffixList[0]))
		}
	}

	autoFlag = StrPad(autoFlag, b.padLength, "0") //不足8位末尾补0
	return StringToInt64(autoFlag), nil
}
