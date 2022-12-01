package openingBalance

import (
	"jcb/db"
)

func Find(year int64) (int64, error) {
	return db.FindOpeningBalance(year)
}

func Save(cents int64, year int64) error {
	return db.SaveOpeningBalance(cents, year)
}
