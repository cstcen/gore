package gocore

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func SnakeCase(camel string) string {
	snake := matchFirstCap.ReplaceAllString(camel, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func JoinVertical(format []string) string {
	return strings.Join(format, "|")
}

func UnixToDate(unix string) (string, error) {
	t, err := strconv.ParseInt(unix, 10, 64)
	if err != nil {
		return "", err
	}
	return time.Unix(t, 0).Format("2006-01-02 15:04:05"), nil
}
