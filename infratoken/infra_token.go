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
	"github.com/sirupsen/logrus"
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

	ctxLog := log.WithContext(c)

	var response *Response
	var err error

	if goreCache.Instance() != nil {
		response, err = get(c, cacheKey)
		if response != nil {
			return response.AccessToken, nil
		}
	}

	response, err = request()
	if err != nil {
		ctxLog.WithError(err).Warn("get infra token fail")
		return "", err
	}

	// async
	go save(response, cacheKey, goreCache.Instance(), ctxLog)

	return response.AccessToken, nil
}

func request() (*Response, error) {

	url := gonfig.Instance().GetString("tenvine.infraToken.url")
	req := Request{
		GrantType:    gonfig.Instance().GetString("tenvine.infraToken.grantType"),
		ClientId:     gonfig.Instance().GetString("tenvine.infraToken.clientId"),
		ClientSecret: gonfig.Instance().GetString("tenvine.infraToken.clientSecret"),
		MacAddress:   macAddr,
	}

	result := new(Response)
	if err := goreHttp.Post(url, constant.ContentTypeApplicationJSON, req, result); err != nil {
		return nil, err
	}

	return result, nil
}

func get(c context.Context, key string) (*Response, error) {
	var wanted Response
	ctx, cancelFunc := context.WithTimeout(c, constant.TimeoutConn)
	if err := goreCache.Instance().Get(ctx, key, &wanted); err != nil {
		return nil, err
	}
	defer cancelFunc()
	return &wanted, nil

}

func save(response *Response, key string, cc *cache.Cache, ctxLog *logrus.Entry) {
	if cc == nil {
		return
	}
	if response == nil {
		return
	}

	ttl := time.Duration(response.ExpiresIn)
	ctx, cancelFunc := context.WithTimeout(context.Background(), constant.TimeoutConn)
	defer cancelFunc()
	if err := goreCache.Instance().Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: response,
		TTL:   ttl,
	}); err != nil {
		ctxLog.WithError(err).Warn("cache infra token fail")
	}

}
