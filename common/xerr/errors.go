package xerr

import "fmt"

type CodeError struct {
	errCode uint32
	errMsg  string
}
//返回给前端的错位码
func (e *CodeError)GetErrCode() uint32 {
	return e.errCode
}
//返回给前端的错误信息
func (e *CodeError)GetErrMsg() string {
	return e.errMsg
}
//实现error接口
func (e *CodeError)Error() string {
	return fmt.Sprintf("error code: %d, error message: %s", e.errCode, e.errMsg)
}
//
func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{
		errCode: errCode,
		errMsg:  errMsg,
	}
}
//
func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{
		errCode: errCode,
		errMsg:  MapErrMsg(errCode),
	}
}
//
func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{
		errCode:  SERVER_COMMON_ERROR,
			errMsg:  errMsg,
	}
}