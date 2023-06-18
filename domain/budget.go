package domain

import (
	"strings"
	"time"
)

type Budget struct {
	Id         int64
	Date       time.Time
	Category   string
	Cents      int64
	Notes      string
	Cumulative bool
}

func (b Budget) CategoryString() string {
	return strings.Trim(b.Category, " ")
}
