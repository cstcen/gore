package common

import "net/http"

var (
	BaseResultSuccess                          = BaseResult{Err{HttpStatus: http.StatusOK, Code: 0, Message: "OK"}}
	BaseResultInvalidArgument                  = BaseResult{Err{HttpStatus: http.StatusOK, Code: 400, Message: "Invalid argument"}}
	BaseResultFailedPrecondition               = BaseResult{Err{HttpStatus: http.StatusOK, Code: 400, Message: "Failed precondition"}}
	BaseResultOutOfRange                       = BaseResult{Err{HttpStatus: http.StatusOK, Code: 400, Message: "Out of range"}}
	BaseResultUnauthenticated                  = BaseResult{Err{HttpStatus: http.StatusOK, Code: 401, Message: "Unauthenticated"}}
	BaseResultPermissionDenied                 = BaseResult{Err{HttpStatus: http.StatusOK, Code: 403, Message: "Permission denied"}}
	BaseResultNotFound                         = BaseResult{Err{HttpStatus: http.StatusOK, Code: 404, Message: "Not found"}}
	BaseResultAborted                          = BaseResult{Err{HttpStatus: http.StatusOK, Code: 409, Message: "Aborted"}}
	BaseResultAlreadyExists                    = BaseResult{Err{HttpStatus: http.StatusOK, Code: 409, Message: "Already exists"}}
	BaseResultResourceExhausted                = BaseResult{Err{HttpStatus: http.StatusOK, Code: 429, Message: "Resource exhausted"}}
	BaseResultCancelled                        = BaseResult{Err{HttpStatus: http.StatusOK, Code: 499, Message: "Cancelled"}}
	BaseResultService                          = BaseResult{Err{HttpStatus: http.StatusOK, Code: 500, Message: "Service error, please contact us"}}
	BaseResultDataLoss                         = BaseResult{Err{HttpStatus: http.StatusOK, Code: 500, Message: "Data loss"}}
	BaseResultNetwork                          = BaseResult{Err{HttpStatus: http.StatusOK, Code: 500, Message: "Network error, please try again later"}}
	BaseResultNotImplemented                   = BaseResult{Err{HttpStatus: http.StatusOK, Code: 501, Message: "Not implemented"}}
	BaseResultUnavailable                      = BaseResult{Err{HttpStatus: http.StatusOK, Code: 503, Message: "Unavailable"}}
	BaseResultDeadlineExceeded                 = BaseResult{Err{HttpStatus: http.StatusOK, Code: 504, Message: "Deadline exceeded"}}
	BaseResultAccountNotFound                  = BaseResult{Err{HttpStatus: http.StatusOK, Code: 40001, Message: "Account not found"}}
	BaseResultMoneyNotFound                    = BaseResult{Err{HttpStatus: http.StatusOK, Code: 40002, Message: "Money not found"}}
	BaseResultQuickSlotNotFound                = BaseResult{Err{HttpStatus: http.StatusOK, Code: 40003, Message: "Quick slot not found"}}
	BaseResultHomelandNotFound                 = BaseResult{Err{HttpStatus: http.StatusOK, Code: 40004, Message: "Homeland not found"}}
	BaseResultInventoryNotFound                = BaseResult{Err{HttpStatus: http.StatusOK, Code: 40005, Message: "Inventory not found"}}
	BaseResultEquipNotFound                    = BaseResult{Err{HttpStatus: http.StatusOK, Code: 40006, Message: "Equip not found"}}
	BaseResultCharacterNotFound                = BaseResult{Err{HttpStatus: http.StatusOK, Code: 40007, Message: "Character not found"}}
	BaseResultWorldNotFound                    = BaseResult{Err{HttpStatus: http.StatusOK, Code: 40008, Message: "World not found"}}
	BaseResultItemNotFound                     = BaseResult{Err{HttpStatus: http.StatusOK, Code: 40009, Message: "Item not found"}}
	BaseResultTransactionNotFound              = BaseResult{Err{HttpStatus: http.StatusOK, Code: 40010, Message: "Transaction not found"}}
	BaseResultSurnameNotFound                  = BaseResult{Err{HttpStatus: http.StatusOK, Code: 40011, Message: "Surname not found"}}
	BaseResultLoggedIn                         = BaseResult{Err{HttpStatus: http.StatusOK, Code: 41001, Message: "Logged in"}}
	BaseResultLoggedOut                        = BaseResult{Err{HttpStatus: http.StatusOK, Code: 41002, Message: "Logged out"}}
	BaseResultNotConnected                     = BaseResult{Err{HttpStatus: http.StatusOK, Code: 41003, Message: "Not connected"}}
	BaseResultTransactionCancelled             = BaseResult{Err{HttpStatus: http.StatusOK, Code: 41004, Message: "Transaction cancelled"}}
	BaseResultTransactionUsed                  = BaseResult{Err{HttpStatus: http.StatusOK, Code: 41005, Message: "Transaction used"}}
	BaseResultLackOfGold                       = BaseResult{Err{HttpStatus: http.StatusOK, Code: 41006, Message: "Lack of gold"}}
	BaseResultLackOfStarDiamond                = BaseResult{Err{HttpStatus: http.StatusOK, Code: 41007, Message: "Lack of star diamond"}}
	BaseResultInvalidRemark                    = BaseResult{Err{HttpStatus: http.StatusOK, Code: 41008, Message: "Invalid remark"}}
	BaseResultCannotDeleteItem                 = BaseResult{Err{HttpStatus: http.StatusOK, Code: 50001, Message: "Cannot delete item"}}
	BaseResultCannotModifyItem                 = BaseResult{Err{HttpStatus: http.StatusOK, Code: 50002, Message: "Cannot modify item"}}
	BaseResultPermissionDeniedModifyGold       = BaseResult{Err{HttpStatus: http.StatusOK, Code: 50003, Message: "Permission denied modify gold"}}
	BaseResultPermissionDeniedModifySarDiamond = BaseResult{Err{HttpStatus: http.StatusOK, Code: 50004, Message: "Permission denied modify sar diamond"}}
	BaseResultTCloudIM                         = BaseResult{Err{HttpStatus: http.StatusOK, Code: 60001, Message: "Tcloud im"}}
	BaseResultTCloudIMMessage                  = BaseResult{Err{HttpStatus: http.StatusOK, Code: 60002, Message: "Tcloud immessage"}}
	BaseResultMetaServer                       = BaseResult{Err{HttpStatus: http.StatusOK, Code: 60101, Message: "MetaServer"}}
	BaseResultMetaServerInvalidAccessLanguage  = BaseResult{Err{HttpStatus: http.StatusOK, Code: 60102, Message: "Meta server invalid access language"}}
	BaseResultMetaServerDatabaseUpgrading      = BaseResult{Err{HttpStatus: http.StatusOK, Code: 60103, Message: "Meta server database upgrading"}}
	BaseResultMetaServerDatabaseNotFound       = BaseResult{Err{HttpStatus: http.StatusOK, Code: 60104, Message: "Meta server database not found"}}
)

type BaseResult struct {
	Err
}

func NewBaseResult(code int32, message string) *BaseResult {
	return &BaseResult{Err{Code: code, Message: message}}
}
