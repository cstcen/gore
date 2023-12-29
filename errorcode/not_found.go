package errorcode

import "net/http"

var (
	NotFound                = NewNotFound(404000, "the requested operation failed because a resource associated with the request could not be found")
	NotFoundProtocol        = NewNotFound(404001, "the protocol used in the request is not supported")
	NotFoundMember          = NewNotFound(404002, "member was not found")
	NotFoundCharacter       = NewNotFound(404003, "character was not found")
	NotFoundUgc             = NewNotFound(404004, "ugc was not found")
	NotFoundUgcTeam         = NewNotFound(404005, "ugc team was not found")
	NotFoundProduct         = NewNotFound(404006, "product was not found")
	NotFoundProductSnapshot = NewNotFound(404007, "product snapshot was not found")
	NotFoundOrder           = NewNotFound(404008, "order was not found")
	NotFoundCashType        = NewNotFound(404009, "cash type was not found")

	NotFoundActivity = NewNotFound(404401, "activity not found")
	NotFoundMission  = NewNotFound(404402, "mission not found")
)

type NotFoundErr struct {
	Err
}

func NewNotFound(code int32, message string) *NotFoundErr {
	return &NotFoundErr{Err{Code: code, Message: message, HttpStatus: http.StatusNotFound}}
}
