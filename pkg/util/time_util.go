package util

import "time"

func GetTimestamp() int64 {
	return time.Now().Unix()
}

func ToTimestamp(seconds int64) time.Time {
	return time.Unix(seconds, 0)
}
