package domain

import (
	"time"
)

type Transaction struct {
	Id          int64
	Date        time.Time
	Description string
	Cents       int64
	Notes       string
	Category    string
}

type StringTransaction struct {
	Id          string
	Date        string
	Description string
	Cents       string
	Notes       string
	Category    string
}

type Attributes struct {
	Committed bool
	Notes     bool
	Saved     bool
	Credit    bool
}
