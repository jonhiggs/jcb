package dates

import (
	"jcb/db"
	"time"
)

func FirstUncommitted() time.Time {
	return db.DateFirstUncommitted()
}

func LastCommitted() time.Time {
	return db.DateLastCommitted()
}

func LastUncommitted() time.Time {
	return db.DateLastUncommitted()
}
