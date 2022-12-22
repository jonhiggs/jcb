package domain

import "time"

type Transaction struct {
	Id          int64
	Date        time.Time
	Description string
	Cents       int64
	Notes       string
}

type StringTransaction struct {
	Id          string
	Date        string
	Description string
	Cents       string
	Notes       string
}

type Attributes struct {
	Committed bool
	Repeated  bool
	Notes     bool
	Saved     bool
}
