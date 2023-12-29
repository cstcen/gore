package auth

import (
	"git.tenvine.cn/backend/gore/common"
	"github.com/golang-jwt/jwt/v5"
	"slices"
	"strings"
	"time"
)

type MemberClaims struct {
	Member
	jwt.RegisteredClaims
	DeviceId    string            `json:"did,omitempty"`
	ClientType  common.ClientType `json:"clt,omitempty"`
	Environment string            `json:"env,omitempty"`
	Simulator   bool              `json:"simulator"`
}

func (c *MemberClaims) IsAccessToken() bool {
	return strings.Contains(c.ID, string(JtiTypeAccessToken))
}

func (c *MemberClaims) IsRefreshToken() bool {
	return strings.Contains(c.ID, string(JtiTypeRefreshToken)) && c.IsUser()
}

func (c *MemberClaims) IsSvr() bool {
	return slices.Contains(c.Audience, string(AudienceSvr))
}

func (c *MemberClaims) IsUser() bool {
	return slices.Contains(c.Audience, string(AudienceUser))
}

func (c *MemberClaims) ExpiresIn() int64 {
	expiresAt := c.ExpiresAt
	if expiresAt == nil {
		return 0
	}
	return expiresAt.Unix() - time.Now().Unix()
}

type Member struct {
	// MemberNo 星空屋用户唯一标识符
	MemberNo int64 `json:"men,omitempty"`
	// ApplicationNo 应用编号
	ApplicationNo int `json:"apn,omitempty"`
	// OpenId 星空屋OpenId
	OpenId string `json:"opi,omitempty"`
	// Nickname 星空屋昵称
	Nickname string `json:"nic,omitempty"`
	// ProfileImg 星空屋头像
	ProfileImg string `json:"pri,omitempty"`
	// LoginType 本次登录类型
	LoginType common.LoginType `json:"lot,omitempty"`
	// LoginValue 本次登录值
	LoginValue string `json:"lov,omitempty"`
	// IsFirstLogin 是否时第一次登录
	IsFirstLogin bool `json:"ifl,omitempty"`
	// ChannelId 本次登录的星空屋渠道号
	ChannelId string `json:"chi,omitempty"`
	// RegisteredDate 账号注册时间
	RegisteredDate *time.Time `json:"red,omitempty"`
	// LastPrivilegeLaunchDate 最后一次特权登录时间
	LastPrivilegeLaunchDate string `json:"lpld,omitempty"`
}
