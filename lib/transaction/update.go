package transaction

func UpdateCategory(s string, ts []*Transaction) []*Transaction {
	var modifedTransactions []*Transaction

	for _, t := range ts {
		if t.Category.GetText() != s {
			t.Category.SetText(s)
			modifedTransactions = append(modifedTransactions, t)
		}
	}

	return modifedTransactions
}
