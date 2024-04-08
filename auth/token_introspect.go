package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"git.tenvine.cn/backend/gore/auth/v2/api"
	"git.tenvine.cn/backend/gore/common"
	"git.tenvine.cn/backend/gore/gonfig"
	goreHttp "git.tenvine.cn/backend/gore/http"
	"strings"
	"time"
)

func CheckToken(ctx context.Context, authorization string, options ...any) (*MemberClaims, error) {

	if len(strings.Split(authorization, ".")) == 3 {
		return ParseJwt(authorization)
	}

	return TokenIntrospect(ctx, authorization, options)
}

func TokenIntrospect(ctx context.Context, authorization string, options []any) (*MemberClaims, error) {
	authorization = strings.TrimPrefix(authorization, PrefixBearerAuth)
	clientWithResponses, err := api.NewClientWithResponses(gonfig.Instance().GetString("xk5.host.external"), func(client *api.Client) error {
		client.Client = goreHttp.Instance()
		return nil
	})
	if err != nil {
		return nil, err
	}

	ctx, cancelFunc := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancelFunc()
	clientId := gonfig.Instance().GetString("xk5.auth.client_id")
	clientSecret := gonfig.Instance().GetString("xk5.auth.client_secret")
	for _, option := range options {
		switch x := option.(type) {
		case ClientIdGetting:
			clientId = x.GetClientId()
		case ClientSecretGetting:
			clientSecret = x.GetClientSecret()
		}
	}
	reqBody := api.PostAuthV40Oauth2IntrospectJSONRequestBody{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Token:        authorization,
	}

	response, err := clientWithResponses.PostAuthV40Oauth2IntrospectWithResponse(ctx, reqBody)
	if err != nil {
		return nil, err
	}
	if response.JSON200 == nil {
		return nil, common.ErrUnauthenticated
	}
	var claims MemberClaims
	if err := json.NewDecoder(bytes.NewReader(response.Body)).Decode(&claims); err != nil {
		return nil, err
	}

	return &claims, nil
}

type ClientIdGetting interface {
	GetClientId() string
}

type ClientIdSettingFn func() string

func (fn ClientIdSettingFn) GetClientId() string {
	return fn()
}

type ClientSecretGetting interface {
	GetClientSecret() string
}

type ClientSecretGettingFn func() string

func (fn ClientSecretGettingFn) GetClientSecret() string {
	return fn()
}
