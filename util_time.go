package gocore

import "time"

func MakeMillisecond(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func MakeLocalMillisecond(t time.Time) int64 {
	return MakeMillisecond(t.In(time.Local))
}
