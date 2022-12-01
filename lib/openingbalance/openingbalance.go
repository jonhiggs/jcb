package openingBalance

import (
	"jcb/db"
)

func Find(year int64) (int64, error) {
	return db.FindOpeningBalance(year)
}
