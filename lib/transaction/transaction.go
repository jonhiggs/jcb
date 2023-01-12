package transaction

import (
	"jcb/db"
	"time"
)

func Sum(startTime time.Time, endTime time.Time) int64 {
	return db.TransactionSum(startTime, endTime)
}
