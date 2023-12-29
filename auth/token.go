package auth

import (
	"fmt"
	"git.tenvine.cn/backend/gore/common"
	"git.tenvine.cn/backend/gore/gonfig"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

const (
	PrefixBearerAuth = "Bearer "
)

func IsBearerAuth(req *http.Request) bool {
	return strings.HasPrefix(req.Header.Get("Authorization"), PrefixBearerAuth)
}

func GetSgxqSvrTokenString() (string, error) {
	return GetSvrTokenString(10001, gonfig.Instance().GetString("name"))
}
func GetSvrTokenString(appId int32, serviceName string) (string, error) {
	return GenerateJwt(NewSvrAccessTokenClaims(appId, serviceName), []byte(serviceName))
}

func ParseJwt(jwtStr string) (*MemberClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(strings.TrimPrefix(jwtStr, PrefixBearerAuth), &MemberClaims{})
	if err != nil {
		return nil, fmt.Errorf("ParseUnverified: %w", err)
	}

	claims, ok := token.Claims.(*MemberClaims)
	if !ok {
		return nil, fmt.Errorf("claims error: %w", err)
	}
	return claims, nil
}

func GenerateJwt(claims *MemberClaims, key []byte) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
}

func NewUserTokenJti(jtiType JtiType, memberNo int64, appId int32, clientType common.ClientType) string {
	cliType := ConvertClientTypeToString(clientType)
	return fmt.Sprintf("%s:auth:v4:%s:member:%v:app:%v:client:%s", gonfig.Instance().GetString("env"), jtiType, memberNo, appId, cliType)
}

func NewSvrTokenJti(appId int32, clientType common.ClientType) string {
	cliType := ConvertClientTypeToString(clientType)
	return fmt.Sprintf("%s:auth:v4:%s:app:%v:client:%s", gonfig.Instance().GetString("env"), JtiTypeAccessToken, appId, cliType)
}

func ConvertClientTypeToString(clientType common.ClientType) string {
	switch clientType {
	default:
		return "web"
	case common.ClientTypeAndroid, common.ClientTypeIos:
		return "mobile"
	case common.ClientTypePc:
		return "pc"
	}
}

func NewToken(sub string, aud []string, jti string, did string, clt common.ClientType, goreAuthMember *Member, ttl time.Duration) *MemberClaims {
	now := time.Now()
	exp := now.Add(ttl)
	xk5Host := "xk5.com"
	iss := fmt.Sprintf("%s/auth/v4.0", xk5Host)

	memberClaims := MemberClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    iss,
			Subject:   sub,
			Audience:  aud,
			ExpiresAt: jwt.NewNumericDate(exp),
			ID:        jti,
		},
		DeviceId:    did,
		ClientType:  clt,
		Environment: gonfig.Instance().GetString("env"),
	}
	if goreAuthMember != nil {
		memberClaims.Member = *goreAuthMember
	}

	return &memberClaims
}
