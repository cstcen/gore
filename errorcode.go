package gore

import "encoding/json"

var (
	ErrorCodeSuccess                          = NewErrorCode(0, "OK")
	ErrorCodeInvalidArgument                  = NewErrorCode(400, "Invalid argument")
	ErrorCodeFailedPrecondition               = NewErrorCode(400, "Failed precondition")
	ErrorCodeOutOfRange                       = NewErrorCode(400, "Out of range")
	ErrorCodeUnauthenticated                  = NewErrorCode(401, "Unauthenticated")
	ErrorCodePermissionDenied                 = NewErrorCode(403, "Permission denied")
	ErrorCodeNotFound                         = NewErrorCode(404, "Not found")
	ErrorCodeAborted                          = NewErrorCode(409, "Aborted")
	ErrorCodeAlreadyExists                    = NewErrorCode(409, "Already exists")
	ErrorCodeResourceExhausted                = NewErrorCode(429, "Resource exhausted")
	ErrorCodeCancelled                        = NewErrorCode(499, "Cancelled")
	ErrorCodeService                          = NewErrorCode(500, "Service error, please contact us")
	ErrorCodeDataLoss                         = NewErrorCode(500, "Data loss")
	ErrorCodeNetwork                          = NewErrorCode(500, "Network error, please try again later")
	ErrorCodeNotImplemented                   = NewErrorCode(501, "Not implemented")
	ErrorCodeUnavailable                      = NewErrorCode(503, "Unavailable")
	ErrorCodeDeadlineExceeded                 = NewErrorCode(504, "Deadline exceeded")
	ErrorCodeAccountNotFound                  = NewErrorCode(40001, "Account not found")
	ErrorCodeMoneyNotFound                    = NewErrorCode(40002, "Money not found")
	ErrorCodeQuickSlotNotFound                = NewErrorCode(40003, "Quick slot not found")
	ErrorCodeHomelandNotFound                 = NewErrorCode(40004, "Homeland not found")
	ErrorCodeInventoryNotFound                = NewErrorCode(40005, "Inventory not found")
	ErrorCodeEquipNotFound                    = NewErrorCode(40006, "Equip not found")
	ErrorCodeCharacterNotFound                = NewErrorCode(40007, "Character not found")
	ErrorCodeWorldNotFound                    = NewErrorCode(40008, "World not found")
	ErrorCodeItemNotFound                     = NewErrorCode(40009, "Item not found")
	ErrorCodeTransactionNotFound              = NewErrorCode(40010, "Transaction not found")
	ErrorCodeSurnameNotFound                  = NewErrorCode(40011, "Surname not found")
	ErrorCodeLoggedIn                         = NewErrorCode(41001, "Logged in")
	ErrorCodeLoggedOut                        = NewErrorCode(41002, "Logged out")
	ErrorCodeNotConnected                     = NewErrorCode(41003, "Not connected")
	ErrorCodeTransactionCancelled             = NewErrorCode(41004, "Transaction cancelled")
	ErrorCodeTransactionUsed                  = NewErrorCode(41005, "Transaction used")
	ErrorCodeLackOfGold                       = NewErrorCode(41006, "Lack of gold")
	ErrorCodeLackOfStarDiamond                = NewErrorCode(41007, "Lack of star diamond")
	ErrorCodeInvalidRemark                    = NewErrorCode(41008, "Invalid remark")
	ErrorCodeCannotDeleteItem                 = NewErrorCode(50001, "Cannot delete item")
	ErrorCodeCannotModifyItem                 = NewErrorCode(50002, "Cannot modify item")
	ErrorCodePermissionDeniedModifyGold       = NewErrorCode(50003, "Permission denied modify gold")
	ErrorCodePermissionDeniedModifySarDiamond = NewErrorCode(50004, "Permission denied modify sar diamond")
	ErrorCodeTCloudIM                         = NewErrorCode(60001, "Tcloud im")
	ErrorCodeTCloudIMMessage                  = NewErrorCode(60002, "Tcloud immessage")
	ErrorCodeMetaServer                       = NewErrorCode(60101, "MetaServer")
	ErrorCodeMetaServerInvalidAccessLanguage  = NewErrorCode(60102, "Meta server invalid access language")
	ErrorCodeMetaServerDatabaseUpgrading      = NewErrorCode(60103, "Meta server database upgrading")
	ErrorCodeMetaServerDatabaseNotFound       = NewErrorCode(60104, "Meta server database not found")
)

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
