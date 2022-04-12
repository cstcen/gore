package auth

type Member struct {
	Agent          string `json:"agent,omitempty"`
	ApplicationNo  int    `json:"application_no,omitempty"`
	BirthDt        int    `json:"birth_dt,omitempty"`
	Check          string `json:"check,omitempty"`
	Env            string `json:"env,omitempty"`
	ExpireTime     int    `json:"expire_time,omitempty"`
	LastLaunchTime string `json:"last_launch_time,omitempty"`
	LogNo          int    `json:"log_no,omitempty"`
	LoginType      string `json:"login_type,omitempty"`
	MarketGameId   string `json:"market_game_id,omitempty"`
	MemberNo       int    `json:"member_no,omitempty"`
	Nickname       string `json:"nickname,omitempty"`
	OsId           string `json:"os_id,omitempty"`
	SvrId          string `json:"svr_id,omitempty"`
	Timestamp      int    `json:"timestamp,omitempty"`
	Token          string `json:"token,omitempty"`
	TransactionId  string `json:"transaction_id,omitempty"`
	OpenId         string `json:"open_id,omitempty"`
	ProviderOS     string `json:"provider_os,omitempty"`
	CharacterNo    int    `json:"character_no,omitempty"`
	ProfileImg     string `json:"profile_img,omitempty"`
}
