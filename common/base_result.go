package common

import (
	"encoding/json"
)

var (
	BaseResultSuccess                          = NewBaseResult(0, "OK")
	BaseResultInvalidArgument                  = NewBaseResult(400, "Invalid argument")
	BaseResultFailedPrecondition               = NewBaseResult(400, "Failed precondition")
	BaseResultOutOfRange                       = NewBaseResult(400, "Out of range")
	BaseResultUnauthenticated                  = NewBaseResult(401, "Unauthenticated")
	BaseResultPermissionDenied                 = NewBaseResult(403, "Permission denied")
	BaseResultNotFound                         = NewBaseResult(404, "Not found")
	BaseResultAborted                          = NewBaseResult(409, "Aborted")
	BaseResultAlreadyExists                    = NewBaseResult(409, "Already exists")
	BaseResultResourceExhausted                = NewBaseResult(429, "Resource exhausted")
	BaseResultCancelled                        = NewBaseResult(499, "Cancelled")
	BaseResultService                          = NewBaseResult(500, "Service error, please contact us")
	BaseResultDataLoss                         = NewBaseResult(500, "Data loss")
	BaseResultNetwork                          = NewBaseResult(500, "Network error, please try again later")
	BaseResultNotImplemented                   = NewBaseResult(501, "Not implemented")
	BaseResultUnavailable                      = NewBaseResult(503, "Unavailable")
	BaseResultDeadlineExceeded                 = NewBaseResult(504, "Deadline exceeded")
	BaseResultAccountNotFound                  = NewBaseResult(40001, "Account not found")
	BaseResultMoneyNotFound                    = NewBaseResult(40002, "Money not found")
	BaseResultQuickSlotNotFound                = NewBaseResult(40003, "Quick slot not found")
	BaseResultHomelandNotFound                 = NewBaseResult(40004, "Homeland not found")
	BaseResultInventoryNotFound                = NewBaseResult(40005, "Inventory not found")
	BaseResultEquipNotFound                    = NewBaseResult(40006, "Equip not found")
	BaseResultCharacterNotFound                = NewBaseResult(40007, "Character not found")
	BaseResultWorldNotFound                    = NewBaseResult(40008, "World not found")
	BaseResultItemNotFound                     = NewBaseResult(40009, "Item not found")
	BaseResultTransactionNotFound              = NewBaseResult(40010, "Transaction not found")
	BaseResultSurnameNotFound                  = NewBaseResult(40011, "Surname not found")
	BaseResultLoggedIn                         = NewBaseResult(41001, "Logged in")
	BaseResultLoggedOut                        = NewBaseResult(41002, "Logged out")
	BaseResultNotConnected                     = NewBaseResult(41003, "Not connected")
	BaseResultTransactionCancelled             = NewBaseResult(41004, "Transaction cancelled")
	BaseResultTransactionUsed                  = NewBaseResult(41005, "Transaction used")
	BaseResultLackOfGold                       = NewBaseResult(41006, "Lack of gold")
	BaseResultLackOfStarDiamond                = NewBaseResult(41007, "Lack of star diamond")
	BaseResultInvalidRemark                    = NewBaseResult(41008, "Invalid remark")
	BaseResultCannotDeleteItem                 = NewBaseResult(50001, "Cannot delete item")
	BaseResultCannotModifyItem                 = NewBaseResult(50002, "Cannot modify item")
	BaseResultPermissionDeniedModifyGold       = NewBaseResult(50003, "Permission denied modify gold")
	BaseResultPermissionDeniedModifySarDiamond = NewBaseResult(50004, "Permission denied modify sar diamond")
	BaseResultTCloudIM                         = NewBaseResult(60001, "Tcloud im")
	BaseResultTCloudIMMessage                  = NewBaseResult(60002, "Tcloud immessage")
	BaseResultMetaServer                       = NewBaseResult(60101, "MetaServer")
	BaseResultMetaServerInvalidAccessLanguage  = NewBaseResult(60102, "Meta server invalid access language")
	BaseResultMetaServerDatabaseUpgrading      = NewBaseResult(60103, "Meta server database upgrading")
	BaseResultMetaServerDatabaseNotFound       = NewBaseResult(60104, "Meta server database not found")
)

type Error interface {
	error
	Code() int32
	Message() string
}

type BaseResult struct {
	code    int32
	message string
}

func NewBaseResult(code int32, message string) *BaseResult {
	return &BaseResult{code: code, message: message}
}

func (b *BaseResult) Code() int32 {
	return b.code
}

func (b *BaseResult) Message() string {
	return b.message
}

func (b *BaseResult) SetCode(code int32) Error {
	tmp := *b
	tmp.code = code
	return &tmp
}

func (b *BaseResult) SetMsg(msg string) Error {
	tmp := *b
	tmp.message = msg
	return &tmp
}

func (b *BaseResult) Error() string {
	raw, _ := b.MarshalJSON()
	return string(raw)
}

// MarshalJSON implements the JSON encoding interface
func (b *BaseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"code":    b.code,
		"message": b.message,
	})
}

func (b *BaseResult) UnmarshalJSON(data []byte) error {
	var result struct {
		Code    int32
		Message string
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}
	_ = b.SetCode(result.Code)
	_ = b.SetMsg(result.Message)
	return nil
}
