package errorcode

import "net/http"

var (
	Forbidden                      = NewForbidden(403000, "the requested operation is forbidden and cannot be completed")
	ForbiddenPostboxExceeded       = NewForbidden(403001, "the number of posts in the postbox has exceeded")
	ForbiddenMemberPostboxExceeded = NewForbidden(403002, "the number of posts in the member's postbox has exceeded")
	ForbiddenMemberBannedUploadMod = NewForbidden(403003, "member is baned from uploading mod")
	// ForbiddenMemberBanedUploadMod Deprecated, see ForbiddenMemberBannedUploadMod
	ForbiddenMemberBanedUploadMod          = ForbiddenMemberBannedUploadMod
	ForbiddenUgcRejected                   = NewForbidden(403004, "ugc is in a rejected state and cannot be sent")
	ForbiddenAddFriendsDailyLimit          = NewForbidden(403005, "exceeded the daily limit of requests to add friends")
	ForbiddenAddingFriendApplicantExceeded = NewForbidden(403006, "the applicant for adding friends has exceeded the maximum number of friends")
	ForbiddenAddingFriendRecipientExceeded = NewForbidden(403007, "the recipient for adding friends has exceeded the maximum number of friends")
	ForbiddenSimulatorRestrictedLogin      = NewForbidden(403008, "simulator restricted login")
	ForbiddenMemberBanedLogin              = NewForbidden(403009, "member baned login")
	ForbiddenMemberRepeatLogin             = NewForbidden(403010, "member repeat login")
	ForbiddenMemberNotInLoginWhitelist     = NewForbidden(403011, "member is not in login whitelist")
	ForbiddenChannelLoginDisabled          = NewForbidden(403012, "channel login disabled")
	ForbiddenMemberCharacterExceeded       = NewForbidden(403013, "The number of characters for member has exceeded")

	ForbiddenMissionNotComplete    = NewForbidden(403401, "mission not completed")
	ForbiddenRewardAlreadyObtained = NewForbidden(403402, "reward has been obtained")
)

type ForbiddenErr struct {
	Err
}

func NewForbidden(code int32, message string) *ForbiddenErr {
	return &ForbiddenErr{Err{Code: code, Message: message, HttpStatus: http.StatusForbidden}}
}
