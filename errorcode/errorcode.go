package errorcode

import (
	"encoding/json"
)

var (
	Success                          = NewErrorCode(0, "OK")
	InvalidArgument                  = NewErrorCode(400, "Invalid argument")
	FailedPrecondition               = NewErrorCode(400, "Failed precondition")
	OutOfRange                       = NewErrorCode(400, "Out of range")
	Unauthenticated                  = NewErrorCode(401, "Unauthenticated")
	PermissionDenied                 = NewErrorCode(403, "Permission denied")
	NotFound                         = NewErrorCode(404, "Not found")
	Aborted                          = NewErrorCode(409, "Aborted")
	AlreadyExists                    = NewErrorCode(409, "Already exists")
	ResourceExhausted                = NewErrorCode(429, "Resource exhausted")
	Cancelled                        = NewErrorCode(499, "Cancelled")
	Service                          = NewErrorCode(500, "Service error, please contact us")
	DataLoss                         = NewErrorCode(500, "Data loss")
	Network                          = NewErrorCode(500, "Network error, please try again later")
	NotImplemented                   = NewErrorCode(501, "Not implemented")
	Unavailable                      = NewErrorCode(503, "Unavailable")
	DeadlineExceeded                 = NewErrorCode(504, "Deadline exceeded")
	AccountNotFound                  = NewErrorCode(40001, "Account not found")
	MoneyNotFound                    = NewErrorCode(40002, "Money not found")
	QuickSlotNotFound                = NewErrorCode(40003, "Quick slot not found")
	HomelandNotFound                 = NewErrorCode(40004, "Homeland not found")
	InventoryNotFound                = NewErrorCode(40005, "Inventory not found")
	EquipNotFound                    = NewErrorCode(40006, "Equip not found")
	CharacterNotFound                = NewErrorCode(40007, "Character not found")
	WorldNotFound                    = NewErrorCode(40008, "World not found")
	ItemNotFound                     = NewErrorCode(40009, "Item not found")
	TransactionNotFound              = NewErrorCode(40010, "Transaction not found")
	SurnameNotFound                  = NewErrorCode(40011, "Surname not found")
	LoggedIn                         = NewErrorCode(41001, "Logged in")
	LoggedOut                        = NewErrorCode(41002, "Logged out")
	NotConnected                     = NewErrorCode(41003, "Not connected")
	TransactionCancelled             = NewErrorCode(41004, "Transaction cancelled")
	TransactionUsed                  = NewErrorCode(41005, "Transaction used")
	LackOfGold                       = NewErrorCode(41006, "Lack of gold")
	LackOfStarDiamond                = NewErrorCode(41007, "Lack of star diamond")
	InvalidRemark                    = NewErrorCode(41008, "Invalid remark")
	CannotDeleteItem                 = NewErrorCode(50001, "Cannot delete item")
	CannotModifyItem                 = NewErrorCode(50002, "Cannot modify item")
	PermissionDeniedModifyGold       = NewErrorCode(50003, "Permission denied modify gold")
	PermissionDeniedModifySarDiamond = NewErrorCode(50004, "Permission denied modify sar diamond")
	TCloudIM                         = NewErrorCode(60001, "Tcloud im")
	TCloudIMMessage                  = NewErrorCode(60002, "Tcloud immessage")
	MetaServer                       = NewErrorCode(60101, "MetaServer")
	MetaServerInvalidAccessLanguage  = NewErrorCode(60102, "Meta server invalid access language")
	MetaServerDatabaseUpgrading      = NewErrorCode(60103, "Meta server database upgrading")
	MetaServerDatabaseNotFound       = NewErrorCode(60104, "Meta server database not found")
)

type ErrorCode interface {
	error
	Code() int32
}

type errorCode struct {
	code    int32
	message string
}

func (e *errorCode) Code() int32 {
	return e.code
}

func (e *errorCode) Message() string {
	return e.message
}

func (e *errorCode) SetCode(code int32) *errorCode {
	e.code = code
	return e
}

func (e *errorCode) SetMessage(message string) *errorCode {
	e.message = message
	return e
}

func (e *errorCode) Error() string {
	result, _ := e.MarshalJSON()
	return string(result)
}

// MarshalJSON implements the JSON encoding interface
func (e *errorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"code":    e.code,
		"message": e.message,
	})
}

func New(code int32, message string) *errorCode {
	return &errorCode{code: code, message: message}
}

// NewErrorCode
// Deprecated
func NewErrorCode(code int32, message string) *errorCode {
	return New(code, message)
}
