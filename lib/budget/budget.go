package budget

import (
	"jcb/db"
	"jcb/domain"
)

func Insert(t domain.Budget) (int64, error) {
	return db.InsertBudget(t)
}
