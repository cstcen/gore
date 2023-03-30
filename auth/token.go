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
	var result common.DataResult[*Member]
	if err := goreHttp.Do(req, &result); err != nil {
		return nil, err
	}
	if result.GetCode() != common.BaseResultSuccess.GetCode() || result.Data == nil {
		return nil, errors.New("token verification failed")
	}

	return result.Data, nil
}
