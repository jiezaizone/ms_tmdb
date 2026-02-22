package xerr

import "fmt"

// 业务错误码
const (
	OK               = 0
	ServerError      = 10001
	ParamError       = 10002
	NotFound         = 10003
	TmdbRequestError = 10004
	DatabaseError    = 10005
	SyncError        = 10006
)

// CodeError 业务错误
type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) *CodeError {
	return &CodeError{Code: code, Msg: msg}
}

func NewParamError(msg string) *CodeError {
	return &CodeError{Code: ParamError, Msg: msg}
}

func NewNotFoundError(msg string) *CodeError {
	return &CodeError{Code: NotFound, Msg: msg}
}

func NewTmdbError(msg string) *CodeError {
	return &CodeError{Code: TmdbRequestError, Msg: msg}
}

func NewDBError(msg string) *CodeError {
	return &CodeError{Code: DatabaseError, Msg: msg}
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("错误码: %d, 信息: %s", e.Code, e.Msg)
}

func (e *CodeError) Data() *CodeError {
	return e
}
