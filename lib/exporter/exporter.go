package exporter

import (
	"fmt"
	"jcb/lib/transaction"
)

func Tsv() {
	start, end := transaction.DateRange()
	for _, t := range transaction.All(start, end) {
		fmt.Printf("%s\t%s\t%s\t%s\t%s\n",
			t.Date.GetText(),
			t.GetCategory(false),
			t.Description.GetText(),
			t.GetAmount(false),
			t.GetNotes(),
		)
	}
}
