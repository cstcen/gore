package util

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z\\d])([A-Z])")
)

// SnakeCase CamelCase -> snake_case
func SnakeCase(camel string) string {
	snake := matchFirstCap.ReplaceAllString(camel, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// JoinVertical []string{"a", "b"} -> "a|b"
func JoinVertical(format []string) string {
	return strings.Join(format, "|")
}

// UnixToDate 1520386162 -> 2018-03-07 09:29:22
func UnixToDate(unix string) (string, error) {
	t, err := strconv.ParseInt(unix, 10, 64)
	if err != nil {
		return "", err
	}
	return time.Unix(t, 0).Format("2006-01-02 15:04:05"), nil
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func Strval(value any) string {
	var wanted string
	if value == nil {
		return wanted
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		wanted = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		wanted = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		wanted = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		wanted = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		wanted = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		wanted = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		wanted = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		wanted = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		wanted = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		wanted = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		wanted = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		wanted = strconv.FormatUint(it, 10)
	case string:
		wanted = value.(string)
	case []byte:
		wanted = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		wanted = string(newValue)
	}

	return wanted
}

func Interface2String(inter any) string {
	return Strval(inter)
}

func Interface2Int(inter any) int {
	var wanted int
	switch v := inter.(type) {
	case string:
		atoi, err := strconv.Atoi(v)
		if err == nil {
			wanted = atoi
		}
		break
	case int:
		wanted = v
	}
	return wanted
}

// FirstUpper get-item -> GetItem
func FirstUpper(s string) string {

	fields := strings.FieldsFunc(s, func(r rune) bool {
		return r == '-'
	})

	wanted := make([]string, len(fields))
	for index, item := range fields {
		split := strings.Split(item, "")
		split[0] = strings.ToUpper(split[0])
		wanted[index] = strings.Join(split, "")
	}

	return strings.Join(wanted, "")
}

// FirstLower GetItem -> get-item
func FirstLower(s string) string {

	if len(s) <= 0 {
		return s
	}

	var sb strings.Builder

	for i, r := range s {
		if unicode.IsUpper(r) {
			if i != 0 {
				sb.WriteRune('-')
			}
			sb.WriteRune(unicode.ToLower(r))
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

// IntsJoinComma []int{1,2} -> "1,2"
func IntsJoinComma(a []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), ","), "[]")
}
