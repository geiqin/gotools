package errors

import "fmt"

// 通用错误类型
type CommonError struct {
	ErrCode int32
	Message string
}

// 实现error接口的Error方法
func (e *CommonError) Error() string {
	return fmt.Sprintf("Error Code: %d, Message: %s", e.ErrCode, e.Message)
}

func NewCommonError(errMsg string, errCode ...int32) error {
	var errorCode int32 = 100 //普通错误
	if errCode != nil {
		errorCode = errCode[0]
	}
	return &CommonError{
		ErrCode: errorCode,
		Message: errMsg,
	}
}
