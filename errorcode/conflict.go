package errorcode

import "net/http"

var (
	Conflict                          = NewConflict(409000, "resource conflict")
	ConflictActivityCodeAlreadyExists = NewConflict(409001, "activity code already exists")
)

type ConflictErr struct {
	Err
}

func NewConflict(code int32, message string) *ConflictErr {
	return &ConflictErr{Err{Code: code, Message: message, HttpStatus: http.StatusConflict}}
}
