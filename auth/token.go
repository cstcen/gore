package auth

import (
	"context"
	"errors"
	"git.tenvine.cn/backend/gore/common"
	goreHttp "git.tenvine.cn/backend/gore/http"
	"net/http"
)

func Check(ctx context.Context, token string, url string) (*Member, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)
	resp, err := goreHttp.GetInstance().Do(req)
	if err != nil {
		return nil, err
	}
	var result common.DataResult[Member]
	if err := goreHttp.RespHandler(resp, &result); err != nil {
		return nil, err
	}
	if result.Code() != common.BaseResultSuccess.Code() {
		return nil, errors.New("token verification failed")
	}

	return &result.Data, nil
}
