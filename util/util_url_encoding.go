package util

import (
	"golang.org/x/text/encoding"
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

func URLEncodeAndUTF8(str string) string {
	enstr, err := encoding.Nop.NewDecoder().String(str)
	if err != nil {
		return str
	}

	escape := url.QueryEscape(enstr)

	return escape
}
