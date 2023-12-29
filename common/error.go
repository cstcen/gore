package common

import (
	"encoding/json"
	"net/http"
)

var (
	// ErrSuccess http status code
	ErrSuccess           = Err{HttpStatus: http.StatusOK, Code: 0, Message: "OK"}
	ErrInvalidArgument   = Err{HttpStatus: http.StatusBadRequest, Code: 400, Message: "Invalid argument"}
	ErrUnauthenticated   = Err{HttpStatus: http.StatusUnauthorized, Code: 401, Message: "Unauthenticated"}
	ErrPermissionDenied  = Err{HttpStatus: http.StatusForbidden, Code: 403, Message: "Permission denied"}
	ErrNotFound          = Err{HttpStatus: http.StatusNotFound, Code: 404, Message: "Not found"}
	ErrAborted           = Err{HttpStatus: http.StatusConflict, Code: 409, Message: "Aborted"}
	ErrAlreadyExists     = Err{HttpStatus: http.StatusConflict, Code: 409, Message: "Already exists"}
	ErrResourceExhausted = Err{HttpStatus: http.StatusTooManyRequests, Code: 429, Message: "Resource exhausted"}
	ErrCancelled         = Err{HttpStatus: http.StatusBadRequest, Code: 499, Message: "Cancelled"}
	ErrService           = Err{HttpStatus: http.StatusInternalServerError, Code: 500, Message: "Service error, please contact us"}
	ErrNotImplemented    = Err{HttpStatus: http.StatusNotImplemented, Code: 501, Message: "Not implemented"}
	ErrUnavailable       = Err{HttpStatus: http.StatusServiceUnavailable, Code: 503, Message: "Unavailable"}
	ErrDeadlineExceeded  = Err{HttpStatus: http.StatusGatewayTimeout, Code: 504, Message: "Deadline exceeded"}

	// ErrParamsNotFound 参数找不到问题相关
	ErrParamsNotFound      = Err{HttpStatus: http.StatusNotFound, Code: 40000, Message: "Account not found"}
	ErrAccountNotFound     = Err{HttpStatus: http.StatusNotFound, Code: 40001, Message: "Account not found"}
	ErrMoneyNotFound       = Err{HttpStatus: http.StatusNotFound, Code: 40002, Message: "Money not found"}
	ErrQuickSlotNotFound   = Err{HttpStatus: http.StatusNotFound, Code: 40003, Message: "Quick slot not found"}
	ErrHomelandNotFound    = Err{HttpStatus: http.StatusNotFound, Code: 40004, Message: "Homeland not found"}
	ErrInventoryNotFound   = Err{HttpStatus: http.StatusNotFound, Code: 40005, Message: "Inventory not found"}
	ErrEquipNotFound       = Err{HttpStatus: http.StatusNotFound, Code: 40006, Message: "Equip not found"}
	ErrCharacterNotFound   = Err{HttpStatus: http.StatusNotFound, Code: 40007, Message: "Character not found"}
	ErrWorldNotFound       = Err{HttpStatus: http.StatusNotFound, Code: 40008, Message: "World not found"}
	ErrItemNotFound        = Err{HttpStatus: http.StatusNotFound, Code: 40009, Message: "Item not found"}
	ErrTransactionNotFound = Err{HttpStatus: http.StatusNotFound, Code: 40010, Message: "Transaction not found"}
	ErrSurnameNotFound     = Err{HttpStatus: http.StatusNotFound, Code: 40011, Message: "Surname not found"}

	// ErrValidation 验证问题相关
	ErrValidation                  = Err{HttpStatus: http.StatusUnauthorized, Code: 41000, Message: "Logged in"}
	ErrLoggedIn                    = Err{HttpStatus: http.StatusUnauthorized, Code: 41001, Message: "Logged in"}
	ErrLoggedOut                   = Err{HttpStatus: http.StatusUnauthorized, Code: 41002, Message: "Logged out"}
	ErrNotConnected                = Err{HttpStatus: http.StatusUnauthorized, Code: 41003, Message: "Not connected"}
	ErrTransactionCancelled        = Err{HttpStatus: http.StatusUnauthorized, Code: 41004, Message: "Transaction cancelled"}
	ErrTransactionUsed             = Err{HttpStatus: http.StatusUnauthorized, Code: 41005, Message: "Transaction used"}
	ErrLackOfGold                  = Err{HttpStatus: http.StatusUnauthorized, Code: 41006, Message: "Lack of gold"}
	ErrLackOfStarDiamond           = Err{HttpStatus: http.StatusUnauthorized, Code: 41007, Message: "Lack of star diamond"}
	ErrInvalidRemark               = Err{HttpStatus: http.StatusUnauthorized, Code: 41008, Message: "Invalid remark"}
	ErrInvalidClientAuthentication = Err{HttpStatus: http.StatusUnauthorized, Code: 41009, Message: "无效的客户端，请检查所申请的应用ID"}
	ErrInvalidPassword             = Err{HttpStatus: http.StatusUnauthorized, Code: 41010, Message: "账号或密码错误"}
	ErrInvalidRedirectUri          = Err{HttpStatus: http.StatusUnauthorized, Code: 41011, Message: "非法的 redirect_uri"}
	ErrInvalidGrantType            = Err{HttpStatus: http.StatusUnauthorized, Code: 41012, Message: "非法的 grant_type"}
	ErrInvalidAuthorizationCode    = Err{HttpStatus: http.StatusUnauthorized, Code: 41013, Message: "非法的授权码"}
	ErrInvalidRefreshToken         = Err{HttpStatus: http.StatusUnauthorized, Code: 41014, Message: "无效的refresh_token"}
	ErrInvalidReplyingParty        = Err{HttpStatus: http.StatusUnauthorized, Code: 41015, Message: "无效的依赖方"}
	ErrInvalidAuthenticationMode   = Err{HttpStatus: http.StatusUnauthorized, Code: 41016, Message: "无效的认证方式"}
	ErrInvalidResponseType         = Err{HttpStatus: http.StatusUnauthorized, Code: 41017, Message: "无效的响应方式"}
	ErrInvalidVerificationCode     = Err{HttpStatus: http.StatusUnauthorized, Code: 41018, Message: "无效的手机验证码"}

	// ErrPermission 权限问题相关
	ErrPermission                       = Err{HttpStatus: http.StatusForbidden, Code: 50000, Message: "Cannot delete item"}
	ErrCannotDeleteItem                 = Err{HttpStatus: http.StatusForbidden, Code: 50001, Message: "Cannot delete item"}
	ErrCannotModifyItem                 = Err{HttpStatus: http.StatusForbidden, Code: 50002, Message: "Cannot modify item"}
	ErrPermissionDeniedModifyGold       = Err{HttpStatus: http.StatusForbidden, Code: 50003, Message: "Permission denied modify gold"}
	ErrPermissionDeniedModifySarDiamond = Err{HttpStatus: http.StatusForbidden, Code: 50004, Message: "Permission denied modify sar diamond"}

	// ErrTCloud 腾讯云服务相关
	ErrTCloud           = Err{HttpStatus: http.StatusOK, Code: 60000, Message: "Tcloud im"}
	ErrTCloudIM         = Err{HttpStatus: http.StatusOK, Code: 60001, Message: "Tcloud im"}
	ErrTCloudIMMessage  = Err{HttpStatus: http.StatusOK, Code: 60002, Message: "Tcloud immessage"}
	ErrTCloudInvalidSig = Err{HttpStatus: http.StatusOK, Code: 60003, Message: "Tcloud invalid sig"}

	// ErrMetaServer Meta 服务相关
	ErrMetaServer                      = Err{HttpStatus: http.StatusOK, Code: 60101, Message: "MetaServer"}
	ErrMetaServerInvalidAccessLanguage = Err{HttpStatus: http.StatusOK, Code: 60102, Message: "Meta server invalid access language"}
	ErrMetaServerDatabaseUpgrading     = Err{HttpStatus: http.StatusOK, Code: 60103, Message: "Meta server database upgrading"}
	ErrMetaServerDatabaseNotFound      = Err{HttpStatus: http.StatusOK, Code: 60104, Message: "Meta server database not found"}

	// ErrAuthBegin Auth 服务相关 10000 ~ 29999
	ErrAuthBegin                                 = Err{HttpStatus: http.StatusOK, Code: 10000, Message: "Error : Auth Server"}
	ErrAuthSvrUnknown                            = Err{HttpStatus: http.StatusOK, Code: 10001, Message: "Error : Unknown Error."}
	ErrAuthSvrException                          = Err{HttpStatus: http.StatusOK, Code: 10002, Message: "Error : Exception Error."}
	ErrAuthSvrNoData                             = Err{HttpStatus: http.StatusOK, Code: 10003, Message: "Error : No Data."}
	ErrAuthSvrClosedService                      = Err{HttpStatus: http.StatusOK, Code: 10021, Message: "Error : Closed service."}
	ErrAuthSvrUnknownMemberNo                    = Err{HttpStatus: http.StatusOK, Code: 10100, Message: "Error : Not exists member number."}
	ErrAuthSvrDiffPassword                       = Err{HttpStatus: http.StatusOK, Code: 10101, Message: "Error : Password is wrong."}
	ErrAuthSvrUnknownMemberId                    = Err{HttpStatus: http.StatusOK, Code: 10102, Message: "Error : Member ID doesn't exist."}
	ErrAuthSvrFailToGetMemberNo                  = Err{HttpStatus: http.StatusOK, Code: 10103, Message: "Error : Fail to get member number."}
	ErrAuthSvrNicknameIsShort                    = Err{HttpStatus: http.StatusOK, Code: 10104, Message: "Error : Nickname is short."}
	ErrAuthSvrNicknameIsLong                     = Err{HttpStatus: http.StatusOK, Code: 10105, Message: "Error : Nickname is long."}
	ErrAuthSvrExistsNickname                     = Err{HttpStatus: http.StatusOK, Code: 10106, Message: "Error : Nickname already exists."}
	ErrAuthSvrPasswordIsShort                    = Err{HttpStatus: http.StatusOK, Code: 10107, Message: "Error : Password is short."}
	ErrAuthSvrPasswordIsLong                     = Err{HttpStatus: http.StatusOK, Code: 10108, Message: "Error : Password is long."}
	ErrAuthSvrMemberidIsShort                    = Err{HttpStatus: http.StatusOK, Code: 10109, Message: "Error : Member ID is short."}
	ErrAuthSvrMemberidIsLong                     = Err{HttpStatus: http.StatusOK, Code: 10110, Message: "Error : Member ID is long."}
	ErrAuthSvrLoginFailOvercount                 = Err{HttpStatus: http.StatusOK, Code: 10120, Message: "Error : Login fail count is over."}
	ErrAuthSvrAlreadyExistsAccount               = Err{HttpStatus: http.StatusOK, Code: 10121, Message: "Error : Already exists account type"}
	ErrAuthSvrBanMember                          = Err{HttpStatus: http.StatusOK, Code: 10122, Message: "Error : This member is a ban member."}
	ErrAuthSvrBadNickname                        = Err{HttpStatus: http.StatusOK, Code: 10123, Message: "Error : Bad nickname."}
	ErrAuthSvrAlreadyExistsMemberid              = Err{HttpStatus: http.StatusOK, Code: 10124, Message: "Error : Already exists member ID."}
	ErrAuthSvrLoginRestrict                      = Err{HttpStatus: http.StatusOK, Code: 10125, Message: "Error : It works only on normal devices. If this error persists, please contact Customer Service."}
	ErrAuthSvrFailToGetNicknameno                = Err{HttpStatus: http.StatusOK, Code: 10126, Message: "Error : Fail to get nickname number."}
	ErrAuthSvrUnregisterMember                   = Err{HttpStatus: http.StatusOK, Code: 10127, Message: "Error : Unregister member."}
	ErrAuthSvrChangedNicknameAlready             = Err{HttpStatus: http.StatusOK, Code: 10128, Message: "Error : Nickname is changed already."}
	ErrAuthSvrRegisterMember                     = Err{HttpStatus: http.StatusOK, Code: 10129, Message: "Error : Register member."}
	ErrAuthSvrUnknownClientId                    = Err{HttpStatus: http.StatusOK, Code: 10130, Message: "Error : Client ID doesn't exist."}
	ErrAuthSvrTrySendTooMany                     = Err{HttpStatus: http.StatusOK, Code: 10131, Message: "Error : SMS verification try to send too many."}
	ErrAuthSvrSmsAuthWrong                       = Err{HttpStatus: http.StatusOK, Code: 10132, Message: "Error : SMS verification failed."}
	ErrAuthSvrSmsAuthTimeOver                    = Err{HttpStatus: http.StatusOK, Code: 10133, Message: "Error : SMS verification time is over."}
	ErrAuthSvrSmsAuthAttemptOver                 = Err{HttpStatus: http.StatusOK, Code: 10134, Message: "Error : SMS verification to attempt is over."}
	ErrAuthSvrSmsSending                         = Err{HttpStatus: http.StatusOK, Code: 10135, Message: "Error : SMS message didn't send."}
	ErrAuthSvrValidCharacterId                   = Err{HttpStatus: http.StatusOK, Code: 10135, Message: "Error : character id doesn't exist."}
	ErrAuthSvrValidCharacterNickname             = Err{HttpStatus: http.StatusOK, Code: 10136, Message: "Error : nickname is blank."}
	ErrAuthSvrValidGameVersion                   = Err{HttpStatus: http.StatusOK, Code: 10137, Message: "Error : game version is not valid."}
	ErrAuthSvrValidEmailCertification            = Err{HttpStatus: http.StatusOK, Code: 10138, Message: "Error : fail email certification"}
	ErrAuthSvrValidChangePassword                = Err{HttpStatus: http.StatusOK, Code: 10139, Message: "Error : fail change password."}
	ErrAuthSvrCertifiedEmail                     = Err{HttpStatus: http.StatusOK, Code: 10140, Message: "Error : email is certified already."}
	ErrAuthSvrEmailIsBlank                       = Err{HttpStatus: http.StatusOK, Code: 10141, Message: "Error : Email is Blank."}
	ErrAuthSvrRemainTheTimeForEmailCertification = Err{HttpStatus: http.StatusOK, Code: 10142, Message: "Error : remain the time for email certification."}
	ErrAuthSvrNotExistsRequestGameDrop           = Err{HttpStatus: http.StatusOK, Code: 10143, Message: "Error : Game Drop request is not exists."}
	ErrAuthSvrGameDrop                           = Err{HttpStatus: http.StatusOK, Code: 10144, Message: "Error : Game Drop is failed."}
	ErrAuthSvrNicknameDelUpdate                  = Err{HttpStatus: http.StatusOK, Code: 10145, Message: "Error : Nickname delete is failed."}
	ErrAuthSvrGameDropCancel                     = Err{HttpStatus: http.StatusOK, Code: 10146, Message: "Error : Withdrawal recovered."}
	ErrAuthSvrAlreadyGameDrop                    = Err{HttpStatus: http.StatusOK, Code: 10147, Message: "Error : Withdrawal has already been completed."}
	ErrAuthSvrNotExistsLoginLog                  = Err{HttpStatus: http.StatusOK, Code: 10148, Message: "Error : Login log is not exists."}
	ErrAuthSvrLoginDateFormat                    = Err{HttpStatus: http.StatusOK, Code: 10149, Message: "Error : Login date format error."}
	ErrAuthSvrMigrationAlreadyLinked             = Err{HttpStatus: http.StatusOK, Code: 10150, Message: "Error : Already been linked."}
	ErrAuthSvrNotExistsMigrationId               = Err{HttpStatus: http.StatusOK, Code: 10151, Message: "Error : Migration id is not exists."}
	ErrAuthSvrValidGameId                        = Err{HttpStatus: http.StatusOK, Code: 10152, Message: "Error : game id is not valid."}
	ErrAuthSvrJoinDateFormat                     = Err{HttpStatus: http.StatusOK, Code: 10153, Message: "Error : Join date format error."}
	ErrAuthSvrNotExistsJoinLog                   = Err{HttpStatus: http.StatusOK, Code: 10154, Message: "Error : Join log is not exists."}
	ErrAuthSvrNotRegisterWhitelist               = Err{HttpStatus: http.StatusOK, Code: 10155, Message: "Error : The device is not registered on whitelist."}
	ErrAuthSvrInvalidRegisterAccountType         = Err{HttpStatus: http.StatusOK, Code: 10156, Message: "Error : Register account type is invalid"}
	ErrAuthSvrNotExistsPushKey                   = Err{HttpStatus: http.StatusOK, Code: 10157, Message: "Error : Push key is not exists."}
	ErrAuthSvrPhoneNoSizeInvalid                 = Err{HttpStatus: http.StatusOK, Code: 10158, Message: "Error : Phone no size not valid."}
	ErrAuthSvrPhoneCorpSizeInvalid               = Err{HttpStatus: http.StatusOK, Code: 10159, Message: "Error : Phone corp size not valid."}
	ErrAuthSvrBirthdaySizeInvalid                = Err{HttpStatus: http.StatusOK, Code: 10160, Message: "Error : Birthday size not valid."}
	ErrAuthSvrGenderSizeInvalid                  = Err{HttpStatus: http.StatusOK, Code: 10161, Message: "Error : Gender size not valid."}
	ErrAuthSvrNationSizeInvalid                  = Err{HttpStatus: http.StatusOK, Code: 10162, Message: "Error : Nation size not valid."}
	ErrAuthSvrNameSizeInvalid                    = Err{HttpStatus: http.StatusOK, Code: 10163, Message: "Error : Name size not valid."}
	ErrAuthSvrResultSizeInvalid                  = Err{HttpStatus: http.StatusOK, Code: 10164, Message: "Error : Result size not valid."}
	ErrAuthSvrCertMetSizeInvalid                 = Err{HttpStatus: http.StatusOK, Code: 10165, Message: "Error : Cert met size not valid."}
	ErrAuthSvrIpSizeInvalid                      = Err{HttpStatus: http.StatusOK, Code: 10166, Message: "Error : IP size not valid."}
	ErrAuthSvrDiSizeInvalid                      = Err{HttpStatus: http.StatusOK, Code: 10167, Message: "Error : DI size not valid."}
	ErrAuthSvrPasswordWrongPattern               = Err{HttpStatus: http.StatusOK, Code: 10168, Message: "비밀번호는 안전을 위해 영문,숫자가 조합된 8~64자 이내의 문자로 생성해 주세요"}
	ErrAuthSvrNotExistsFacebookAccountInfo       = Err{HttpStatus: http.StatusOK, Code: 10169, Message: "Error : Facebook account info is not exists."}
	ErrAuthSvrGuestMember                        = Err{HttpStatus: http.StatusOK, Code: 10170, Message: "Error : This member is guest."}
	ErrAuthSvrDidNotChangePassword               = Err{HttpStatus: http.StatusOK, Code: 10171, Message: "Error : This member did not change password."}
	ErrAuthSvrGameDropRequest                    = Err{HttpStatus: http.StatusOK, Code: 10172, Message: "Error : The withdrawal request members."}
	ErrAuthSvrAlreadyMemberDropCancel            = Err{HttpStatus: http.StatusOK, Code: 10173, Message: "Error : Withdrawal has already been canceled."}
	ErrAuthSvrNotPlatformMember                  = Err{HttpStatus: http.StatusOK, Code: 10174, Message: "Error : This member is not xk5 platform member."}
	ErrAuthSvrAlreadyPlatformDrop                = Err{HttpStatus: http.StatusOK, Code: 10175, Message: "Error : Withdrawal a member"}
	ErrAuthSvrNotExistsPolicy                    = Err{HttpStatus: http.StatusOK, Code: 10176, Message: "Error : Not exists policy."}
	ErrAuthSvrAlreadyAgreedPolicy                = Err{HttpStatus: http.StatusOK, Code: 10177, Message: "Error : Already agreed to the policy."}
	ErrAuthSvrHoldWithdrawMember                 = Err{HttpStatus: http.StatusOK, Code: 10178, Message: "탈퇴 신청이 완료되어 삭제될 예정입니다.\n만약 취소를 원하시는 경우 고객센터로 문의해주세요."}
	ErrAuthSvrAlreadyAgreedGameTerms             = Err{HttpStatus: http.StatusOK, Code: 10179, Message: "Error : Already agreed to the game term."}
	ErrAuthSvrGuestLogoutFail                    = Err{HttpStatus: http.StatusOK, Code: 10180, Message: "Error : Guest does not provide logout."}
	ErrAuthSvrAlreadyStoveAcoountLink            = Err{HttpStatus: http.StatusOK, Code: 10181, Message: "Error : Already xk5 Account Link."}
	ErrAuthSvrAccountTypeSimple                  = Err{HttpStatus: http.StatusOK, Code: 10182, Message: "Error : You must upgrade to full account first."}
	ErrAuthSvrEmailVerificationExceeded          = Err{HttpStatus: http.StatusOK, Code: 10183, Message: "Error : Email verification count exceeded."}
	ErrAuthSvrLinkLimitExceeded                  = Err{HttpStatus: http.StatusOK, Code: 10190, Message: "Error : You have exceeded the maximum overwrite attempts. Please link another account."}
	ErrAuthSvrLinkBanMember                      = Err{HttpStatus: http.StatusOK, Code: 10191, Message: "Error : You cannot link a banned account."}
	ErrAuthSvrServiceUseAgeRestirct              = Err{HttpStatus: http.StatusOK, Code: 10192, Message: "Error : The age at which the service is unavailable."}
	ErrAuthSvrMgsBroken                          = Err{HttpStatus: http.StatusOK, Code: 11000, Message: "Error : Message is broken."}
	ErrAuthSvrArgIsWrong                         = Err{HttpStatus: http.StatusOK, Code: 11001, Message: "Error : Argument is wrong."}
	ErrAuthSvrUnknownMsgid                       = Err{HttpStatus: http.StatusOK, Code: 11002, Message: "Error : Unknown Message ID."}
	ErrAuthSvrAltIsNull                          = Err{HttpStatus: http.StatusOK, Code: 11010, Message: "Error : ALT is NULL."}
	ErrAuthSvrAltException                       = Err{HttpStatus: http.StatusOK, Code: 11011, Message: "Error : ALT is wrong."}
	ErrAuthSvrAltTimeout                         = Err{HttpStatus: http.StatusOK, Code: 11012, Message: "Error : ALT is timeout."}
	ErrAuthSvrAltDecryptError                    = Err{HttpStatus: http.StatusOK, Code: 11013, Message: "Error : ALT is wrong."}
	ErrAuthSvrAltItemCountWrong                  = Err{HttpStatus: http.StatusOK, Code: 11014, Message: "Error : ALT is wrong."}
	ErrAuthSvrAltDiffMemberno                    = Err{HttpStatus: http.StatusOK, Code: 11020, Message: "Error : ALT is wrong."}
	ErrAuthSvrAltDiffDeviceid                    = Err{HttpStatus: http.StatusOK, Code: 11021, Message: "Error : ALT is wrong."}
	ErrAuthSvrAltDiffStagestring                 = Err{HttpStatus: http.StatusOK, Code: 11022, Message: "Error : ALT is wrong."}
	ErrAuthSvrAltCreateFail                      = Err{HttpStatus: http.StatusOK, Code: 11030, Message: "Error : ALT can't create."}
	ErrAuthSvrStIsNull                           = Err{HttpStatus: http.StatusOK, Code: 11110, Message: "Error : ST is NULL."}
	ErrAuthSvrStException                        = Err{HttpStatus: http.StatusOK, Code: 11111, Message: "Error : ST is wrong."}
	ErrAuthSvrStTimeout                          = Err{HttpStatus: http.StatusOK, Code: 11112, Message: "Error : ST is timeout."}
	ErrAuthSvrStDecryptError                     = Err{HttpStatus: http.StatusOK, Code: 11113, Message: "Error : ST is wrong."}
	ErrAuthSvrStItemCountWrong                   = Err{HttpStatus: http.StatusOK, Code: 11114, Message: "Error : ST is wrong."}
	ErrAuthSvrStDiffMemberno                     = Err{HttpStatus: http.StatusOK, Code: 11020, Message: "Error : ST is wrong."}
	ErrAuthSvrStCreateFail                       = Err{HttpStatus: http.StatusOK, Code: 11030, Message: "Error : ST can't create."}
	ErrAuthSvrSstIsNull                          = Err{HttpStatus: http.StatusOK, Code: 11210, Message: "Error : SST is NULL."}
	ErrAuthSvrSstException                       = Err{HttpStatus: http.StatusOK, Code: 11211, Message: "Error : SST is wrong."}
	ErrAuthSvrSstTimeout                         = Err{HttpStatus: http.StatusOK, Code: 11212, Message: "Error : SST is timeout."}
	ErrAuthSvrSstDecryptError                    = Err{HttpStatus: http.StatusOK, Code: 11213, Message: "Error : SST is wrong."}
	ErrAuthSvrSstItemCountWrong                  = Err{HttpStatus: http.StatusOK, Code: 11214, Message: "Error : SST is wrong."}
	ErrAuthSvrSstDiffMemberno                    = Err{HttpStatus: http.StatusOK, Code: 11220, Message: "Error : SST is wrong."}
	ErrAuthSvrSstCreateFail                      = Err{HttpStatus: http.StatusOK, Code: 11230, Message: "Error : SST can't create."}
	ErrAuthSvrSstUnknown                         = Err{HttpStatus: http.StatusOK, Code: 11231, Message: "Error : SST doesn't exist."}
	ErrAuthSvrSstAlreadyUsed                     = Err{HttpStatus: http.StatusOK, Code: 11232, Message: "Error : SST is already used."}
	ErrAuthSvrSstDiff                            = Err{HttpStatus: http.StatusOK, Code: 11233, Message: "Error : SST is different in DB."}
	ErrAuthSvrDecryptError                       = Err{HttpStatus: http.StatusOK, Code: 11234, Message: "Error : Decrypte error."}
	ErrAuthSvrRefreshTokenException              = Err{HttpStatus: http.StatusOK, Code: 11235, Message: "Error : Refresh token is wrong."}
	ErrAuthSvrAccessTokenException               = Err{HttpStatus: http.StatusOK, Code: 11236, Message: "Error : Access token is wrong."}
	ErrAuthSvrGameAccessTokenException           = Err{HttpStatus: http.StatusOK, Code: 11237, Message: "Error : Game Access Token is wrong."}
	ErrAuthSvrOauthAuthorizationExcetpion        = Err{HttpStatus: http.StatusOK, Code: 11238, Message: "Error : OAuth Authorization is wrong."}
	ErrAuthSvrMemberSleep                        = Err{HttpStatus: http.StatusOK, Code: 11239, Message: "Error : This member is a sleep member."}
	ErrAuthSvrPcGameAccessTokenMissing           = Err{HttpStatus: http.StatusOK, Code: 11240, Message: "Error : PC Game Access Token missing."}
	ErrAuthSvrPcGameAccessTokenException         = Err{HttpStatus: http.StatusOK, Code: 11241, Message: "Error : PC Game Access Token is wrong."}
	ErrAuthSvrPcWegameAccessTokenException       = Err{HttpStatus: http.StatusOK, Code: 11242, Message: "Error : PC WeGame Access Token is wrong."}
	ErrAuthSvrNotInEditorWhiteList               = Err{HttpStatus: http.StatusOK, Code: 11243, Message: "Error : the user not in editor white list."}
	ErrAuthSvrOtherDeviceAlreadyLogin4Editor     = Err{HttpStatus: http.StatusOK, Code: 11244, Message: "Error : the account has logged in on another device."}
	ErrAuthSvrOtherDeviceForceLogin4Editor       = Err{HttpStatus: http.StatusOK, Code: 11245, Message: "Error : the account has logged in on another device,your login is invalid"}
	ErrAuthSvrNetLinkIsNull                      = Err{HttpStatus: http.StatusOK, Code: 12000, Message: "Error : Link Info is NULL."}
	ErrAuthSvrDbDisconnect                       = Err{HttpStatus: http.StatusOK, Code: 12001, Message: "Error : DB connecion error."}
	ErrAuthSvrDbException                        = Err{HttpStatus: http.StatusOK, Code: 12002, Message: "Error : Data access exception."}
	ErrAuthSvrGameIdIsNull                       = Err{HttpStatus: http.StatusOK, Code: 12003, Message: "Error : Game Id is error"}
	ErrAuthSvrInvalidAccountTypeCd               = Err{HttpStatus: http.StatusOK, Code: 13000, Message: "Error : Invalid Account Type CD"}
	ErrAuthSvrUnknownMemberNo2                   = Err{HttpStatus: http.StatusOK, Code: 13001, Message: "Error : Not exists Palmple member number."}
	ErrAuthSvrDropMemberNo                       = Err{HttpStatus: http.StatusOK, Code: 13002, Message: "Error : Drop member."}
	ErrAuthSvrAlreadyTransferStove               = Err{HttpStatus: http.StatusOK, Code: 13003, Message: "Error : Already transfer account."}
	ErrAuthSvrAlreadyGameAccount                 = Err{HttpStatus: http.StatusOK, Code: 13004, Message: "Error : Already existing game account."}
	ErrAuthSvrTicketParseException               = Err{HttpStatus: http.StatusOK, Code: 13005, Message: "Error : Ticket is wrong."}
	ErrAuthSvrAlreadyTransferAccountRegister     = Err{HttpStatus: http.StatusOK, Code: 13006, Message: "Error : Already transfer account regist"}
	ErrAuthSvrAlreadyGameInfo                    = Err{HttpStatus: http.StatusOK, Code: 13007, Message: "Error : Already Game Info"}
	ErrAuthSvrMandatoryMissing                   = Err{HttpStatus: http.StatusOK, Code: 13008, Message: "Error : Mandatory Parameter missing"}
	ErrAuthSvrAlreadyTransfer                    = Err{HttpStatus: http.StatusOK, Code: 13009, Message: "Error : Already transfer account."}
	ErrAuthSvrInvalid                            = Err{HttpStatus: http.StatusOK, Code: 14000, Message: "Error : Invalid nickname."}
	ErrAuthSvrBannedWord                         = Err{HttpStatus: http.StatusOK, Code: 14001, Message: "Error : It contains a banned word."}
	ErrAuthSvrDuplicate                          = Err{HttpStatus: http.StatusOK, Code: 14002, Message: "Error : Duplicated nickname."}
	ErrAuthSvrDeactivateService                  = Err{HttpStatus: http.StatusOK, Code: 15000, Message: "Error : Deactivate Service Member"}
	ErrAuthSvrExistsCharacter                    = Err{HttpStatus: http.StatusOK, Code: 20050, Message: "Error : The member's character information already exists."}
	ErrAuthSvrNotWhiteMember                     = Err{HttpStatus: http.StatusOK, Code: 20051, Message: "Error : Sorry, you are not eligible for testing."}
	ErrAuthSvrInvalidProviderId                  = Err{HttpStatus: http.StatusOK, Code: 20052, Message: "Error : Invalid login ID."}
	ErrAuthSvrInvalidMappingGameId               = Err{HttpStatus: http.StatusOK, Code: 20053, Message: "Error : Invalid Mapping Game ID."}
	ErrAuthSvrNotExistsMappingGameId             = Err{HttpStatus: http.StatusOK, Code: 20054, Message: "Error : Not exist Mapping Game ID."}
	ErrAuthSvrAccountLimitExceeded               = Err{HttpStatus: http.StatusOK, Code: 20055, Message: "Error : The number of accounts has exceeded the maximum"}
	ErrAuthSvrSimulatorLogin                     = Err{HttpStatus: http.StatusOK, Code: 20056, Message: "Error : Sorry, The simulator cannot log in."}
	ErrAuthSvrRepeatLogin                        = Err{HttpStatus: http.StatusOK, Code: 20057, Message: "Error : repeat login."}
	ErrAuthSvrChannelWhitelistDisabled           = Err{HttpStatus: http.StatusOK, Code: 20058, Message: "Error : channel whitelist disabled."}
	ErrAuthSvrNotOnChannelWhitelist              = Err{HttpStatus: http.StatusOK, Code: 20059, Message: "Error : not on the channel whitelist."}
	ErrIdipError                                 = Err{HttpStatus: http.StatusOK, Code: -299, Message: "Error : UNKNOWN ERROR"}
	ErrIdipMandatoryParameterMissing             = Err{HttpStatus: http.StatusOK, Code: -201, Message: "Error : MANDATORY PARAMETER MISSING"}
	ErrIdipNotExistMember                        = Err{HttpStatus: http.StatusOK, Code: -202, Message: "Error : NOT EXIST MEMBER"}
	ErrIdipNotExistCharacter                     = Err{HttpStatus: http.StatusOK, Code: -203, Message: "Error : NOT EXIST GAME CHARACTER"}
	ErrIdipDiscordantPlatformType                = Err{HttpStatus: http.StatusOK, Code: -204, Message: "Error : DISCORDANT PLATFORM TYPE"}

	// ErrPostboxWrongApiUsage Postbox 返回码
	ErrPostboxWrongApiUsage                   = Err{HttpStatus: http.StatusOK, Code: -1, Message: "Wrong API Usage"}
	ErrPostboxServiceError                    = Err{HttpStatus: http.StatusOK, Code: -2, Message: "Service Error"}
	ErrPostboxAuthFail                        = Err{HttpStatus: http.StatusOK, Code: -3, Message: "Authentication Failure"}
	ErrPostboxUserNotFound                    = Err{HttpStatus: http.StatusOK, Code: -4, Message: "User NotFound"}
	ErrPostboxUnauthorized                    = Err{HttpStatus: http.StatusOK, Code: -5, Message: "Unauthorized"}
	ErrPostboxExceededPostboxSize             = Err{HttpStatus: http.StatusOK, Code: -6, Message: "Exceeded Postbox Size"}
	ErrPostboxExceededMaxSizeOfRecipients     = Err{HttpStatus: http.StatusOK, Code: -7, Message: "Exceeded Recipient Max Size"}
	ErrPostboxPartialSuccess                  = Err{HttpStatus: http.StatusOK, Code: -100, Message: "Partial Success"}
	ErrPostboxAllFailed                       = Err{HttpStatus: http.StatusOK, Code: -101, Message: "All Failed"}
	ErrPostboxDuplicatePostboxId              = Err{HttpStatus: http.StatusOK, Code: -90001, Message: "Duplicate PostboxId"}
	ErrPostboxPostboxNotFound                 = Err{HttpStatus: http.StatusOK, Code: -90002, Message: "Postbox NotFound"}
	ErrPostboxDuplicateServiceId              = Err{HttpStatus: http.StatusOK, Code: -90003, Message: "Duplicate ServiceId"}
	ErrPostboxServiceNotFound                 = Err{HttpStatus: http.StatusOK, Code: -90004, Message: "ServiceInfo NotFound"}
	ErrPostboxInvalidRecipients               = Err{HttpStatus: http.StatusOK, Code: -90005, Message: "Invalid Recipients"}
	ErrPostboxInvalidTransactionId            = Err{HttpStatus: http.StatusOK, Code: -90006, Message: "Invalid Transaction Id"}
	ErrPostboxAlreadySentItem                 = Err{HttpStatus: http.StatusOK, Code: -90007, Message: "Already Sent Item"}
	ErrPostboxNotAttachedItem                 = Err{HttpStatus: http.StatusOK, Code: -90008, Message: "Not Attached Item"}
	ErrPostboxDuplicateTransactionId          = Err{HttpStatus: http.StatusOK, Code: -90009, Message: "Duplicate Transaction Id"}
	ErrPostboxPostNotFound                    = Err{HttpStatus: http.StatusOK, Code: -90010, Message: "Post NotFound"}
	ErrPostboxInvalidReservationTime          = Err{HttpStatus: http.StatusOK, Code: -90012, Message: "Invalid ReservationTime"}
	ErrPostboxReservationPostNotFound         = Err{HttpStatus: http.StatusOK, Code: -90013, Message: "Reservation Post NotFound"}
	ErrPostboxAlreadySentReservationPost      = Err{HttpStatus: http.StatusOK, Code: -90014, Message: "Already Sent Reservation Post"}
	ErrPostboxCanceledReservationPost         = Err{HttpStatus: http.StatusOK, Code: -90015, Message: "Canceled Reservation Post"}
	ErrPostboxUserPostNotFound                = Err{HttpStatus: http.StatusOK, Code: -90016, Message: "User Post NotFound"}
	ErrPostboxInvalidExpireAt                 = Err{HttpStatus: http.StatusOK, Code: -90017, Message: "Invalid Expire Time"}
	ErrPostboxNotAllowModifyPost              = Err{HttpStatus: http.StatusOK, Code: -90018, Message: "Not Allow Modify Post"}
	ErrPostboxMoreThan1Recipient              = Err{HttpStatus: http.StatusOK, Code: -90019, Message: "There is more than 1 recipient"}
	ErrPostboxWrongPageNumber                 = Err{HttpStatus: http.StatusOK, Code: -90020, Message: "Page number must be greater than or equal to 1."}
	ErrPostboxUserIdIsMandatoryForAllType     = Err{HttpStatus: http.StatusOK, Code: -90021, Message: "User ID is mandatory for all-type post."}
	ErrPostboxGdisWrongApiUsage               = Err{HttpStatus: http.StatusOK, Code: -5001, Message: "Wrong API Usage"}
	ErrPostboxGdisServiceError                = Err{HttpStatus: http.StatusOK, Code: -5002, Message: "Service Error"}
	ErrPostboxGdisAuthFail                    = Err{HttpStatus: http.StatusOK, Code: -5003, Message: "Authentication Failure"}
	ErrPostboxGdisUserNotFound                = Err{HttpStatus: http.StatusOK, Code: -5004, Message: "User NotFound"}
	ErrPostboxGdisPostNotFound                = Err{HttpStatus: http.StatusOK, Code: -5100, Message: "Post NotFound"}
	ErrPostboxGdisDuplicateTransactionId      = Err{HttpStatus: http.StatusOK, Code: -5101, Message: "Duplicate Transaction Id"}
	ErrPostboxGdisPostboxNotFound             = Err{HttpStatus: http.StatusOK, Code: -5102, Message: "Postbox NotFound"}
	ErrPostboxGdisServiceNotFound             = Err{HttpStatus: http.StatusOK, Code: -5103, Message: "ServiceInfo NotFound"}
	ErrPostboxGdisExceededPostboxSize         = Err{HttpStatus: http.StatusOK, Code: -5106, Message: "Exceeded Postbox Size"}
	ErrPostboxGdisUserIdIsMandatoryForAllType = Err{HttpStatus: http.StatusOK, Code: -5107, Message: "User ID is mandatory for all-type post."}
	ErrPostboxGdisNotAttachedItem             = Err{HttpStatus: http.StatusOK, Code: -5108, Message: "Not Attached Item"}
	ErrPostboxGdisDeletedPost                 = Err{HttpStatus: http.StatusOK, Code: -5109, Message: "Deleted Post"}
	ErrPostboxGdisCouldNotCheck               = Err{HttpStatus: http.StatusOK, Code: -5110, Message: "Could Not Check"}
	ErrPostboxGdisWrongPageNumber             = Err{HttpStatus: http.StatusOK, Code: -5111, Message: "Page number must be greater than or equal to 1."}
	ErrPostboxGdisReservationPostNotFound     = Err{HttpStatus: http.StatusOK, Code: -5112, Message: "Reservation Post NotFound"}
	ErrPostboxGdisAlreadySentReservationPost  = Err{HttpStatus: http.StatusOK, Code: -5113, Message: "Already Sent Reservation Post"}
	ErrPostboxGdisCanceledReservationPost     = Err{HttpStatus: http.StatusOK, Code: -5114, Message: "Canceled Reservation Post"}

	// ErrUgcSuccess UGC返回码
	ErrUgcSuccess                               = Err{HttpStatus: http.StatusOK, Code: 000, Message: "OK"}
	ErrUgcUnknown                               = Err{HttpStatus: http.StatusOK, Code: 501, Message: "Internal server handling error"}
	ErrUgcInvalidHandshakingTicket              = Err{HttpStatus: http.StatusOK, Code: 552, Message: "Invalid handshaking ticket"}
	ErrUgcInvalidToken                          = Err{HttpStatus: http.StatusOK, Code: 553, Message: "Invalid access token : %s"}
	ErrUgcInfraTokenGetError                    = Err{HttpStatus: http.StatusOK, Code: 555, Message: "Failure to acquire infra Token  : %s"}
	ErrUgcMemberInfoNull                        = Err{HttpStatus: http.StatusOK, Code: 505, Message: "Get Member Info error memberNo : %s"}
	ErrUgcMemberServerError                     = Err{HttpStatus: http.StatusOK, Code: 506, Message: "call member server failed  : %s"}
	ErrUgcMemberNoNull                          = Err{HttpStatus: http.StatusOK, Code: 507, Message: "UGC member no is null"}
	ErrUgcMemberPermissionError                 = Err{HttpStatus: http.StatusOK, Code: 508, Message: "This user does not have permission to operate"}
	ErrUgcCallBanwordError                      = Err{HttpStatus: http.StatusOK, Code: 510, Message: "call banword v2 service error : %s"}
	ErrUgcCallBanImageServerError               = Err{HttpStatus: http.StatusOK, Code: 511, Message: "Call BanImage Server error : %s"}
	ErrUgcMemberBanUpMod                        = Err{HttpStatus: http.StatusOK, Code: 512, Message: "member ban upload mod : %s"}
	ErrUgcExist                                 = Err{HttpStatus: http.StatusOK, Code: 601, Message: "UGC existing : %s "}
	ErrUgcNotFind                               = Err{HttpStatus: http.StatusOK, Code: 603, Message: "UGC is not find UGC ID: %s"}
	ErrUgcTeamNotFind                           = Err{HttpStatus: http.StatusOK, Code: 604, Message: "UGC Team is not find"}
	ErrUgcDataNull                              = Err{HttpStatus: http.StatusOK, Code: 605, Message: "UGC Data null"}
	ErrUgcParameterError                        = Err{HttpStatus: http.StatusOK, Code: 606, Message: "UGC Param error : %s"}
	ErrUgcVersionExist                          = Err{HttpStatus: http.StatusOK, Code: 612, Message: "This is the current version"}
	ErrUgcVersionNotFind                        = Err{HttpStatus: http.StatusOK, Code: 613, Message: "UGC version is not find"}
	ErrUgcGameIdNull                            = Err{HttpStatus: http.StatusOK, Code: 614, Message: "UGC gameId is null"}
	ErrUgcIdNull                                = Err{HttpStatus: http.StatusOK, Code: 616, Message: "UGC ID is null! "}
	ErrUgcIdIllegal                             = Err{HttpStatus: http.StatusOK, Code: 616, Message: "illegal UGC ID : %s"}
	ErrUgcIdGenerate                            = Err{HttpStatus: http.StatusOK, Code: 616, Message: "UGC ID Generate failure"}
	ErrUgcRejectedStatus                        = Err{HttpStatus: http.StatusOK, Code: 617, Message: "UGC rejected status Can't send"}
	ErrUgcVersionListNull                       = Err{HttpStatus: http.StatusOK, Code: 618, Message: "UGC Version List is null  UGC ID : %s"}
	ErrUgcTypeError                             = Err{HttpStatus: http.StatusOK, Code: 621, Message: "UGC Type error"}
	ErrUgcGenreError                            = Err{HttpStatus: http.StatusOK, Code: 622, Message: "UGC Genre error : %s"}
	ErrUgcNotMod                                = Err{HttpStatus: http.StatusOK, Code: 623, Message: "UGC is not MOD Can't start the game : %s"}
	ErrUgcDesignSandboxMetaParamNull            = Err{HttpStatus: http.StatusOK, Code: 624, Message: "UGC Design required sandbox meta param is null : %s"}
	ErrUgcModSandboxMetaParamNull               = Err{HttpStatus: http.StatusOK, Code: 625, Message: "UGC MOD required sandbox meta param is null : %s"}
	ErrUgcMediaCheckError                       = Err{HttpStatus: http.StatusOK, Code: 627, Message: "UGC media check fail : %s"}
	ErrUgcSystemModUnableRecommend              = Err{HttpStatus: http.StatusOK, Code: 628, Message: "System MOD unable recommend"}
	ErrUgcFavoriteExist                         = Err{HttpStatus: http.StatusOK, Code: 650, Message: "Favorite exist!"}
	ErrUgcFavoriteNotExist                      = Err{HttpStatus: http.StatusOK, Code: 651, Message: "favorite not exist"}
	ErrUgcEvaluateNotFind                       = Err{HttpStatus: http.StatusOK, Code: 680, Message: "UGC evaluate in not find"}
	ErrUgcEvaluateCreateError                   = Err{HttpStatus: http.StatusOK, Code: 902, Message: "evaluate create error : %s"}
	ErrUgcParamsStatusError                     = Err{HttpStatus: http.StatusOK, Code: 701, Message: "Request params status error !"}
	ErrUgcParamsTypeError                       = Err{HttpStatus: http.StatusOK, Code: 702, Message: "Request params type error : %s"}
	ErrUgcTagExist                              = Err{HttpStatus: http.StatusOK, Code: 801, Message: "Tag is exist"}
	ErrUgcTagInfoError                          = Err{HttpStatus: http.StatusOK, Code: 802, Message: "Tag information error"}
	ErrUgcTagNotFind                            = Err{HttpStatus: http.StatusOK, Code: 803, Message: "Tag not find : %s"}
	ErrUgcGenreIdExist                          = Err{HttpStatus: http.StatusOK, Code: 805, Message: "UGC genre id exist : %s"}
	ErrUgcGenreNameNotFind                      = Err{HttpStatus: http.StatusOK, Code: 806, Message: "UGC genre name not find : %s"}
	ErrUgcGenreTypeError                        = Err{HttpStatus: http.StatusOK, Code: 807, Message: "UGC genre type error : %s"}
	ErrUgcGenreNameNull                         = Err{HttpStatus: http.StatusOK, Code: 808, Message: "UGC genre name is null"}
	ErrUgcGenreIdNull                           = Err{HttpStatus: http.StatusOK, Code: 809, Message: "UGC genre id is null"}
	ErrUgcGenreNull                             = Err{HttpStatus: http.StatusOK, Code: 810, Message: "genre is null"}
	ErrUgcTagOperationTypeExist                 = Err{HttpStatus: http.StatusOK, Code: 811, Message: "Tag Is already operation tag"}
	ErrUgcFileUnpackTheError                    = Err{HttpStatus: http.StatusOK, Code: 1001, Message: "UGC ZIP File Unpack the error"}
	ErrUgcJsonFileError                         = Err{HttpStatus: http.StatusOK, Code: 1002, Message: "UGC Json File error"}
	ErrUgcFileNotFind                           = Err{HttpStatus: http.StatusOK, Code: 1003, Message: "UGC JSON File Not Find"}
	ErrUgcFileUploadFail                        = Err{HttpStatus: http.StatusOK, Code: 1004, Message: "UGC File Upload Fail : %s"}
	ErrUgcFileExportFail                        = Err{HttpStatus: http.StatusOK, Code: 1005, Message: "UGC File Export Fail"}
	ErrUgcFileTypeParamError                    = Err{HttpStatus: http.StatusOK, Code: 1006, Message: "UGC File Type Parameters error : %s"}
	ErrUgcCreateZipFileError                    = Err{HttpStatus: http.StatusOK, Code: 1007, Message: "UGC Create zip file error"}
	ErrUgcFileDownloadError                     = Err{HttpStatus: http.StatusOK, Code: 1008, Message: "UGC File download error : %s"}
	ErrUgcFileUploadNotFindError                = Err{HttpStatus: http.StatusOK, Code: 1009, Message: "ugc.json File Not Find "}
	ErrUgcResourcesFileNotFind                  = Err{HttpStatus: http.StatusOK, Code: 1010, Message: "UGC Resources File Not Find !"}
	ErrUgcUserNotExist                          = Err{HttpStatus: http.StatusOK, Code: 1205, Message: "GMT user not exist : %s"}
	ErrUgcUserNumberNull                        = Err{HttpStatus: http.StatusOK, Code: 1207, Message: "GMT user number is null"}
	ErrUgcReviewAuditorNotExist                 = Err{HttpStatus: http.StatusOK, Code: 1401, Message: "The auditor not exist  : %s"}
	ErrUgcReviewIdError                         = Err{HttpStatus: http.StatusOK, Code: 1402, Message: "reviewId error  : %s"}
	ErrUgcReviewError                           = Err{HttpStatus: http.StatusOK, Code: 1403, Message: "The audit is in progress  : %s"}
	ErrUgcReviewBatchExist                      = Err{HttpStatus: http.StatusOK, Code: 1404, Message: "The audit is in progress  : %s"}
	ErrUgcReviewStatusError                     = Err{HttpStatus: http.StatusOK, Code: 1405, Message: "Review status error  : %s"}
	ErrUgcReviewStatusNull                      = Err{HttpStatus: http.StatusOK, Code: 1406, Message: "Review status is null"}
	ErrUgcReviewExist                           = Err{HttpStatus: http.StatusOK, Code: 1407, Message: "Have received version is reviewing, can't repeat submit  reviewing  versionId: %s"}
	ErrUgcReviewReplyTemplateNotExist           = Err{HttpStatus: http.StatusOK, Code: 1408, Message: "review reply template not exist  templateId : %s"}
	ErrUgcReviewAutoError                       = Err{HttpStatus: http.StatusOK, Code: 1409, Message: "Automatic allocation is ongoing"}
	ErrUgcReviewAutoTypeNull                    = Err{HttpStatus: http.StatusOK, Code: 1410, Message: "Automatically assigned type is null   type : %s"}
	ErrUgcSearchTypeNull                        = Err{HttpStatus: http.StatusOK, Code: 1601, Message: "Search Type is null"}
	ErrUgcSearchTypeError                       = Err{HttpStatus: http.StatusOK, Code: 1602, Message: "Search Type error : %s"}
	ErrUgcArcadeGameNull                        = Err{HttpStatus: http.StatusOK, Code: 1701, Message: "Arcade info is null : %s"}
	ErrUgcArcadeGameNumberExist                 = Err{HttpStatus: http.StatusOK, Code: 1702, Message: "Arcade Number exist: %s"}
	ErrUgcTeamMemberContains                    = Err{HttpStatus: http.StatusOK, Code: 1901, Message: "The member already exists : %s"}
	ErrUgcTeamMemberNotFind                     = Err{HttpStatus: http.StatusOK, Code: 1902, Message: "UGC Team memberNo is not find ：%s"}
	ErrUgcTeamMemberError                       = Err{HttpStatus: http.StatusOK, Code: 1903, Message: "Your Unauthorized operation memberNo : %s"}
	ErrUgcRecommendNotExist                     = Err{HttpStatus: http.StatusOK, Code: 2152, Message: "Recommend not exist recommendId: %s"}
	ErrUgcRecommendDataExist                    = Err{HttpStatus: http.StatusOK, Code: 2156, Message: "Recommend Mark data exist UGC ID : %s"}
	ErrUgcReportTypeError                       = Err{HttpStatus: http.StatusOK, Code: 2201, Message: "report type parameter error : %s"}
	ErrUgcReportCountCeiling                    = Err{HttpStatus: http.StatusOK, Code: 2203, Message: "Today report count ceiling"}
	ErrUgcReportTodayAlready                    = Err{HttpStatus: http.StatusOK, Code: 2204, Message: "Today already report : %s"}
	ErrUgcRankTemplateStatNameNull              = Err{HttpStatus: http.StatusOK, Code: 2301, Message: "rank template is stats.statName is null"}
	ErrUgcRankTemplateOrderError                = Err{HttpStatus: http.StatusOK, Code: 2302, Message: "rank template is order error : %s"}
	ErrUgcRankTemplateScoreTypeError            = Err{HttpStatus: http.StatusOK, Code: 2303, Message: "rank template is scoreType error : %s"}
	ErrUgcRankTemplateStatValueTypeError        = Err{HttpStatus: http.StatusOK, Code: 2304, Message: "rank template is stats.valueType error : %s"}
	ErrUgcRankTemplateExist                     = Err{HttpStatus: http.StatusOK, Code: 2305, Message: "rank template exist templateId : %s"}
	ErrUgcRankTemplateNotExist                  = Err{HttpStatus: http.StatusOK, Code: 2307, Message: "template not exist templateId : %s"}
	ErrUgcRankTemplateStatPermissionError       = Err{HttpStatus: http.StatusOK, Code: 2308, Message: "rank template is stats.permission error : %s"}
	ErrUgcRankTemplateStatUpdateTypeError       = Err{HttpStatus: http.StatusOK, Code: 2309, Message: "rank template is stats.updateType error : %s"}
	ErrUgcRankTemplateStatSyncError             = Err{HttpStatus: http.StatusOK, Code: 2310, Message: "rank template stats sync error : %s"}
	ErrUgcRankTemplateUnStatSyncError           = Err{HttpStatus: http.StatusOK, Code: 2311, Message: "rank template stats unsync error : %s"}
	ErrUgcRankTemplateParameterAggregationError = Err{HttpStatus: http.StatusOK, Code: 2312, Message: "rank template parameter aggregation error : %s"}
	ErrUgcNotAllowedApplyHomeland               = Err{HttpStatus: http.StatusOK, Code: 2401, Message: "Not allowed to apply"}

	// ErrGraphSuccess Graph 返回码
	ErrGraphSuccess                     = Err{HttpStatus: http.StatusOK, Code: 0, Message: "success"}
	ErrGraphRequestTokenIsNull          = Err{HttpStatus: http.StatusOK, Code: 90000, Message: "request token is null."}
	ErrGraphInvalidAccessToken          = Err{HttpStatus: http.StatusOK, Code: 90001, Message: "Invalid access token."}
	ErrGraphFriendNotExist              = Err{HttpStatus: http.StatusOK, Code: 10001, Message: "friend not exist"}
	ErrGraphMemberNotExist              = Err{HttpStatus: http.StatusOK, Code: 10004, Message: "member not exist"}
	ErrGraphNicknameNotAllowNull        = Err{HttpStatus: http.StatusOK, Code: 6001, Message: "nickname is null"}
	ErrGraphSelfRequest                 = Err{HttpStatus: http.StatusOK, Code: 1104, Message: "Can't request to self"}
	ErrGraphNotExistsFriend             = Err{HttpStatus: http.StatusOK, Code: 1102, Message: "Not exists friend"}
	ErrGraphNotExistsFriendRequest      = Err{HttpStatus: http.StatusOK, Code: 1109, Message: "Not Exists a friend request"}
	ErrGraphOtherTypesFriendsNotAllowed = Err{HttpStatus: http.StatusOK, Code: 6000, Message: "Other types of friends are not allowed"}
	ErrGraphAlreadyBlock                = Err{HttpStatus: http.StatusOK, Code: 1211, Message: "already block."}
	ErrGraphNotExistsBlock              = Err{HttpStatus: http.StatusOK, Code: 1212, Message: "blocked not exists"}
	ErrGraphBlockReceiver               = Err{HttpStatus: http.StatusOK, Code: 1213, Message: "blocked receiver."}
	ErrGraphBlockRequester              = Err{HttpStatus: http.StatusOK, Code: 1214, Message: "blocked requester."}
	ErrGraphCreditRequestBlockError     = Err{HttpStatus: http.StatusOK, Code: 5100, Message: "Requester is credit score below 350 or less."}
	ErrGraphCreditReceiveBlockError     = Err{HttpStatus: http.StatusOK, Code: 5102, Message: "Receiver is credit score below 350 or less."}
	ErrGraphAlreadyFriend               = Err{HttpStatus: http.StatusOK, Code: 1101, Message: "Already friend"}
	ErrGraphAlreadyRequest              = Err{HttpStatus: http.StatusOK, Code: 1106, Message: "Already Requested"}
	ErrGraphAlreadyFriendInBlock        = Err{HttpStatus: http.StatusOK, Code: 1113, Message: "Already friend in the case of block"}
	ErrGraphDailyFriendReqExceed        = Err{HttpStatus: http.StatusOK, Code: 1226, Message: "Exceeded the daily friend request limit."}
	ErrGraphRequesterFriendCntExceed    = Err{HttpStatus: http.StatusOK, Code: 1221, Message: "Requester exceeded friend count."}
	ErrGraphReceiverFriendCntExceed     = Err{HttpStatus: http.StatusOK, Code: 1222, Message: "Receiver exceeded friend count."}
	ErrGraphExistClose                  = Err{HttpStatus: http.StatusOK, Code: 6010, Message: "close status is exist"}
	ErrGraphNotExistClose               = Err{HttpStatus: http.StatusOK, Code: 6011, Message: "close status is not exist"}
	ErrGraphMemberDoesNotExist          = Err{HttpStatus: http.StatusOK, Code: 10004, Message: "Member Does not exist."}
	ErrGraphGameFriendCountDailyLimit   = Err{HttpStatus: http.StatusOK, Code: 1230, Message: "Exceeded the daily limit."}
	ErrGraphGameFriendIntervalLimit     = Err{HttpStatus: http.StatusOK, Code: 1231, Message: "Exceeded the interval limit."}
	ErrGraphGameFriendSocialLimit       = Err{HttpStatus: http.StatusOK, Code: 1232, Message: "Been in social limit."}
	ErrGraphGameFriendAllLimit          = Err{HttpStatus: http.StatusOK, Code: 1233, Message: "Over the limit."}
	ErrGraphMemberNoListEmpty           = Err{HttpStatus: http.StatusOK, Code: 1223, Message: "memberNoList should be not empty."}

	// ErrCsReporterNotFound Cs 返回码
	ErrCsReporterNotFound                       = Err{HttpStatus: http.StatusOK, Code: 100001, Message: "reported id not exist"}
	ErrCsRrplyTemplateNotFound                  = Err{HttpStatus: http.StatusOK, Code: 400010, Message: "Reply template was not found"}
	ErrCsCategoryNotFound                       = Err{HttpStatus: http.StatusOK, Code: 400011, Message: "Category was not found"}
	ErrCsFaqNotFound                            = Err{HttpStatus: http.StatusOK, Code: 400012, Message: "Faq was not found"}
	ErrCsObjectNotFound                         = Err{HttpStatus: http.StatusOK, Code: 400013, Message: "Object was not found"}
	ErrCsParamIsNull                            = Err{HttpStatus: http.StatusOK, Code: 400014, Message: "param is null"}
	ErrCsFileUploadFaild                        = Err{HttpStatus: http.StatusOK, Code: 400020, Message: "file upload faild."}
	ErrCsFileFormatError                        = Err{HttpStatus: http.StatusOK, Code: 400021, Message: "file format error."}
	ErrCsFileParamCanNotEmpty                   = Err{HttpStatus: http.StatusOK, Code: 400022, Message: "file param can't empty"}
	ErrCsReportReasonAndReportTypeAlreadyExists = Err{HttpStatus: http.StatusOK, Code: 600001, Message: "reportReason-reportType already exists"}
	ErrCsNotFoundReportReasonById               = Err{HttpStatus: http.StatusOK, Code: 600002, Message: "reportReason was not found"}
	ErrCsNotFoundReasonId                       = Err{HttpStatus: http.StatusOK, Code: 600003, Message: "reasonId was not found"}
	ErrCsNotFoundReasonType                     = Err{HttpStatus: http.StatusOK, Code: 600004, Message: "reportType was not found"}
	ErrCsGeneralTypeMustBeOne                   = Err{HttpStatus: http.StatusOK, Code: 600005, Message: "the size of reportType which contains GENERAL_REPORT  must be one"}
	ErrCsReportNotFound                         = Err{HttpStatus: http.StatusOK, Code: 600009, Message: "report was not found"}
	ErrCsCompleteTaskFailed                     = Err{HttpStatus: http.StatusOK, Code: 600010, Message: "report complete failed"}
	ErrCsReportFailed                           = Err{HttpStatus: http.StatusOK, Code: 600011, Message: "report failed, please try again later."}
	ErrCsInquiryNotFound                        = Err{HttpStatus: http.StatusOK, Code: 500001, Message: "inquiry was not found"}
	ErrCsInquiryFailed                          = Err{HttpStatus: http.StatusOK, Code: 500002, Message: "inquiry failed, please try again later."}
	ErrCsInquiryCallbackFalied                  = Err{HttpStatus: http.StatusOK, Code: 500003, Message: "inquiry callback failed."}
	ErrCsInquiryAgainFailed                     = Err{HttpStatus: http.StatusOK, Code: 500004, Message: "can't reopen consultations not initiated by oneself."}
	ErrCsInquiryFinishedFalied                  = Err{HttpStatus: http.StatusOK, Code: 500005, Message: "inquiry finished failed, please try again later"}

	// ErrZelosServiceError Zelos 返回码
	ErrZelosServiceError              = Err{HttpStatus: http.StatusOK, Code: 500, Message: "Service error, please contact us"}
	ErrZelosAlreadyExistLeaderboardId = Err{HttpStatus: http.StatusOK, Code: 5001, Message: "already exist leaderboard id"}
	ErrZelosCharacterNoNotFound       = Err{HttpStatus: http.StatusOK, Code: 5002, Message: "character no no found"}
	ErrZelosLeaderboardNotFound       = Err{HttpStatus: http.StatusOK, Code: 5003, Message: "leaderboard not found"}
	ErrZelosBanDeletionUserNotFound   = Err{HttpStatus: http.StatusOK, Code: 5004, Message: "ban deletion user not found"}
	ErrZelosRegistrationNotAllowed    = Err{HttpStatus: http.StatusOK, Code: 5005, Message: "registration of scores is not allowed"}

	// ErrShopService Shop 返回码
	ErrShopService                   = Err{HttpStatus: http.StatusOK, Code: 500, Message: "Service error, please contact us"}
	ErrShopProductNotFound           = Err{HttpStatus: http.StatusOK, Code: 40101, Message: "product was not found"}
	ErrShopProductSnapshotNotFound   = Err{HttpStatus: http.StatusOK, Code: 40103, Message: "product snapshot was not found"}
	ErrShopOrderNotFound             = Err{HttpStatus: http.StatusOK, Code: 40105, Message: "order was not found"}
	ErrShopCashTypeNotFound          = Err{HttpStatus: http.StatusOK, Code: 40106, Message: "cash type not found"}
	ErrShopChannelNotFound           = Err{HttpStatus: http.StatusOK, Code: 40107, Message: "channel could not match"}
	ErrShopUserNotFound              = Err{HttpStatus: http.StatusOK, Code: 40108, Message: "user not found"}
	ErrShopPageDuplicateProduct      = Err{HttpStatus: http.StatusOK, Code: 40109, Message: "Duplicate product information appears in the display page,please contact us"}
	ErrShopLackOfPoint               = Err{HttpStatus: http.StatusOK, Code: 41101, Message: "lack of point"}
	ErrShopLackOfGold                = Err{HttpStatus: http.StatusOK, Code: 41006, Message: "lack of gold"}
	ErrShopLackOfDiamond             = Err{HttpStatus: http.StatusOK, Code: 41007, Message: "lack of start diamond"}
	ErrShopLackOfSCoin               = Err{HttpStatus: http.StatusOK, Code: 41008, Message: "lack of s-coin"}
	ErrShopLackOfNCoin               = Err{HttpStatus: http.StatusOK, Code: 41009, Message: "lack of n-coin"}
	ErrShopInvalidArgument           = Err{HttpStatus: http.StatusOK, Code: 41102, Message: "Invalid argument"}
	ErrShopInvalidThirdOrderNo       = Err{HttpStatus: http.StatusOK, Code: 41104, Message: "third order no already exists"}
	ErrShopAlreadyRollback           = Err{HttpStatus: http.StatusOK, Code: 41105, Message: "order no already rollback"}
	ErrShopNotAllowedBuy             = Err{HttpStatus: http.StatusOK, Code: 50101, Message: "not allowed to buy"}
	ErrShopLimitPurchased            = Err{HttpStatus: http.StatusOK, Code: 50102, Message: "limited purchase %s times, It has been purchased %s times"}
	ErrShopLimitDiscount             = Err{HttpStatus: http.StatusOK, Code: 50103, Message: "limited discount %s times, It has been discounted %s times"}
	ErrShopPresentNotAllowedRollback = Err{HttpStatus: http.StatusOK, Code: 50104, Message: "present cash not allowed rollback"}
	ErrShopTransformNotFound         = Err{HttpStatus: http.StatusOK, Code: 50105, Message: "cash type  %s no exists transform cash config"}
	ErrShopMidasInterface            = Err{HttpStatus: http.StatusOK, Code: 60101, Message: "midas interface error"}
	ErrShopBillInterface             = Err{HttpStatus: http.StatusOK, Code: 60102, Message: "billing interface error"}
	ErrShopGdisInterface             = Err{HttpStatus: http.StatusOK, Code: 60103, Message: "gdis interface error"}
	ErrShopGsInterface               = Err{HttpStatus: http.StatusOK, Code: 60104, Message: "gs interface error"}

	// ErrBillingServiceError Billing 返回码
	ErrBillingServiceError        = Err{HttpStatus: http.StatusInternalServerError, Code: 500, Message: "Service error, please contact us"}
	ErrBillingOrderNotFound       = Err{HttpStatus: http.StatusNotFound, Code: 40105, Message: "Order was not found"}
	ErrBillingCashTypeNotFound    = Err{HttpStatus: http.StatusNotFound, Code: 40106, Message: "Cash type not found"}
	ErrBillingUserNotFound        = Err{HttpStatus: http.StatusNotFound, Code: 40108, Message: "User not found"}
	ErrBillingLackOfCash          = Err{HttpStatus: http.StatusBadRequest, Code: 41101, Message: "lack of Cash"}
	ErrBillingInvalidArgument     = Err{HttpStatus: http.StatusBadRequest, Code: 41102, Message: "Invalid argument"}
	ErrBillingAlreadyRollback     = Err{HttpStatus: http.StatusBadRequest, Code: 41105, Message: "order no already rollback"}
	ErrBillingMidasInterfaceError = Err{HttpStatus: http.StatusBadRequest, Code: 601001, Message: "midas interface return error"}

	// ErrModshopServiceError Modshop返回码
	ErrModshopServiceError                  = Err{HttpStatus: http.StatusInternalServerError, Code: 500, Message: "Service error, please contact us"}
	ErrModshopProductIdAlreadyExist         = Err{HttpStatus: http.StatusBadRequest, Code: 40001, Message: "Product id already exists"}
	ErrModshopNoOperationPermission         = Err{HttpStatus: http.StatusBadRequest, Code: 40002, Message: "No operation permission"}
	ErrModshopLackOfCash                    = Err{HttpStatus: http.StatusBadRequest, Code: 41101, Message: "lack of Cash"}
	ErrModshopProductNotFound               = Err{HttpStatus: http.StatusNotFound, Code: 40401, Message: "Product not found"}
	ErrModshopProductNotOnSale              = Err{HttpStatus: http.StatusNotFound, Code: 40402, Message: "Product not on sale"}
	ErrModshopOrderOperationPermission      = Err{HttpStatus: http.StatusNotFound, Code: 40403, Message: "The order doesn't belong to you"}
	ErrModshopProductNameExistSensitiveInfo = Err{HttpStatus: http.StatusNotFound, Code: 40404, Message: "Product name contains sensitive information"}
	ErrModshopProductIconExistSensitiveInfo = Err{HttpStatus: http.StatusNotFound, Code: 40405, Message: "Product icon contains sensitive information"}
	ErrModshopOrderNoNotFound               = Err{HttpStatus: http.StatusNotFound, Code: 40406, Message: "Order not found"}

	// ErrXpnsRequestTokenIsNull Xpns 返回码
	ErrXpnsRequestTokenIsNull                  = Err{HttpStatus: http.StatusOK, Code: 90000, Message: "Request token is null"}
	ErrXpnsInvalidAccessToken                  = Err{HttpStatus: http.StatusOK, Code: 90001, Message: "Invalid access token"}
	ErrXpnsMemberNotExist                      = Err{HttpStatus: http.StatusOK, Code: 1001, Message: "Member does not exist"}
	ErrXpnsOsTypeError                         = Err{HttpStatus: http.StatusOK, Code: 1002, Message: "Os-Type is blank"}
	ErrXpnsOsTypeIncorrect                     = Err{HttpStatus: http.StatusOK, Code: 1003, Message: "Os-Type is incorrect"}
	ErrXpnsServiceTypeInvalid                  = Err{HttpStatus: http.StatusOK, Code: 1004, Message: "ServiceType is invalid"}
	ErrXpnsIpLimitError                        = Err{HttpStatus: http.StatusOK, Code: 30009, Message: "Send message frequently"}
	ErrXpnsSecretAuthError                     = Err{HttpStatus: http.StatusOK, Code: 31002, Message: "Authorization error"}
	ErrXpnsInvalidParameterError               = Err{HttpStatus: http.StatusOK, Code: 30001, Message: "The parameter error: "}
	ErrXpnsIncorrectPhoneNumber                = Err{HttpStatus: http.StatusOK, Code: 30002, Message: "At least one of the phoneNumbers is incorrect"}
	ErrXpnsLimit4ADay                          = Err{HttpStatus: http.StatusOK, Code: 30003, Message: "Get the verification code more than ten times a day"}
	ErrXpnsLimit4AMinute                       = Err{HttpStatus: http.StatusOK, Code: 30004, Message: "The verification code cannot be obtained again within 60 seconds"}
	ErrXpnsInvalidVerificationCodeOrErrorCode  = Err{HttpStatus: http.StatusOK, Code: 30005, Message: "The verification code is not up to date or invalid"}
	ErrXpnsNotFoundTemplateType                = Err{HttpStatus: http.StatusOK, Code: 30006, Message: "The templateType was not found"}
	ErrXpnsInvalidPhoneNumberSize4Verification = Err{HttpStatus: http.StatusOK, Code: 30007, Message: "The phoneNumbers size must be one for verification"}
	ErrXpnsSceneError                          = Err{HttpStatus: http.StatusOK, Code: 30008, Message: "The scene is incorrect"}
	ErrXpnsExternalServiceError                = Err{HttpStatus: http.StatusOK, Code: 31001, Message: "External service error: "}
	ErrXpnsLimit4Count                         = Err{HttpStatus: http.StatusOK, Code: 31003, Message: "The verification code had sent"}
	ErrXpnsInvalidVerificationCode             = Err{HttpStatus: http.StatusOK, Code: 31004, Message: "The verification code is invalid"}
	ErrXpnsIncorrectVerificationCode           = Err{HttpStatus: http.StatusOK, Code: 31005, Message: "The verification code is incorrect"}
	ErrXpnsVerifyCodeLimit                     = Err{HttpStatus: http.StatusOK, Code: 31006, Message: "Verified code limit"}
	ErrXpnsIncorrectEmailAddress               = Err{HttpStatus: http.StatusOK, Code: 31007, Message: "Email address format is incorrect"}

	// ErrFileNeedLeastOneFileId File返回码
	ErrFileNeedLeastOneFileId = Err{HttpStatus: http.StatusOK, Code: 41001, Message: "need least one fileId"}
	ErrFileFileCantBeEmpty    = Err{HttpStatus: http.StatusOK, Code: 41002, Message: "file can't be empty"}
	ErrFileFixedType          = Err{HttpStatus: http.StatusOK, Code: 41003, Message: "type must be [DOWNLOAD|UPLOAD]"}
	ErrFileMetaNotFound       = Err{HttpStatus: http.StatusOK, Code: 40201, Message: "meta was not found"}
	ErrFileFileNotFound       = Err{HttpStatus: http.StatusOK, Code: 40202, Message: "file was not found"}
	ErrFileIllegalImg         = Err{HttpStatus: http.StatusOK, Code: 41009, Message: "Illegal picture"}

	// ErrIgcfFail Igcf 返回码
	ErrIgcfFail                                   = Err{HttpStatus: http.StatusOK, Code: -1, Message: "broker fail"}
	ErrIgcfAccessTokenFail                        = Err{HttpStatus: http.StatusOK, Code: -2, Message: "access token fail"}
	ErrIgcfMessengerApiFail                       = Err{HttpStatus: http.StatusOK, Code: -3, Message: "messenger api fail"}
	ErrIgcfMessengerApiSenderNotFound             = Err{HttpStatus: http.StatusOK, Code: -4, Message: "sender not found"}
	ErrIgcfContainingInappropriateMessage         = Err{HttpStatus: http.StatusOK, Code: -5, Message: "containing inappropriate message"}
	ErrIgcfServiceInfoInsertFail                  = Err{HttpStatus: http.StatusOK, Code: -6, Message: "db insert fail"}
	ErrIgcfBannedUser                             = Err{HttpStatus: http.StatusOK, Code: -7, Message: "messenger banned user"}
	ErrIgcfMessageNotExist                        = Err{HttpStatus: http.StatusOK, Code: -8, Message: "message not exist"}
	ErrIgcfReceiverWasBlocked                     = Err{HttpStatus: http.StatusOK, Code: -9, Message: "receiver was blocked"}
	ErrIgcfSocialGraphApiFail                     = Err{HttpStatus: http.StatusOK, Code: -10, Message: "social graph api fail"}
	ErrIgcfMessengerApiReceiverNotFound           = Err{HttpStatus: http.StatusOK, Code: -11, Message: "receiver not found"}
	ErrIgcfMessageFilterApiFail                   = Err{HttpStatus: http.StatusOK, Code: -12, Message: "message filter api fail"}
	ErrIgcfMessengerApiInvalidRequestParam        = Err{HttpStatus: http.StatusOK, Code: -13, Message: "messenger api invalid request param"}
	ErrIgcfBannedUserApiFail                      = Err{HttpStatus: http.StatusOK, Code: -14, Message: "banned user api fail"}
	ErrIgcfMessengerApiFromJidNotAttendee         = Err{HttpStatus: http.StatusOK, Code: -15, Message: "from jid not attendee"}
	ErrIgcfMessengerApiRequestTooLarge            = Err{HttpStatus: http.StatusOK, Code: -16, Message: "request too large"}
	ErrIgcfLowCreditCutoff                        = Err{HttpStatus: http.StatusOK, Code: -17, Message: "low credit cutoff"}
	ErrIgcfLowCredit10timesWithin24h              = Err{HttpStatus: http.StatusOK, Code: -18, Message: "low credit 10times_within_24hour"}
	ErrIgcfLowCreditOnceWithin10m                 = Err{HttpStatus: http.StatusOK, Code: -19, Message: "low credit once_within_10minute"}
	ErrIgcfSenderWasBlocked                       = Err{HttpStatus: http.StatusOK, Code: -20, Message: "sender was blocked"}
	ErrIgcfContainingInappropriateAsteriskMessage = Err{HttpStatus: http.StatusOK, Code: -21, Message: "containing inappropriate asterisk message"}
	ErrIgcfBannedUserVoice                        = Err{HttpStatus: http.StatusOK, Code: -22, Message: "messenger banned voice"}
	ErrIgcfBannedUserChat                         = Err{HttpStatus: http.StatusOK, Code: -23, Message: "messenger banned chat"}

	// ErrGsSuccess Gs 返回码
	ErrGsSuccess                 = Err{HttpStatus: http.StatusOK, Code: 0, Message: "Success"}
	ErrGsServiceError            = Err{HttpStatus: http.StatusOK, Code: -1, Message: "Service Error"}
	ErrGsRateLimit               = Err{HttpStatus: http.StatusOK, Code: 10202, Message: "Operation Too Frequently"}
	ErrGsCompanyNotFound         = Err{HttpStatus: http.StatusOK, Code: 401001, Message: "company was not found"}
	ErrGsConfigNotFound          = Err{HttpStatus: http.StatusOK, Code: 401002, Message: "the game configuration was not found"}
	ErrGsEnumNotFound            = Err{HttpStatus: http.StatusOK, Code: 401003, Message: "%s was not found"}
	ErrGsGradeNotFound           = Err{HttpStatus: http.StatusOK, Code: 401003, Message: "grade was not found"}
	ErrGsStatusNotFound          = Err{HttpStatus: http.StatusOK, Code: 401003, Message: "status was not found"}
	ErrGsPackageConfigNotFound   = Err{HttpStatus: http.StatusOK, Code: 401004, Message: "package configuration was not found"}
	ErrGsPackageNotFound         = Err{HttpStatus: http.StatusOK, Code: 401005, Message: "package not found"}
	ErrGsGameConfigFieldNotFound = Err{HttpStatus: http.StatusOK, Code: 401006, Message: "the game config field was not found"}
	ErrGsInvalidGameName         = Err{HttpStatus: http.StatusOK, Code: 41101, Message: "game name already exists"}
	ErrGsInvalidGameId           = Err{HttpStatus: http.StatusOK, Code: 41102, Message: "game id already exists"}
	ErrGsInvalidGameNo           = Err{HttpStatus: http.StatusOK, Code: 41103, Message: "game no already exists, please exchange game id"}
	ErrGsInvalidGameIdRuleInd    = Err{HttpStatus: http.StatusOK, Code: 41104, Message: "game id contains only lowercase letters and digits"}
	ErrGsInvalidGameIdRuleAgent  = Err{HttpStatus: http.StatusOK, Code: 41105, Message: "game id contains only uppercase and lowercase letters and digits"}
	ErrGsInvalidGameIdRuleOther  = Err{HttpStatus: http.StatusOK, Code: 41106, Message: "game id contains only uppercase and lowercase letters digits and special characters -_"}
	ErrGsGameConfigExists        = Err{HttpStatus: http.StatusOK, Code: 41107, Message: "the game configuration already exists"}
	ErrGsGameIdIsValid           = Err{HttpStatus: http.StatusOK, Code: 41108, Message: "the game id not conform to the regulation"}
	ErrGsPackageDefaultNotExists = Err{HttpStatus: http.StatusOK, Code: 41207, Message: "the package default web url not exists"}

	ErrActivityNotFound      = Err{HttpStatus: http.StatusNotFound, Code: 404401, Message: "activity not found"}
	ErrMissionNotFound       = Err{HttpStatus: http.StatusNotFound, Code: 404402, Message: "mission not found"}
	ErrMissionNotComplete    = Err{HttpStatus: http.StatusForbidden, Code: 403401, Message: "mission not completed"}
	ErrRewardAlreadyObtained = Err{HttpStatus: http.StatusForbidden, Code: 403402, Message: "reward has been obtained"}
)

type Error interface {
	error
	GetHttpStatus() int
	GetCode() int32
	GetMessage() string
	WithCode(code int32) Error
	WithMsg(msg string) Error
	WithHttpStatus(httpStatus int) Error
}

type Err struct {
	Code       int32  `json:"code"`
	Message    string `json:"message,omitempty"`
	HttpStatus int    `json:"-"`
}

func NewErr(code int32, message string) *Err {
	return &Err{Code: code, Message: message}
}

func (e Err) GetCode() int32 {
	return e.Code
}

func (e Err) GetMessage() string {
	return e.Message
}

func (e Err) WithCode(code int32) Error {
	e.Code = code
	return e
}

func (e Err) WithMsg(msg string) Error {
	e.Message = msg
	return e
}

func (e Err) Error() string {
	raw, _ := json.Marshal(e)
	return string(raw)
}

func (e Err) GetHttpStatus() int {
	if e.HttpStatus == 0 {
		return http.StatusOK
	}
	return e.HttpStatus
}

func (e Err) WithHttpStatus(httpStatus int) Error {
	e.HttpStatus = httpStatus
	return e
}
