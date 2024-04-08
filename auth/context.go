package auth

import (
	"git.tenvine.cn/backend/gore/common"
	"github.com/gin-gonic/gin"
	"time"
)

type Context struct {
	*gin.Context
}

func NewContext(c *gin.Context) *Context {
	return &Context{Context: c}
}
func (c *Context) Claims() *MemberClaims {
	return GetClaimsFromGinContext(c.Context)
}
func (c *Context) MemberNo() int64 {
	return c.Claims().MemberNo
}
func (c *Context) ApplicationNo() int {
	return c.Claims().ApplicationNo
}
func (c *Context) OpenId() string {
	return c.Claims().OpenId
}
func (c *Context) Nickname() string {
	return c.Claims().Nickname
}
func (c *Context) ProfileImg() string {
	return c.Claims().ProfileImg
}
func (c *Context) LoginType() common.LoginType {
	return c.Claims().LoginType
}
func (c *Context) LoginValue() string {
	return c.Claims().LoginValue
}
func (c *Context) IsFirstLogin() bool {
	return c.Claims().IsFirstLogin
}
func (c *Context) ChannelId() string {
	return c.Claims().ChannelId
}
func (c *Context) RegisteredDate() *time.Time {
	return c.Claims().RegisteredDate
}
func (c *Context) LastPrivilegeLaunchDate() string {
	return c.Claims().LastPrivilegeLaunchDate
}
func (c *Context) DeviceId() string {
	return c.Claims().DeviceId
}
func (c *Context) ClientType() common.ClientType {
	return c.Claims().ClientType
}
func (c *Context) Environment() string {
	return c.Claims().Environment
}
func (c *Context) Simulator() bool {
	return c.Claims().Simulator
}
