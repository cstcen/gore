package auth

import (
	"context"
	"errors"
	"fmt"
	"git.tenvine.cn/backend/gore/common"
	"git.tenvine.cn/backend/gore/gonfig"
	goreHttp "git.tenvine.cn/backend/gore/http"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func DecryptToken(ctx context.Context, token string) (*string, error) {
	url := gonfig.Instance().GetString("xk5.auth.userJwt")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)
	var result common.DataResult[struct {
		UserJwt string `json:"user_jwt"`
	}]
	if err := goreHttp.Do(req, &result); err != nil {
		return nil, err
	}
	if result.GetCode() != common.BaseResultSuccess.GetCode() || len(result.Data.UserJwt) == 0 {
		return nil, errors.New("token verification failed")
	}

	return &result.Data.UserJwt, nil
}

type MemberClaims struct {
	Member
	jwt.RegisteredClaims
}

func ParseJwt(jwtStr string) (*MemberClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(jwtStr, &MemberClaims{})
	if err != nil {
		return nil, fmt.Errorf("ParseUnverified: %w", err)
	}

	claims, ok := token.Claims.(*MemberClaims)
	if !ok {
		return nil, fmt.Errorf("claims error: %w", err)
	}
	return claims, nil
}
