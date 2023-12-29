package common

const (
	LoginTypeQq LoginType = iota + 1
	LoginTypeWx
	LoginTypePhone
	LoginTypePwd
	LoginTypeSgxqMsdkQq
	LoginTypeSgxqMsdkWx
	LoginTypeWegame
	LoginTypeApple
)

var (
	LoginTypeStringMap = map[LoginType]string{
		LoginTypeQq:         "Qq",
		LoginTypeWx:         "Wx",
		LoginTypePhone:      "Phone",
		LoginTypePwd:        "Pwd",
		LoginTypeSgxqMsdkQq: "SgxqMsdkQq",
		LoginTypeSgxqMsdkWx: "SgxqMsdkWx",
		LoginTypeWegame:     "Wegame",
		LoginTypeApple:      "Apple",
	}
)

type LoginType uint8

func (t LoginType) IsQq() bool {
	return t == LoginTypeQq || t == LoginTypeSgxqMsdkQq
}

func (t LoginType) IsWx() bool {
	return t == LoginTypeWx || t == LoginTypeSgxqMsdkWx
}

func (t LoginType) String() string {
	s, _ := LoginTypeStringMap[t]
	return s
}
