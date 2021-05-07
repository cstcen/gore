package gore

import "time"

const (
	TimeoutConn = time.Second * 2

	FormatTimestamp = "2006-01-02 15:04:05"
	FormatDate      = "20060102"

	ContentTypeApplicationJSON        = "application/json"
	ContentTypeApplicationJSONCharset = "application/json; charset=UTF-8"
)
