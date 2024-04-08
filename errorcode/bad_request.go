package errorcode

import "net/http"

var (
	BadRequest                         = NewBadRequest(400000, "the API request is invalid or improperly formed")
	BadRequestInvalidHeader            = NewBadRequest(400001, "invalid header.")
	BadRequestInvalidParameter         = NewBadRequest(400002, "invalid parameter")
	BadRequestInvalidQuery             = NewBadRequest(400003, "invalid query")
	BadRequestInvalidBody              = NewBadRequest(400004, "invalid request body")
	BadRequestInvalidClientId          = NewBadRequest(400005, "invalid client_id")
	BadRequestInvalidRecipients        = NewBadRequest(400006, "invalid recipients")
	BadRequestInvalidReservationTime   = NewBadRequest(400007, "invalid reservation_time")
	BadRequestInvalidMemberNoList      = NewBadRequest(400008, "member_no list must not be empty")
	BadRequestInvalidChannel           = NewBadRequest(400009, "invalid channel")
	BadRequestDuplicateProductInfo     = NewBadRequest(400010, "duplicate product information appears in the display page")
	BadRequestNotEnoughPoint           = NewBadRequest(400011, "not enough point")
	BadRequestAuthenticationMode       = NewBadRequest(400012, "invalid authentication mode")
	BadRequestInvalidPassword          = NewBadRequest(400013, "invalid password")
	BadRequestInvalidRedirectUri       = NewBadRequest(400014, "invalid redirect uri")
	BadRequestInvalidResponseType      = NewBadRequest(400015, "invalid response type")
	BadRequestInvalidVerificationCode  = NewBadRequest(400016, "invalid verification code")
	BadRequestInvalidGrantType         = NewBadRequest(400017, "invalid grant type")
	BadRequestInvalidAuthorizationCode = NewBadRequest(400018, "invalid authorization code")
	BadRequestNicknameExists           = NewBadRequest(400019, "nickname exists")
	BadRequestNicknameBanned           = NewBadRequest(400020, "nickname banned")
)

type BadRequestErr struct {
	Err
}

func NewBadRequest(code int32, message string) *BadRequestErr {
	return &BadRequestErr{Err{Code: code, Message: message, HttpStatus: http.StatusBadRequest}}
}
