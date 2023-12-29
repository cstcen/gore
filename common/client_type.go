package common

const (
	// ClientTypeWeb Web应用
	ClientTypeWeb ClientType = iota
	// ClientTypeAndroid 安卓客户端应用
	ClientTypeAndroid
	// ClientTypeIos 苹果客户端应用
	ClientTypeIos
	// ClientTypePc 桌面客户端应用
	ClientTypePc
)

var (
	ClientTypeStringMap = map[ClientType]string{
		ClientTypeWeb:     "WEB",
		ClientTypeAndroid: "ANDROID",
		ClientTypeIos:     "IOS",
		ClientTypePc:      "PC",
	}
)

type ClientType uint8

func (c ClientType) String() string {
	s, _ := ClientTypeStringMap[c]
	return s
}
