package time_util

import (
	"fmt"
	"time"
)

func GetUTCDate() string {
	now := time.Now().UTC()
	return fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day())
}

func GetUTCDateTime() string {
	now := time.Now().UTC()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
}
