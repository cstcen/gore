package infratoken

import (
	"context"
	"git.tenvine.cn/backend/gore/constant"
	goreCache "git.tenvine.cn/backend/gore/db/cache"
	"git.tenvine.cn/backend/gore/gonfig"
	goreHttp "git.tenvine.cn/backend/gore/http"
	"git.tenvine.cn/backend/gore/log"
	"git.tenvine.cn/backend/gore/util"
	"github.com/go-redis/cache/v8"
	"time"
)

var (
	macAddr = util.GetMACAddr()
)

type Request struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	MacAddress   string `json:"macaddress"`
}

type Response struct {
	ReturnCode    uint   `json:"return_code"`
	ReturnMessage string `json:"return_message,omitempty"`
	ExpiresIn     uint   `json:"expires_in,omitempty"`
	ExpiresDt     uint   `json:"expires_dt,omitempty"`
	AccessToken   string `json:"access_token,omitempty"`
	TokenType     string `json:"token_type,omitempty"`
}

func Get(c context.Context) (string, error) {
	cacheKey := gonfig.Instance().GetString("tenvine.keyPrefix") + gonfig.Instance().GetString("tenvine.infraToken.clientId")

	var response *Response
	var err error

	if goreCache.Instance() != nil {
		response, err = get(c, cacheKey)
		if response != nil {
			return response.AccessToken, nil
		}
	}

	response, err = request(c)
	if err != nil {
		log.ErrorCf(c, "get infra token fail: %v", err)
		return "", err
	}

	// async
	go save(response, cacheKey, goreCache.Instance(), c)

	return response.AccessToken, nil
}

func request(c context.Context) (*Response, error) {

	url := gonfig.Instance().GetString("tenvine.infraToken.url")
	req := Request{
		GrantType:    gonfig.Instance().GetString("tenvine.infraToken.grantType"),
		ClientId:     gonfig.Instance().GetString("tenvine.infraToken.clientId"),
		ClientSecret: gonfig.Instance().GetString("tenvine.infraToken.clientSecret"),
		MacAddress:   macAddr,
	}

	result := new(Response)
	if err := goreHttp.Post(c, url, constant.ContentTypeApplicationJSON, req, result); err != nil {
		return nil, err
	}

	return result, nil
}

func get(c context.Context, key string) (*Response, error) {
	var wanted Response
	ctx, cancelFunc := context.WithTimeout(c, constant.TimeoutConn)
	defer cancelFunc()
	if err := goreCache.Instance().Get(ctx, key, &wanted); err != nil {
		return nil, err
	}
	defer cancelFunc()
	return &wanted, nil

}

func save(response *Response, key string, cc *cache.Cache, ctx context.Context) {
	if cc == nil || response == nil || response.ExpiresIn == 0 {
		return
	}

	ttl := time.Duration(response.ExpiresIn) * time.Millisecond
	c, cancelFunc := context.WithTimeout(context.Background(), constant.TimeoutConn)
	defer cancelFunc()
	if err := goreCache.Instance().Set(&cache.Item{
		Ctx:   c,
		Key:   key,
		Value: response,
		TTL:   ttl,
	}); err != nil {
		log.WarningCf(ctx, "cache infra token fail: %v", err)
	}

}
