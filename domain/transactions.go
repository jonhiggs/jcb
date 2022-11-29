package domain

import "time"

type Transaction struct {
	Id          int64
	Date        time.Time
	Description string
	Cents       int64
}
