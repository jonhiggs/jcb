package exporter

import (
	"fmt"
	stringf "jcb/lib/formatter/string"
	"jcb/lib/transaction"
)

func Tsv() {
	for _, t := range transaction.All() {
		st := stringf.Transaction(t)
		fmt.Printf("%s\t%s\t%s\t%s\t%s\n",
			st.Date,
			st.Category,
			st.Description,
			st.Cents,
			st.Notes,
		)
	}
}
