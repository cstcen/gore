package errorcode

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	OK = New(0, "OK")
)

type Error interface {
	error
	GetHttpStatus() int
	GetCode() int32
	GetMessage() string
	WithCode(code int32) Error
	WithMsg(msg string) Error
	WithHttpStatus(httpStatus int) Error
}

type Err struct {
	Code       int32  `json:"code"`
	Message    string `json:"message,omitempty"`
	HttpStatus int    `json:"-"`
}

func New(code int32, message string, httpStatus ...int) *Err {
	status := http.StatusOK
	if len(httpStatus) > 0 {
		status = httpStatus[0]
	}
	return &Err{Code: code, Message: message, HttpStatus: status}
}

func (e *Err) GetCode() int32 {
	return e.Code
}

func (e *Err) GetMessage() string {
	return e.Message
}

func (e *Err) GetHttpStatus() int {
	if e.HttpStatus == 0 {
		return http.StatusOK
	}
	return e.HttpStatus
}

func (e *Err) WithCode(code int32) Error {
	tmp := *e
	tmp.Code = code
	return &tmp
}

func (e *Err) WithMsg(msg string) Error {
	tmp := *e
	tmp.Message = fmt.Sprintf("%s: %s", tmp.Message, msg)
	return &tmp
}

func (e *Err) WithHttpStatus(httpStatus int) Error {
	tmp := *e
	tmp.HttpStatus = httpStatus
	return &tmp
}

func (e *Err) Error() string {
	raw, _ := json.Marshal(e)
	return string(raw)
}
