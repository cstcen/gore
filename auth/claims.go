package auth

import "github.com/gin-gonic/gin"

const (
	GinContextKeyJwtPayload = "JwtPayload"
)

func SetClaimsToGinContext(c *gin.Context, claims *MemberClaims) {
	c.Set(GinContextKeyJwtPayload, claims)
}

func GetClaimsFromGinContext(c *gin.Context) *MemberClaims {
	v, exists := c.Get(GinContextKeyJwtPayload)
	if !exists {
		return nil
	}
	claims, ok := v.(*MemberClaims)
	if !ok {
		return nil
	}
	return claims
}

type JtiType string

const (
	JtiTypeAccessToken  JtiType = "access_token"
	JtiTypeRefreshToken JtiType = "refresh_token"
)

type Audience string

const (
	AudienceUser Audience = "user"
	AudienceSvr  Audience = "svr"
)
