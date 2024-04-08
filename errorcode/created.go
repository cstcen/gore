package errorcode

import "net/http"

var (
	Created = NewConflict(201000, "resource created")
)

type CreatedErr struct {
	Err
}

func NewCreatedErr(code int32, message string) *CreatedErr {
	return &CreatedErr{Err{Code: code, Message: message, HttpStatus: http.StatusConflict}}
}
