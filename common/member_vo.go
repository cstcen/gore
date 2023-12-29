package common

type MemberVO struct {
	// DateOfBirth 生日日期
	DateOfBirth *string `json:"date_of_birth,omitempty"`

	// Gender 性别
	// * `0` - 未知
	// * `1` - 男
	// * `2` - 女
	Gender int32 `json:"gender"`

	// MemberNo 用户唯一标识
	MemberNo int64 `json:"member_no"`

	// Nickname 昵称
	Nickname string `json:"nickname"`

	// OpenId xk5 openId
	OpenId string `json:"open_id"`

	// ProfileImg 用户头像
	ProfileImg string `json:"profile_img"`

	// RegTs 注册时的时间戳（MilliSecond）
	RegTs int64 `json:"reg_ts"`
}

type PunishmentVO struct {
	// Content 惩罚描述内容
	Content string `json:"content"`

	// PunishId 惩罚标识
	PunishId string `json:"punish_id"`

	// Title 惩罚标题
	Title string `json:"title"`

	// BeginTime 惩罚开始时间
	BeginTime int64 `json:"begin_time"`

	// EndTime 惩罚结束时间
	EndTime int64 `json:"end_time"`
}
