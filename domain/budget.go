package domain

import "time"

type Budget struct {
	Id         int64
	Date       time.Time
	Category   string
	Cents      int64
	Notes      string
	Cumulative bool
}
