package model

import (
	"encoding/json"
	"fmt"
)

var (
	BaseResultSuccess = BaseResult{0, "OK"}

	BaseResultInvalidArgument    = BaseResult{400, "Invalid argument"}
	BaseResultFailedPrecondition = BaseResult{400, "Failed precondition"}
	BaseResultOutOfRange         = BaseResult{400, "Out of range"}
	BaseResultUnauthenticated    = BaseResult{401, "Unauthenticated"}
	BaseResultPermissionDenied   = BaseResult{403, "Permission denied"}
	BaseResultNotFound           = BaseResult{404, "Not found"}
	BaseResultAborted            = BaseResult{409, "Aborted"}
	BaseResultAlreadyExists      = BaseResult{409, "Already exists"}
	BaseResultResourceExhausted  = BaseResult{429, "Resource exhausted"}
	BaseResultCancelled          = BaseResult{499, "Cancelled"}
	BaseResultService            = BaseResult{500, "Service error, please contact us"}
	BaseResultDataLoss           = BaseResult{500, "Data loss"}
	BaseResultNetwork            = BaseResult{500, "Network error, please try again later"}
	BaseResultNotImplemented     = BaseResult{501, "Not implemented"}
	BaseResultUnavailable        = BaseResult{503, "Unavailable"}
	BaseResultDeadlineExceeded   = BaseResult{504, "Deadline exceeded"}

	BaseResultAccountNotFound     = BaseResult{40001, "Account not found"}
	BaseResultMoneyNotFound       = BaseResult{40002, "Money not found"}
	BaseResultQuickSlotNotFound   = BaseResult{40003, "Quick slot not found"}
	BaseResultHomelandNotFound    = BaseResult{40004, "Homeland not found"}
	BaseResultInventoryNotFound   = BaseResult{40005, "Inventory not found"}
	BaseResultEquipNotFound       = BaseResult{40006, "Equip not found"}
	BaseResultCharacterNotFound   = BaseResult{40007, "Character not found"}
	BaseResultWorldNotFound       = BaseResult{40008, "World not found"}
	BaseResultItemNotFound        = BaseResult{40009, "Item not found"}
	BaseResultTransactionNotFound = BaseResult{40010, "Transaction not found"}
	BaseResultSurnameNotFound     = BaseResult{40011, "Surname not found"}

	BaseResultLoggedIn             = BaseResult{41001, "Logged in"}
	BaseResultLoggedOut            = BaseResult{41002, "Logged out"}
	BaseResultNotConnected         = BaseResult{41003, "Not connected"}
	BaseResultTransactionCancelled = BaseResult{41004, "Transaction cancelled"}
	BaseResultTransactionUsed      = BaseResult{41005, "Transaction used"}
	BaseResultLackOfGold           = BaseResult{41006, "Lack of gold"}
	BaseResultLackOfStarDiamond    = BaseResult{41007, "Lack of star diamond"}
	BaseResultInvalidRemark        = BaseResult{41008, "Invalid remark"}

	BaseResultCannotDeleteItem                 = BaseResult{50001, "Cannot delete item"}
	BaseResultCannotModifyItem                 = BaseResult{50002, "Cannot modify item"}
	BaseResultPermissionDeniedModifyGold       = BaseResult{50003, "Permission denied modify gold"}
	BaseResultPermissionDeniedModifySarDiamond = BaseResult{50004, "Permission denied modify sar diamond"}

	BaseResultTCloudIM        = BaseResult{60001, "Tcloud im"}
	BaseResultTCloudIMMessage = BaseResult{60002, "Tcloud immessage"}

	BaseResultMetaServer                      = BaseResult{60101, "MetaServer"}
	BaseResultMetaServerInvalidAccessLanguage = BaseResult{60102, "Meta server invalid access language"}
	BaseResultMetaServerDatabaseUpgrading     = BaseResult{60103, "Meta server database upgrading"}
	BaseResultMetaServerDatabaseNotFound      = BaseResult{60104, "Meta server database not found"}
)

type BaseResult struct {
	Code    int
	Message string
}

func (r *BaseResult) Error() string {
	bytes, err := json.Marshal(r)
	if err != nil {
		return fmt.Sprintf("{code: %v, message: %s}", r.Code, r.Message)
	}
	return string(bytes)
}
