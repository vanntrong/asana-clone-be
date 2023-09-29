package utils

import "time"

func FormatTime(t string) (time.Time, error) {
	layout := "1/2/2006, 3:4:5 PM"
	return time.Parse(layout, t)
}
