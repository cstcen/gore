package infratoken

import (
	"context"
	"github.com/cstcen/gore/gonfig"
	goreHttp "github.com/cstcen/gore/http"
	"github.com/cstcen/gore/log"
	"github.com/cstcen/gore/util"
	"net/http"
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

func SetAuthorizationInHeader(request *http.Request) error {
	token, err := Get(request.Context())
	if err != nil {
		return err
	}
	request.Header.Set("Authorization", token)
	return nil
}

func Get(c context.Context) (string, error) {
	var response *Response
	var err error

	response, err = request(c)
	if err != nil {
		log.ErrorCf(c, "get infra token fail: %v", err)
		return "", err
	}

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
	if err := goreHttp.Post(c, url, "application/json; charset=utf8", req, result); err != nil {
		return nil, err
	}

	return result, nil
}
