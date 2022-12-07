package balance

import (
	"jcb/db"
)

func SetClosing(year int, balance int64) {
	db.SetClosing(year, balance)
	return
}

func GetClosing(year int) int64 {
	return db.GetClosing(year)
}
