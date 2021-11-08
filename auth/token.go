package auth

import (
	"errors"
	"git.tenvine.cn/backend/gore/gonfig"
	goreHttp "git.tenvine.cn/backend/gore/http"
	"git.tenvine.cn/backend/gore/vo"
	"net/http"
)

func Check(token string) (*Member, error) {
	url := gonfig.Instance().GetString("tenvine.api.host") + gonfig.Instance().GetString("tenvine.api.verifyToken")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)
	resp, err := goreHttp.GetInstance().Do(req)
	if err != nil {
		return nil, err
	}
	member := new(Member)
	result := &vo.DataResult{Data: member}
	if err := goreHttp.RespHandler(resp, result); err != nil {
		return nil, err
	}
	if result.Code != vo.BaseResultSuccess.Code {
		return nil, errors.New("token verification failed")
	}
	if result.Data == nil {
		return nil, errors.New("no member was found during token verification")
	}

	if member.Agent != "XK5_SERVER" {
		return nil, errors.New("invalid token")
	}

	return member, nil
}
