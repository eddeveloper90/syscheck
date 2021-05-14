package utils

import (
	"fmt"
	"time"
)

const DATETIME_LOG_FORMAT = "2006-01-02 15:04:05"

func Now() string {
	return time.Now().Format(DATETIME_LOG_FORMAT)
}

func Unix() int64 {
	return time.Now().Unix()
}

func UnixToTime(unix int64) time.Time {
	return time.Unix(unix, 0)
}

func UnixToDuration(unix int64) string {
	duration := Unix() - unix
	d, h, m, s := 0, 0, 0, 0
	if duration >= 24*3600 {
		d = int(duration / (24 * 3600))
		duration = duration % (24 * 3600)
	}

	if duration >= 3600 {
		h = int(duration / 3600)
		duration = duration % 3600
	}

	if duration >= 60 {
		m = int(duration / 60)
		duration = duration % 60
	}

	s = int(duration)

	return fmt.Sprintf("%d days %d hours %d minutes %d seconds", d, h, m, s)
}

func UnixToTs(unix int64) string {
	return UnixToTime(unix).Format(DATETIME_LOG_FORMAT)
}
