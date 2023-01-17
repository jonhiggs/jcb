package db

import (
	"time"
)

const timeLayout = "2006-01-02 15:04:05.999999999-07:00"

func ParseDate(s string) time.Time {
	d, _ := time.Parse(timeLayout, s)
	return d
}

func TimeNow() string {
	return time.Now().Format(timeLayout)
}
