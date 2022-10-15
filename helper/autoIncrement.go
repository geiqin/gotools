package helper

import (
	"errors"
	"fmt" 
	"strings"
)

//生成自动增长ID
type AutoIncrement struct {
	prefix        int64
	parentId      int64
	isCheck       bool
	padLength     int
	customIndex   int64
	suffixStr     string
	shortParentId int64
}

func NewAutoIncrement(padLength int, prefix int64, suffixStrings ...string) *AutoIncrement {
	suffixStr := "00"
	if suffixStrings != nil {
		suffixStr = suffixStrings[0]
	}
	return &AutoIncrement{prefix: prefix, padLength: padLength, suffixStr: suffixStr}
}

func (b *AutoIncrement) Check(parentId int64, customIndex int64) error {
	limitIndex := StringToInt64("1" + b.suffixStr)
	if b.prefix <= 0 {
		return errors.New("编号前缀数字必须大于0，即:prefix值必须大于0")
	}
	if b.customIndex >= limitIndex && customIndex < 0 {
		return errors.New(fmt.Sprintf("自定义序号只能在0到%d之间", limitIndex))
	}
	var shortParentIdStr string
	if parentId > 0 {
		shortParentIdStr = b.trimRight(ToString(parentId))
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

//从右去掉0字符串
func (b *AutoIncrement) trimRight(idStr string) string {
	idStr = strings.TrimSuffix(idStr, b.suffixStr)
	if strings.HasSuffix(idStr, b.suffixStr) {
		return b.trimRight(idStr)
	}
	return idStr
}
//生成新的ID
func (b *AutoIncrement) MakeNewId(currentMaxId int64) (int64, error) {
	var autoFlag string
	if !b.isCheck {
		return 0, errors.New("检查未通过，不能生成ID,请确认调用check函数是否规范")
	}
	if currentMaxId > 0 {
		currentMaxIdStr := b.trimRight(ToString(currentMaxId))
		currentMaxId = StringToInt64(currentMaxIdStr)
		autoFlag = fmt.Sprintf("%d", currentMaxId+1)
	} else {
		autoFlag = fmt.Sprintf("%d%02d", b.shortParentId, 1)
	}

	autoFlag = StrPad(autoFlag, b.padLength, "0") //不足8位末尾补0
	return StringToInt64(autoFlag), nil
}
