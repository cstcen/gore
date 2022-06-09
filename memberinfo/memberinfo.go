package memberinfo

import (
	"context"
	"encoding/json"
	"fmt"
	"git.tenvine.cn/backend/gore"
	goreCache "git.tenvine.cn/backend/gore/db/cache"
	"github.com/go-redis/cache/v8"
	"io"
	"net/http"
	"strconv"
)

type Request struct {
	Ctx context.Context
}

func (r *Request) Get(memberNo int64) (*MemberInfo, error) {
	if goreCache.Instance() != nil {
		var memberInfo MemberInfo
		if goreCache.Instance().Get(r.Ctx, r.cacheKey(memberNo), &memberInfo) == nil {
			return &memberInfo, nil
		}
	}

	url := gore.Viper().GetString("memberinfo.host") + gore.Viper().GetString("memberinfo.uri")
	req, err := http.NewRequestWithContext(r.Ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("memberNo", strconv.FormatInt(memberNo, 10))
	req.URL.RawQuery = q.Encode()

	resp, err := gore.HttpClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var memberInfo MemberInfo
	if err := json.Unmarshal(body, &memberInfo); err != nil {
		return nil, err
	}

	if goreCache.Instance() != nil {
		if err := goreCache.Instance().Set(&cache.Item{
			Ctx:   r.Ctx,
			Key:   r.cacheKey(memberNo),
			Value: memberInfo,
			TTL:   gore.Viper().GetDuration("memberinfo.duration"),
		}); err != nil {
			return nil, err
		}
	}

	return &memberInfo, nil
}

func (r *Request) cacheKey(memberNo int64) string {
	return fmt.Sprintf("MemberInfo:%v", memberNo)
}

type MemberInfo struct {
	MemberNo    int64  `json:"member_no,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
	ProfileImg  string `json:"profile_img,omitempty"`
	CharacterNo int64  `json:"character_no,omitempty"`
	OpenId      string `json:"open_id,omitempty"`
	ProviderCd  string `json:"provider_cd,omitempty"`
	ProviderOs  string `json:"provider_os,omitempty"`
	PwdType     int    `json:"pwd_type,omitempty"`
	ServerId    string `json:"server_id,omitempty"`
	ChannelId   string `json:"channel_id,omitempty"`
	GameId      string `json:"game_id,omitempty"`
	RegDt       string `json:"reg_dt,omitempty"`
	CountryCd   string `json:"country_cd,omitempty"`
}
