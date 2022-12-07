package dates

import (
	"jcb/db"
	"time"
)

func Opening() time.Time {
	t, err := db.FindTransaction(0)
	if err != nil {
		return time.Unix(0, 0)
	}
	return t.Date
}

// year of -1 will scan all years
func FirstUncommitted(year int) time.Time {
	return db.DateFirstUncommitted(year)
}

// year of -1 will scan all years
func LastCommitted(year int) time.Time {
	return db.DateLastCommitted(year)
}

// year of -1 will scan all years
func LastUncommitted(year int) time.Time {
	return db.DateLastUncommitted(year)
}

// returns date of opening balance final day of the year after the final transaction
func Range() (time.Time, time.Time) {
	var start time.Time
	var end time.Time

	start = Opening()
	end = time.Date(start.Year()+1, 12, 31, 23, 59, 59, 59, time.UTC)

	lc := LastCommitted(-1)
	if lc.Unix() > end.Unix() {
		end = time.Date(lc.Year()+1, 12, 31, 23, 59, 59, 59, time.UTC)
	}

	lu := LastUncommitted(-1)
	if lu.Unix() > end.Unix() {
		end = time.Date(lu.Year()+1, 12, 31, 23, 59, 59, 59, time.UTC)
	}

	return start, end
}
