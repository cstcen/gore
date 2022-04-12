package auth

import (
	"context"
	"errors"
	goreCache "git.tenvine.cn/backend/gore/db/cache"
	goreHttp "git.tenvine.cn/backend/gore/http"
	"git.tenvine.cn/backend/gore/log"
	"git.tenvine.cn/backend/gore/vo"
	"github.com/go-redis/cache/v8"
	"net/http"
	"time"
)

func Check(ctx context.Context, token string, url string) (*Member, error) {
	var member Member
	if goreCache.Instance() != nil {
		if goreCache.Instance().Get(ctx, cacheKey(token), &member) == nil {
			return &member, nil
		}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)
	resp, err := goreHttp.GetInstance().Do(req)
	if err != nil {
		return nil, err
	}
	result := vo.DataResult{Data: &member}
	if err := goreHttp.RespHandler(resp, &result); err != nil {
		return nil, err
	}
	if result.Code != vo.BaseResultSuccess.Code {
		return nil, errors.New("token verification failed")
	}
	if result.Data == nil {
		return nil, errors.New("no member was found during token verification")
	}

	if goreCache.Instance() != nil {
		if err := goreCache.Instance().Set(&cache.Item{
			Ctx:   ctx,
			Key:   cacheKey(token),
			Value: member,
			TTL:   time.Duration(member.ExpireTime-member.Timestamp) * time.Millisecond,
		}); err != nil {
			log.ErrorCE(ctx, err)
		}
	}

	return &member, nil
}
