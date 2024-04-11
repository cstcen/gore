package auth

import (
	"context"
	"github.com/cstcen/gore/common"
	"github.com/gin-gonic/gin"
	"github.com/go-jose/go-jose/v3/jwt"
	"net/http"
	"time"
)

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

const (
	ContextKeyJwtPayload = "JwtPayload"
)

func SetClaimsToContext(req *http.Request, claims *MemberClaims) *http.Request {
	return req.WithContext(context.WithValue(req.Context(), ContextKeyJwtPayload, claims))
}

func GetClaimsFromContext(c context.Context) *MemberClaims {
	claims, _ := c.Value(ContextKeyJwtPayload).(*MemberClaims)
	return claims
}

type MemberClaims struct {
	jwt.Claims
	// MemberNo 平台用户唯一标识符
	MemberNo int64 `json:"men,omitempty"`
	// CharacterNo 游戏角色ID（目前用于手工星球，因为它是单一角色）
	CharacterNo int64 `json:"chn,omitempty"`
	// ApplicationNo 应用编号
	ApplicationNo int `json:"apn,omitempty"`
	// GameId 游戏ID，例如：SGXQ
	GameId string `json:"gid,omitempty"`
	// OpenId 平台OpenId
	OpenId string `json:"opi,omitempty"`
	// Nickname 平台昵称
	Nickname string `json:"nic,omitempty"`
	// ProfileImg 平台头像
	ProfileImg string `json:"pri,omitempty"`
	// LoginType 登录类型
	LoginType common.LoginType `json:"lot,omitempty"`
	// LoginValue 登录值
	LoginValue string `json:"lov,omitempty"`
	// FirstLogin 是否第一次登录
	FirstLogin bool `json:"fil"`
	// ChannelId 登录的平台渠道号
	ChannelId string `json:"chi,omitempty"`
	// RegisteredDateTime 账号注册时间
	RegisteredDateTime time.Time `json:"rdt,omitempty"`
	// LastLaunchTime 最后一次特权登录时间
	LastLaunchTime string `json:"llt,omitempty"`
	// DeviceId 登录的设备ID
	DeviceId string `json:"did,omitempty"`
	// DeviceId 登录的设备操作系统
	DeviceOs common.ClientType `json:"dos,omitempty"`
	// Environment 环境变量
	Environment string `json:"env,omitempty"`
	// Simulator 是否是模拟器登录
	Simulator bool `json:"sim"`
}
