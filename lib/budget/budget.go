package budget

import "jcb/domain"

func Insert(t domain.Budget) (int64, error) {
	//if t.Date.Unix() < dates.LastCommitted().Unix() {
	//	return -1, errors.New("Cannot insert before a committed transation")
	//}
	//return db.InsertTransaction(t)

	return -1, nil
}
