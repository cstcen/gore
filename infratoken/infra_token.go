package infratoken

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"git.tenvine.cn/backend/gore/constant"
	"git.tenvine.cn/backend/gore/log"
	"git.tenvine.cn/backend/gore/util"
	"github.com/go-redis/cache/v8"
	"io"
	"net/http"
	"time"
)

const (
	GrantType      = "client_credentials"
	AuthID         = "infra_billing_server"
	AuthSecret     = "alkjsdf8jsf9n3onf78s9dhfjlk398f9hlksdfuihaoisdhf"
	AuthHostFormat = "https://m-apis-%s.xk5.com/auth/v2/infra_server/init"
)

var (
	req *Request
)

type Request struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	MacAddress   string `json:"mac_address"`
}

type Response struct {
	ReturnCode    uint   `json:"return_code"`
	ReturnMessage string `json:"return_message,omitempty"`
	ExpiresIn     uint   `json:"expires_in,omitempty"`
	AccessToken   string `json:"access_token,omitempty"`
}

func init() {
	req = &Request{
		GrantType:    GrantType,
		ClientId:     AuthID,
		ClientSecret: AuthSecret,
		MacAddress:   util.GetMACAddr(),
	}
}

func Get(c context.Context, appName, env string, cc *cache.Cache) (string, error) {
	cacheKey := fmt.Sprintf("%s:%s:infra_token", env, appName)

	contextLog := log.WithContext(c)

	response, err := get(c, cacheKey, cc)
	if response != nil {
		return response.AccessToken, nil
	}

	response, err = request(env)
	if err != nil {
		contextLog.WithError(err).Warn("get infra token fail")
		return "", err
	}

	if err = save(response, cacheKey, cc); err != nil {
		contextLog.WithError(err).Warn("cache infra token fail")
		return "", err
	}

	return response.AccessToken, nil
}

func request(env string) (*Response, error) {

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	authURL := fmt.Sprintf(AuthHostFormat, env)

	resp, err := http.Post(authURL, constant.ContentTypeApplicationJSONCharset, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	result := new(Response)
	if err = json.Unmarshal(body, result); err != nil {
		return nil, err
	}

	return result, nil
}

func get(c context.Context, key string, cc *cache.Cache) (*Response, error) {
	if cc == nil {
		return nil, nil
	}

	var wanted Response
	ctx, cancelFunc := context.WithTimeout(c, constant.TimeoutConn)
	if err := cc.Get(ctx, key, &wanted); err != nil {
		return nil, err
	}
	defer cancelFunc()
	return &wanted, nil

}

func save(response *Response, key string, cc *cache.Cache) error {
	if response != nil {
		ttl := time.Duration(response.ExpiresIn)
		ctx, cancelFunc := context.WithTimeout(context.Background(), constant.TimeoutConn)
		defer cancelFunc()
		if err := cc.Set(&cache.Item{
			Ctx:   ctx,
			Key:   key,
			Value: response,
			TTL:   ttl,
		}); err != nil {
			return err
		}
	}

	return nil
}
