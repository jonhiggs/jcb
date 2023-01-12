package exporter

import (
	"fmt"
	"jcb/lib/transaction2"
)

func Tsv() {
	start, end := transaction2.DateRange()
	for _, t := range transaction2.All(start, end) {
		fmt.Printf("%s\t%s\t%s\t%s\t%s\n",
			t.GetDateString(),
			t.GetCategory(false),
			t.GetDescription(false),
			t.GetAmount(false),
			t.GetNotes(),
		)
	}
}
