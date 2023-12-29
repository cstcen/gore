package auth

import (
	"git.tenvine.cn/backend/gore/common"
	"strconv"
	"strings"
	"time"
)

const (
	AccessTokenExpDuration = time.Hour * 6
)

func RefreshAccessTokenClaims(refreshToken *MemberClaims, deviceId string) *MemberClaims {
	sub := refreshToken.Subject
	aud := refreshToken.Audience
	jti := strings.ReplaceAll(refreshToken.ID, "refresh_token", "access_token")
	did := deviceId
	return NewAccessTokenClaims(sub, aud, jti, did, refreshToken.ClientType, &refreshToken.Member)
}

func NewUserAccessTokenClaims(memberNo int64, appId int32, deviceId string, clientId string, clientType common.ClientType, goreAuthMember *Member) *MemberClaims {
	jti := NewUserTokenJti(JtiTypeAccessToken, memberNo, appId, clientType)
	sub := strconv.FormatInt(memberNo, 10)
	var aud = []string{clientId, string(AudienceUser)}
	return NewAccessTokenClaims(sub, aud, jti, deviceId, clientType, goreAuthMember)
}

func NewSvrAccessTokenClaims(appId int32, clientId string) *MemberClaims {
	clientType := common.ClientTypeWeb
	jti := NewSvrTokenJti(appId, clientType)
	appIdString := strconv.FormatInt(int64(appId), 10)
	sub := appIdString
	var aud = []string{clientId, string(AudienceSvr)}
	return NewAccessTokenClaims(sub, aud, jti, "", clientType, nil)
}

func NewAccessTokenClaims(sub string, aud []string, jti string, did string, clt common.ClientType, memberClaims *Member) *MemberClaims {
	return NewToken(sub, aud, jti, did, clt, memberClaims, AccessTokenExpDuration)
}
