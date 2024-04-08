package errorcode

import "net/http"

var (
	MethodNotAllowed = NewMethodNotAllowed(405000, "the HTTP method associated with the request is not supported")
)

type MethodNotAllowedErr struct {
	Err
}

func NewMethodNotAllowed(code int32, message string) *MethodNotAllowedErr {
	return &MethodNotAllowedErr{Err{Code: code, Message: message, HttpStatus: http.StatusMethodNotAllowed}}
}
