package auth

type Member struct {
	Agent          string `json:"agent"`
	ApplicationNo  int    `json:"application_no"`
	BirthDt        int    `json:"birth_dt"`
	Check          string `json:"check"`
	Env            string `json:"env"`
	ExpireTime     int    `json:"expire_time"`
	LastLaunchTime string `json:"last_launch_time"`
	LogNo          int    `json:"log_no"`
	LoginType      string `json:"login_type"`
	MarketGameId   string `json:"market_game_id"`
	MemberNo       int    `json:"member_no"`
	Nickname       string `json:"nickname"`
	OsId           string `json:"os_id"`
	SvrId          string `json:"svr_id"`
	Timestamp      int    `json:"timestamp"`
	Token          string `json:"token"`
	TransactionId  string `json:"transaction_id"`
	OpenId         string `json:"open_id"`
	ProviderOS     string `json:"provider_os"`
	CharacterNo    string `json:"character_no"`
	ProfileImg     string `json:"profile_img"`
}
