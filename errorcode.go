package gore

import "encoding/json"

type ErrorCode struct {
	code    int32
	message string
}

func (e *ErrorCode) Code() int32 {
	return e.code
}

func (e *ErrorCode) Message() string {
	return e.message
}

func (e *ErrorCode) SetCode(code int32) {
	e.code = code
}

func (e *ErrorCode) SetMessage(message string) {
	e.message = message
}

func (e *ErrorCode) Error() string {
	b, _ := json.Marshal(map[string]any{
		"code":    e.code,
		"message": e.message,
	})
	return string(b)
}

func NewErrorCode(code int32, message string) *ErrorCode {
	return &ErrorCode{code: code, message: message}
}
