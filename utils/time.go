package utils

import (
	"time"
)

func FormatTime(t string) (time.Time, error) {
	if t == "" {
		return time.Time{}, nil
	}
	layout := "02/01/2006, 03:04:05 PM"
	return time.Parse(layout, t)
}

func GetTimeNow() time.Time {
	return time.Now().UTC()
}
