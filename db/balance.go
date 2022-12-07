package db

import "log"

func SetClosing(year int, balance int64) {
	statement, err := db.Prepare("UPDATE balances SET closing = ? WHERE year = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(balance, year)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func GetClosing(year int) int64 {
	var closing int64
	statement, _ := db.Prepare("SELECT closing FROM balances WHERE year = ?")
	statement.QueryRow(year).Scan(&closing)
	return closing
}
