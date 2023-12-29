package errorcode

import "net/http"

var (
	ActivityCodeAlreadyExists = NewConflict(409000, "activity code already exists")
)

type ConflictErr struct {
	Err
}

func NewConflict(code int32, message string) *ForbiddenErr {
	return &ForbiddenErr{Err{Code: code, Message: message, HttpStatus: http.StatusConflict}}
}
