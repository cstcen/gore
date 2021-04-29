package gocore

import (
	"net/url"
	"strings"
)

func URLDecode(encodedString string) string {

	wanted, err := url.QueryUnescape(encodedString)

	if err != nil {
		return ""
	}

	return wanted

}

func URLEncode(str string) string {

	return strings.ReplaceAll(url.QueryEscape(str), "+", "%20")

}
