package lock

import (
	"jcb/db"
)

func Create(id int64) error {
	return db.LockCreateId(id)
}

func Clear(id int64) error {
	return db.LockClearId(id)
}

// return the earliest editable timestamp
//func LockedUntil() time.Time {
//}
