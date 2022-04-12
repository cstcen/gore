package auth

import (
	"context"
	"git.tenvine.cn/backend/gore/gonfig"
)

func ExternalCheck(ctx context.Context, token string) (*Member, error) {
	return Check(ctx, token, gonfig.Instance().GetString("tenvine.api.host")+gonfig.Instance().GetString("tenvine.api.verifyUserToken"))
}
