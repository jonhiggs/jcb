package exporter

import (
	"fmt"
	"jcb/lib/transaction"
)

func Tsv() {
	start, end := transaction.DateRange()
	for _, t := range transaction.All(start, end) {
		fmt.Printf("%s\t%s\t%s\t%s\t%s\n",
			t.GetDateString(),
			t.GetCategory(false),
			t.GetDescription(false),
			t.GetAmount(false),
			t.GetNotes(),
		)
	}
}
