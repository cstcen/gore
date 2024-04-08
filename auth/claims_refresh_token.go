package auth

import (
	"git.tenvine.cn/backend/gore/common"
	"strconv"
	"time"
)

const (
	RefreshTokenExpDuration = time.Hour * 24 * 30
)

func NewUserRefreshTokenClaims(memberNo int64, appId int32, deviceId string, clientId string, clientType common.ClientType, goreAuthMember *Member) *MemberClaims {
	jti := NewUserTokenJti(JtiTypeRefreshToken, memberNo, appId, clientType)
	sub := strconv.FormatInt(memberNo, 10)
	var aud = []string{clientId, string(AudienceUser)}
	return NewRefreshTokenClaims(sub, aud, jti, deviceId, clientType, goreAuthMember)
}

func NewRefreshTokenClaims(sub string, aud []string, jti string, did string, clt common.ClientType, memberClaims *Member) *MemberClaims {
	return NewToken(sub, aud, jti, did, clt, memberClaims, RefreshTokenExpDuration)
}
