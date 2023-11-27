package utils

import (
	"time"
)

func FormatTime(t string) (time.Time, error) {
	if t == "" {
		return time.Time{}, nil
	}
	layout := "02/01/2006, 03:04:05 PM -07:00"
	return time.Parse(layout, t)
}

func GetTimeNow() string {
	return time.Now().Format("2006-01-02 03:04:05 PM -07:00")
}
