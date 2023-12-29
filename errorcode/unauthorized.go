package errorcode

import "net/http"

var (
	Unauthorized                    = NewUnauthorized(401000, "the user is not authorized make the request")
	UnauthorizedInvalidAccessToken  = NewUnauthorized(401001, "invalid access_token")
	UnauthorizedInvalidRefreshToken = NewUnauthorized(401002, "invalid refresh_token")
)

type UnauthorizedErr struct {
	Err
}

func NewUnauthorized(code int32, message string) *UnauthorizedErr {
	return &UnauthorizedErr{Err{Code: code, Message: message, HttpStatus: http.StatusUnauthorized}}
}
