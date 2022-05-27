package gore

import "encoding/json"

type ErrorCode struct {
	code    int32
	message string
}

func NewErrorCode(code int32, message string) *ErrorCode {
	return &ErrorCode{code: code, message: message}
}

func (e *ErrorCode) Error() string {
	b, _ := json.Marshal(map[string]any{
		"code":    e.code,
		"message": e.message,
	})
	return string(b)
}

func (e *ErrorCode) Code() int32 {
	return e.code
}
