package auth

import (
	"context"
	"errors"
	"git.tenvine.cn/backend/gore/constant"
	goreCache "git.tenvine.cn/backend/gore/db/cache"
	goreHttp "git.tenvine.cn/backend/gore/http"
	"git.tenvine.cn/backend/gore/log"
	"git.tenvine.cn/backend/gore/vo"
	"github.com/go-redis/cache/v8"
	"net/http"
	"time"
)

func Check(ctx context.Context, token string, url string) (*Member, error) {
	if goreCache.Instance() != nil {
		log.DebugCf(ctx, "load member from cache")
		var member Member
		if goreCache.Instance().Get(ctx, cacheKey(token), &member) == nil {
			return &member, nil
		}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(constant.HeaderAuthorization, token)
	resp, err := goreHttp.GetInstance().Do(req)
	if err != nil {
		return nil, err
	}
	var result vo.DataResult[Member]
	if err := goreHttp.RespHandler(resp, &result); err != nil {
		return nil, err
	}
	if result.Code != vo.BaseResultSuccess.Code {
		return nil, errors.New("token verification failed")
	}

	if goreCache.Instance() != nil {
		expireTime := result.Data.ExpireTime
		timestamp := result.Data.Timestamp
		if err := goreCache.Instance().Set(&cache.Item{
			Ctx:   ctx,
			Key:   cacheKey(token),
			Value: result.Data,
			TTL:   time.Duration(expireTime-timestamp) * time.Millisecond,
		}); err != nil {
			log.ErrorCf(ctx, "set redis fail", err)
		}
	}

	return &result.Data, nil
}

func deleteCheckCache(token string) error {
	if goreCache.Instance() == nil {
		return nil
	}

	if err := goreCache.Instance().Delete(context.Background(), cacheKey(token)); err != nil {
		return err
	}

	return nil
}
