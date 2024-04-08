package auth

import (
	"context"
	"github.com/cstcen/gore/auth"
	"github.com/cstcen/gore/common"
	"github.com/cstcen/gore/gonfig"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"strings"
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
	jwt.RegisteredClaims
	// MemberNo 星空屋用户唯一标识符
	MemberNo int64 `json:"men,omitempty"`
	// CharacterNo 游戏角色ID（目前用于手工星球，因为它是单一角色）
	CharacterNo int64 `json:"chn,omitempty"`
	// ApplicationNo 应用编号
	ApplicationNo int `json:"apn,omitempty"`
	// GameId 游戏ID，例如：SGXQ
	GameId string `json:"gid,omitempty"`
	// OpenId 星空屋OpenId
	OpenId string `json:"opi,omitempty"`
	// Nickname 星空屋昵称
	Nickname string `json:"nic,omitempty"`
	// ProfileImg 星空屋头像
	ProfileImg string `json:"pri,omitempty"`
	// LoginType 登录类型
	LoginType common.LoginType `json:"lot,omitempty"`
	// LoginValue 登录值
	LoginValue string `json:"lov,omitempty"`
	// FirstLogin 是否第一次登录
	FirstLogin bool `json:"fil"`
	// ChannelId 登录的星空屋渠道号
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

func NewClaimsFromV1Member(m *auth.Member) *MemberClaims {
	claims := MemberClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "",
			Subject:   strconv.Itoa(m.MemberNo),
			Audience:  []string{m.SvrId},
			ExpiresAt: &jwt.NumericDate{Time: time.UnixMilli(int64(m.ExpireTime))},
			NotBefore: nil,
			IssuedAt:  &jwt.NumericDate{Time: time.UnixMilli(int64(m.Timestamp))},
			ID:        m.SvrId,
		},
		MemberNo:           int64(m.MemberNo),
		CharacterNo:        int64(m.CharacterNo),
		ApplicationNo:      m.ApplicationNo,
		GameId:             "SGXQ",
		OpenId:             "",
		Nickname:           m.Nickname,
		ProfileImg:         m.ProfileImg,
		LoginType:          0,
		LoginValue:         m.OpenId,
		FirstLogin:         false,
		ChannelId:          "",
		RegisteredDateTime: time.Time{},
		LastLaunchTime:     m.LastLaunchTime,
		DeviceId:           m.OsId,
		DeviceOs:           0,
		Environment:        gonfig.Instance().GetString("env"),
		Simulator:          false,
	}
	if strings.EqualFold(m.LoginType, "QQ") {
		claims.LoginType = common.LoginTypeSgxqMsdkQq
	} else if strings.EqualFold(m.LoginType, "WECHAT") {
		claims.LoginType = common.LoginTypeSgxqMsdkWx
	}
	if strings.EqualFold(m.ProviderOS, "IOS") {
		claims.DeviceOs = common.ClientTypeIos
	} else if strings.EqualFold(m.ProviderOS, "ANDROID") {
		claims.DeviceOs = common.ClientTypeAndroid
	} else if strings.EqualFold(m.ProviderOS, "PC") {
		claims.DeviceOs = common.ClientTypePc
	}
	return &claims
}

func (c *MemberClaims) CompatibleV1Member() *auth.Member {
	m := auth.Member{}
	if c.MemberNo == 0 {
		m.Agent = "XK5_SERVER"
	} else {
		m.Agent = "USER"
		m.MemberNo = int(c.MemberNo)
		m.Nickname = c.Nickname
		m.CharacterNo = int(c.CharacterNo)
	}
	m.ApplicationNo = c.ApplicationNo
	m.ExpireTime = int(c.ExpiresAt.UnixMilli())

	if c.LoginType.IsWx() {
		m.LoginType = "WECHAT"
		m.OpenId = c.LoginValue
	} else if c.LoginType.IsQq() {
		m.LoginType = "QQ"
		m.OpenId = c.LoginValue
	}
	m.OsId = c.DeviceId
	m.ProfileImg = c.ProfileImg
	if c.DeviceOs == common.ClientTypeIos {
		m.ProviderOS = "iOS"
	} else if c.DeviceOs == common.ClientTypeAndroid {
		m.ProviderOS = "AOS"
	}
	if c.IssuedAt != nil {
		m.Timestamp = int(c.IssuedAt.UnixMilli())
	}
	return &m
}
