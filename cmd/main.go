package main

import (
	"jcb/db"
	"jcb/ui"
	"log"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Fatal(err)
	}

	ui.Start()

	//newT := domain.Transaction{-1, time.Now(), "test new t", 666}
	//db.InsertTransaction(newT)
	//if err != nil {
	//	log.Fatal(err)
	//}

	////trans, _ := db.UncommittedTransactions()
	////for _, t := range trans {
	////	fmt.Println(t.Id)
	////}

	//cs, _ := transaction.CommitSet(4, 30000)

	//fmt.Println(len(cs))
	//for _, t := range cs {
	//	fmt.Println(t.Id)
	//}
}
