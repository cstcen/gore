package auth

import (
	"fmt"
	"git.tenvine.cn/backend/gore/gonfig"
)

func cacheKey(token string) string {
	return fmt.Sprintf("%s:%s", gonfig.Instance().GetString("env"), token)
}
