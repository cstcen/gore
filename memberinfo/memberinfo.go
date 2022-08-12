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

func Get(ctx context.Context, memberNo int64) (*MemberInfo, error) {
	if goreCache.Instance() != nil {
		var memberInfo MemberInfo
		if goreCache.Instance().Get(ctx, cacheKey(memberNo), &memberInfo) == nil {
			return &memberInfo, nil
		}
	}

	url := gore.Viper().GetString("memberinfo.host") + gore.Viper().GetString("memberinfo.uri")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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
			Ctx:   ctx,
			Key:   cacheKey(memberNo),
			Value: memberInfo,
			TTL:   gore.Viper().GetDuration("memberinfo.duration"),
		}); err != nil {
			return nil, err
		}
	}

	return &memberInfo, nil
}

func cacheKey(memberNo int64) string {
	return fmt.Sprintf("MemberInfo:%v", memberNo)
}

type MemberInfo struct {
	MemberNo          int64  `json:"member_no"`
	CharacterNo       int    `json:"character_no"`
	Nickname          string `json:"nickname"`
	ProfileImg        string `json:"profile_img"`
	OpenId            string `json:"open_id"`
	ProviderOs        string `json:"provider_os"`
	ProviderCd        string `json:"provider_cd"`
	PwdType           int    `json:"pwd_type"`
	MemberType        int    `json:"member_type"`
	ServerId          string `json:"server_id"`
	ChannelId         string `json:"channel_id"`
	GameId            string `json:"game_id"`
	RegDt             string `json:"reg_dt"`
	CountryCd         string `json:"country_cd"`
	LoginClientIp     string `json:"login_client_ip"`
	LoginChannelNum   string `json:"login_channel_num"`
	RegChannelNum     string `json:"reg_channel_num"`
	WithdrawRequestDt string `json:"withdraw_request_dt"`
	GameMemberNo      int64  `json:"game_member_no"`
	Xk5MemberNo       int64  `json:"xk5_member_no"`
}
